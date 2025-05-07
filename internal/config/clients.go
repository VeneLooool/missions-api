package config

import (
	"context"
	"os"
)

const (
	EnvKeyFieldsApiHost     = "FIELDS_API_HOST"
	EnvKeyFieldsApiGrpcPort = "FIELDS_API_GRPC_PORT"
)

type FieldsApiClientConfig struct {
	Host string
	Port string
}

func NewFieldsApiClientConfig(ctx context.Context) *FieldsApiClientConfig {
	return &FieldsApiClientConfig{
		Host: os.Getenv(EnvKeyFieldsApiHost),
		Port: os.Getenv(EnvKeyFieldsApiGrpcPort),
	}
}
