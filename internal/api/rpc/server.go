package rpc

import (
	"context"

	api "github.com/gkalombo21/mbongo-chain/internal/api"
	"github.com/gkalombo21/mbongo-chain/internal/runtime"
)

// Server wraps the base RPC server with runtime-specific wiring.
type Server struct {
	base    *api.RPCServer
	runtime *runtime.Engine
}

// NewServer constructs a server bound to the provided runtime engine.
func NewServer(rt *runtime.Engine) *Server {
	return &Server{
		base:    api.NewRPCServer(),
		runtime: rt,
	}
}

// Register attaches a handler to the base RPC server.
func (s *Server) Register(method string, handler func(context.Context, []byte) ([]byte, error)) {
	if s.base == nil {
		return
	}
	s.base.Register(method, handler)
}

// Listen delegates to the underlying RPC server implementation.
func (s *Server) Listen(address string) error {
	if s.base == nil {
		return nil
	}
	return s.base.Listen(address)
}

// Serve begins handling incoming requests using the runtime context.
func (s *Server) Serve(ctx context.Context) error {
	if s.base == nil {
		return nil
	}
	return s.base.Serve(ctx)
}

// Runtime exposes the attached runtime for handler registration.
func (s *Server) Runtime() *runtime.Engine {
	return s.runtime
}
