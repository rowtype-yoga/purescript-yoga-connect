package ffi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"sync"

	"connectrpc.com/connect"
)

// Connection is a server's address and the HTTP/2 connections to it.
type Connection = *connection

// ClientCall is one streaming call in flight. Exactly one of the streams is set.
type ClientCall = *clientCall

type connection struct {
	http    *http.Client
	address string
	clients sync.Map
}

type clientCall struct {
	serverStream *connect.ServerStreamForClient[[]byte]
	clientStream *connect.ClientStreamForClient[[]byte, []byte]
	bidiStream   *connect.BidiStreamForClient[[]byte, []byte]
	received     []byte
	failure      error
}

// ConnectImpl speaks HTTP/2 without TLS, as gRPC servers do on a plain port.
func ConnectImpl(address string) *connection {
	protocols := new(http.Protocols)
	protocols.SetUnencryptedHTTP2(true)
	return &connection{http: &http.Client{Transport: &http.Transport{Protocols: protocols}}, address: address}
}

// client is connect-go's client for the procedure, made once per connection.
func (c *connection) client(procedure string) *connect.Client[[]byte, []byte] {
	if existing, ok := c.clients.Load(procedure); ok {
		return existing.(*connect.Client[[]byte, []byte])
	}
	made := connect.NewClient[[]byte, []byte](c.http, c.address+procedure, connect.WithGRPC(), rawBytes)
	existing, _ := c.clients.LoadOrStore(procedure, made)
	return existing.(*connect.Client[[]byte, []byte])
}

func CallUnaryImpl[R any](failed func(int32) func(string) R, succeeded func([]byte) R, c *connection, procedure string, request []byte) R {
	response, err := c.client(procedure).CallUnary(context.Background(), connect.NewRequest(&request))
	if err != nil {
		return failure(failed, err)
	}
	return succeeded(*response.Msg)
}

func OpenServerStreamImpl(c *connection, procedure string, request []byte) *clientCall {
	stream, err := c.client(procedure).CallServerStream(context.Background(), connect.NewRequest(&request))
	return &clientCall{serverStream: stream, failure: err}
}

func OpenClientStreamImpl(c *connection, procedure string) *clientCall {
	return &clientCall{clientStream: c.client(procedure).CallClientStream(context.Background())}
}

func OpenBidiStreamImpl(c *connection, procedure string) *clientCall {
	return &clientCall{bidiStream: c.client(procedure).CallBidiStream(context.Background())}
}

func SendImpl(call *clientCall, request []byte) bool {
	if call.bidiStream != nil {
		return call.bidiStream.Send(&request) == nil
	}
	return call.clientStream.Send(&request) == nil
}

func CloseSendingImpl(call *clientCall) {
	call.bidiStream.CloseRequest()
}

func CloseAndReceiveImpl[R any](failed func(int32) func(string) R, succeeded func([]byte) R, call *clientCall) R {
	response, err := call.clientStream.CloseAndReceive()
	if err != nil {
		return failure(failed, err)
	}
	return succeeded(*response.Msg)
}

func ReceiveImpl(call *clientCall) bool {
	if call.failure != nil {
		return false
	}
	if call.bidiStream != nil {
		received, err := call.bidiStream.Receive()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				call.failure = err
			}
			return false
		}
		call.received = *received
		return true
	}
	if !call.serverStream.Receive() {
		call.failure = call.serverStream.Err()
		return false
	}
	call.received = *call.serverStream.Msg()
	return true
}

func ReceivedImpl(call *clientCall) []byte { return call.received }

func OutcomeImpl[R any](failed func(int32) func(string) R, finished R, call *clientCall) R {
	CloseImpl(call)
	if call.failure != nil {
		return failure(failed, call.failure)
	}
	return finished
}

func CloseImpl(call *clientCall) {
	switch {
	case call.bidiStream != nil:
		call.bidiStream.CloseResponse()
	case call.serverStream != nil:
		call.serverStream.Close()
	}
}

func failure[R any](failed func(int32) func(string) R, err error) R {
	var connectError *connect.Error
	if errors.As(err, &connectError) {
		return failed(int32(connectError.Code()))(connectError.Message())
	}
	return failed(int32(connect.CodeUnknown))(err.Error())
}
