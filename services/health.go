package services

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.com/loveplus/data-ingest/proto"
)

type healthService struct {}

func (service *healthService) HealthCheck(context.Context, *empty.Empty) (*proto.HealthStatus, error) {
	return &proto.HealthStatus{Status: "OK"}, nil
}

// NewHealthService - Returns new implementation of health service
func NewHealthService() proto.HealthServer {
	return new(healthService)
}
