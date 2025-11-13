package api

import (
	"context"
	"net"
)

// RPCServer exposes minimal controls for registering services and managing
// lifecycle events. Concrete protocol handling will be added in future work.
type RPCServer struct {
	listener net.Listener
	handlers map[string]func(context.Context, []byte) ([]byte, error)
}

// NewRPCServer constructs an unbound server.
func NewRPCServer() *RPCServer {
	return &RPCServer{
		handlers: make(map[string]func(context.Context, []byte) ([]byte, error)),
	}
}

// Register associates a handler function with an RPC method name.
// TODO: add middleware, authentication, and version negotiation.
func (s *RPCServer) Register(method string, handler func(context.Context, []byte) ([]byte, error)) {
	s.handlers[method] = handler
}

// Listen binds the server to an address.
// TODO: support TLS and graceful shutdown.
func (s *RPCServer) Listen(address string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	s.listener = ln
	return nil
}

// Serve starts handling connections.
// TODO: implement request routing and concurrency model.
func (s *RPCServer) Serve(ctx context.Context) error {
	if s.listener == nil {
		return nil
	}
	<-ctx.Done()
	return s.listener.Close()
}

