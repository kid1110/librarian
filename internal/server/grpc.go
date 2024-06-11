package server

import (
	"github.com/tuihub/librarian/internal/conf"
	"github.com/tuihub/librarian/internal/lib/libapp"
	"github.com/tuihub/librarian/internal/lib/libauth"
	"github.com/tuihub/librarian/internal/lib/libobserve"
	"github.com/tuihub/librarian/internal/lib/libsentry"
	pb "github.com/tuihub/protos/pkg/librarian/sephirah/v1"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.SephirahServer,
	auth *libauth.Auth,
	greeter pb.LibrarianSephirahServiceServer,
	app *libapp.Settings,
	observer *libobserve.BuiltInObserver,
) (*grpc.Server, error) {
	validator, err := libapp.NewValidator()
	if err != nil {
		return nil, err
	}
	var middlewares = []middleware.Middleware{
		logging.Server(libapp.GetLogger()),
		ratelimit.Server(),
		tracing.Server(),
		validator,
	}
	if app.EnablePanicRecovery {
		middlewares = append(middlewares, recovery.Recovery())
	}
	middlewares = append(middlewares, libsentry.Server())
	middlewares = append(middlewares, libobserve.Server(observer))
	middlewares = append(middlewares, NewTokenMatcher(auth)...)
	var opts = []grpc.ServerOption{
		grpc.Middleware(middlewares...),
	}
	if c.GetGrpc().GetNetwork() != "" {
		opts = append(opts, grpc.Network(c.GetGrpc().GetNetwork()))
	}
	if c.GetGrpc().GetAddr() != "" {
		opts = append(opts, grpc.Address(c.GetGrpc().GetAddr()))
	}
	if c.GetGrpc().GetTimeout() != nil {
		opts = append(opts, grpc.Timeout(c.GetGrpc().GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterLibrarianSephirahServiceServer(srv, greeter)
	return srv, nil
}
