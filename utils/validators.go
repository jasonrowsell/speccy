package utils

import (
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateLaptopId(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid ID: parsing \"%s\" failed: %v", id, err)
	}
	return nil
}
