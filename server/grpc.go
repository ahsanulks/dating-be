package server

import (
	v1 "app/api/helloworld/v1"
	"app/configs"
	"app/internal/service"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
// func NewGRPCServer(c *configs.ApplicationConfig, logger log.Logger) *grpc.Server {
func NewGRPCServer(c *configs.ApplicationConfig, greeter *service.GreeterService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Server.GRPC.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.GRPC.Addr))
	}
	if c.Server.GRPC.Timeout != 0 {
		opts = append(opts, grpc.Timeout(time.Duration(c.Server.GRPC.Timeout)))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	return srv
}
