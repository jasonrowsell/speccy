package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	LaptopStore LaptopStore
	ImageStore  ImageStore
}

func NewLaptopServer(
	laptopStore LaptopStore,
	imageStore ImageStore,
) *LaptopServer {
	return &LaptopServer{
		LaptopStore: laptopStore,
		ImageStore:  imageStore,
	}
}

// Unary RPC handler to create a new laptop and save it to the store
func (server *LaptopServer) CreateLaptop(
	context context.Context,
	request *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := request.GetLaptop()
	log.Printf("CreateLaptop: %v", laptop)

	if len(laptop.GetId()) > 0 {
		if err := utils.ValidateLaptopId(laptop.GetId()); err != nil {
			return nil, err
		}
	} else {
		if err := assignLaptopId(laptop); err != nil {
			return nil, err
		}
	}

	if err := contextError(context); err != nil {
		return nil, err
	}

	if err := saveLaptop(server.LaptopStore, laptop); err != nil {
		return nil, err
	} else {
		log.Printf("Saved laptop with ID: %s", laptop.GetId())
	}

	res := &pb.CreateLaptopResponse{
		Id: laptop.GetId(),
	}

	return res, nil
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
	if err := store.SaveLaptop(laptop); err != nil {
		return err
	}
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Errorf(codes.Canceled, "request cancelled")
	case context.DeadlineExceeded:
		return status.Errorf(codes.DeadlineExceeded, "request timed out")
	default:
		return nil
	}
}

// Client-streaming RPC handler to upload a laptop image
func (server *LaptopServer) UploadLaptopImage(
	stream pb.LaptopService_UploadLaptopImageServer,
) error {
	log.Printf("UploadLaptopImage: %v", stream)
	return nil
}
