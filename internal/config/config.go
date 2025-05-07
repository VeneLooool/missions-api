package config

import (
	"context"
	"os"
)

const (
	EnvKeyHttpPort = "MISSION_HTTP_PORT"
	EnvKeyGrpcPort = "MISSION_GRPC_PORT"
)

type Config struct {
	HttpPort string
	GrpcPort string

	FieldsApiClientConfig *FieldsApiClientConfig
}

func New(ctx context.Context) (*Config, error) {
	return &Config{
		HttpPort: os.Getenv(EnvKeyHttpPort),
		GrpcPort: os.Getenv(EnvKeyGrpcPort),

		FieldsApiClientConfig: NewFieldsApiClientConfig(ctx),
	}, nil
}
