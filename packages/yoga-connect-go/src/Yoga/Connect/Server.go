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
	s.routes.Handle(procedure, connect.NewUnaryHandler(procedure, func(_ context.Context, request *connect.Request[[]byte]) (*connect.Response[[]byte], error) {
		answered := &call{request: *request.Msg}
		if failure := run(handle, answered); failure != nil {
			return nil, failure
		}
		return connect.NewResponse(&answered.response), nil
	}, rawBytes))
}

func ServerStreamImpl(procedure string, handle func(*call), s *server) {
	s.routes.Handle(procedure, connect.NewServerStreamHandler(procedure, func(_ context.Context, request *connect.Request[[]byte], stream *connect.ServerStream[[]byte]) error {
		return asError(run(handle, &call{request: *request.Msg, serverStream: stream}))
	}, rawBytes))
}

func ClientStreamImpl(procedure string, handle func(*call), s *server) {
	s.routes.Handle(procedure, connect.NewClientStreamHandler(procedure, func(_ context.Context, stream *connect.ClientStream[[]byte]) (*connect.Response[[]byte], error) {
		answered := &call{clientStream: stream}
		if failure := run(handle, answered); failure != nil {
			return nil, failure
		}
		return connect.NewResponse(&answered.response), nil
	}, rawBytes))
}

func BidiStreamImpl(procedure string, handle func(*call), s *server) {
	s.routes.Handle(procedure, connect.NewBidiStreamHandler(procedure, func(_ context.Context, stream *connect.BidiStream[[]byte, []byte]) error {
		return asError(run(handle, &call{bidiStream: stream}))
	}, rawBytes))
}

// run calls the handler and returns the call's failure. A panic is an internal error rather
// than a dropped connection.
func run(handle func(*call), answered *call) (failure *connect.Error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			failure = connect.NewError(connect.CodeInternal, fmt.Errorf("%v", recovered))
		}
	}()
	handle(answered)
	return answered.failure
}

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

func FailImpl(c *call, code int32, message string) {
	if c.failure == nil {
		c.failure = connect.NewError(connect.Code(code), errors.New(message))
	}
}

// ListenImpl binds the address and serves HTTP/1.1 and HTTP/2 without TLS, which gRPC clients
// use, in the background.
func ListenImpl[R any](failed func(string) R, bound func(int32) R, host string, port int32, s *server) R {
	listener, err := net.Listen("tcp", net.JoinHostPort(host, strconv.Itoa(int(port))))
	if err != nil {
		return failed(err.Error())
	}
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	s.http = &http.Server{Handler: s.routes, Protocols: protocols}
	eventloop.Hold()
	go s.http.Serve(listener)
	return bound(int32(listener.Addr().(*net.TCPAddr).Port))
}
