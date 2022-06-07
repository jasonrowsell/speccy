package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jasonrowsell/speccy/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
}

func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

// Unary RPC handler to create a new laptop and save it to the store
func (server *LaptopServer) CreateLaptop(
	context context.Context,
	request *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := request.GetLaptop()
	log.Printf("CreateLaptop with ID: %s", laptop.GetId())

	if len(laptop.GetId()) > 0 {
		if err := validateLaptopId(laptop.GetId()); err != nil {
			return nil, err
		} else if err := assignLaptopId(laptop); err != nil {
			return nil, err
		}
	}

	if err := saveLaptop(server.Store, laptop); err != nil {
		return nil, err
	} else {
		log.Printf("Created laptop with ID: %s", laptop.GetId())
	}

	res := &pb.CreateLaptopResponse{
		Id: laptop.GetId(),
	}

	return res, nil
}

func validateLaptopId(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return status.Errorf(codes.InvalidArgument, "invalid ID: %v", err)
	}
	return nil
}

func assignLaptopId(laptop *pb.Laptop) error {
	if len(laptop.GetId()) > 0 {
		return nil
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return status.Errorf(codes.Internal, "internal error: %v", err)
	}
	laptop.Id = id.String()

	return nil
}

func saveLaptop(store LaptopStore, laptop *pb.Laptop) error {
	err := store.SaveLaptop(laptop)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot save laptop: %v", err)
	} else {
		log.Printf("Created laptop with ID: %s", laptop.GetId())
	}

	return nil
}
