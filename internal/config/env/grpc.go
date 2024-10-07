package env

import (
	"errors"
	"github.com/Chigvero/grpcProject/internal/config"
	"net"
	"os"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port string
}

func NewGRPCCongfig() (*grpcConfig, error) {
	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("incorrect port")
	}
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("incorrect host name")
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}
