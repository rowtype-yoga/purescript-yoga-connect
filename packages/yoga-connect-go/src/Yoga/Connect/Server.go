package ffi

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"

	"connectrpc.com/connect"
	"github.com/i-am-the-slime/iris-go/eventloop"
)

// Server is the routes and, once listening, the HTTP server.
type Server = *server

// Call is one call a route answers.
type Call = *call

type server struct {
	routes *http.ServeMux
	http   *http.Server
}

// call holds what the PureScript handler reads and writes. Exactly one of the streams is set
// for a streaming call.
type call struct {
	context      context.Context
	request      []byte
	response     []byte
	received     []byte
	failure      *connect.Error
	serverStream *connect.ServerStream[[]byte]
	clientStream *connect.ClientStream[[]byte]
	bidiStream   *connect.BidiStream[[]byte, []byte]
}

func CreateServerImpl() *server {
	return &server{routes: http.NewServeMux()}
}

func UnaryImpl(procedure string, handle func(*call), s *server) {
	s.routes.Handle(procedure, connect.NewUnaryHandler(procedure, func(ctx context.Context, request *connect.Request[[]byte]) (*connect.Response[[]byte], error) {
		answered := &call{context: ctx, request: *request.Msg}
		if failure := run(handle, answered); failure != nil {
			return nil, failure
		}
		return connect.NewResponse(&answered.response), nil
	}, rawBytes))
}

func ServerStreamImpl(procedure string, handle func(*call), s *server) {
	s.routes.Handle(procedure, connect.NewServerStreamHandler(procedure, func(ctx context.Context, request *connect.Request[[]byte], stream *connect.ServerStream[[]byte]) error {
		return asError(run(handle, &call{context: ctx, request: *request.Msg, serverStream: stream}))
	}, rawBytes))
}

func ClientStreamImpl(procedure string, handle func(*call), s *server) {
	s.routes.Handle(procedure, connect.NewClientStreamHandler(procedure, func(ctx context.Context, stream *connect.ClientStream[[]byte]) (*connect.Response[[]byte], error) {
		answered := &call{context: ctx, clientStream: stream}
		if failure := run(handle, answered); failure != nil {
			return nil, failure
		}
		return connect.NewResponse(&answered.response), nil
	}, rawBytes))
}

func BidiStreamImpl(procedure string, handle func(*call), s *server) {
	s.routes.Handle(procedure, connect.NewBidiStreamHandler(procedure, func(ctx context.Context, stream *connect.BidiStream[[]byte, []byte]) error {
		return asError(run(handle, &call{context: ctx, bidiStream: stream}))
	}, rawBytes))
}

// run calls the handler and returns the call's failure. A panic is an internal error rather
// than a dropped connection. A call the client cancelled, or whose deadline passed, ends with
// that, whatever the handler made of it.
func run(handle func(*call), answered *call) (failure *connect.Error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			failure = connect.NewError(connect.CodeInternal, fmt.Errorf("%v", recovered))
		}
		if ended := answered.context.Err(); ended != nil {
			failure = connect.NewError(codeOfEnded(ended), ended)
		}
	}()
	handle(answered)
	return answered.failure
}

func codeOfEnded(ended error) connect.Code {
	if errors.Is(ended, context.DeadlineExceeded) {
		return connect.CodeDeadlineExceeded
	}
	return connect.CodeCanceled
}

// ContextImpl is the call's context, done once the client cancels or the deadline passes.
func ContextImpl(c *call) context.Context { return c.context }

func asError(failure *connect.Error) error {
	if failure == nil {
		return nil
	}
	return failure
}

func RequestImpl(c *call) []byte { return c.request }

func RespondImpl(c *call, response []byte) { c.response = response }

func SendImpl(c *call, response []byte) bool {
	if c.bidiStream != nil {
		return c.bidiStream.Send(&response) == nil
	}
	return c.serverStream.Send(&response) == nil
}

func ReceiveImpl(c *call) bool {
	if c.bidiStream != nil {
		received, err := c.bidiStream.Receive()
		if err != nil {
			c.ended(err)
			return false
		}
		c.received = *received
		return true
	}
	if !c.clientStream.Receive() {
		c.ended(c.clientStream.Err())
		return false
	}
	c.received = *c.clientStream.Msg()
	return true
}

// ended notes why the client's stream ended, unless the client simply sent all it had.
func (c *call) ended(err error) {
	if err != nil && !errors.Is(err, io.EOF) && c.failure == nil {
		c.failure = connect.NewError(connect.CodeOf(err), err)
	}
}

func ReceivedImpl(c *call) []byte { return c.received }

// FailedImpl says whether the call has failed, such as when the client's stream broke off.
func FailedImpl(c *call) bool { return c.failure != nil }

func FailImpl(c *call, code int, message string) {
	if c.failure == nil {
		c.failure = connect.NewError(connect.Code(code), errors.New(message))
	}
}

// ListenImpl binds the address and serves HTTP/1.1 and HTTP/2 without TLS, which gRPC clients
// use, in the background.
func ListenImpl[R any](failed func(string) R, bound func(int) R, host string, port int, s *server) R {
	listener, err := net.Listen("tcp", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		return failed(err.Error())
	}
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	s.http = &http.Server{Handler: s.routes, Protocols: protocols}
	eventloop.Hold()
	go s.http.Serve(listener)
	return bound(listener.Addr().(*net.TCPAddr).Port)
}
