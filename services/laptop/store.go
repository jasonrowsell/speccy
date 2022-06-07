package service

import (
	"sync"

	"github.com/jasonrowsell/speccy/pb"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopStore interface {
	SaveLaptop(laptop *pb.Laptop) error
}

type InMemoryLaptopStore struct {
	mutex sync.Mutex
	store map[string]*pb.Laptop
}

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		store: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) SaveLaptop(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.store[laptop.Id] != nil {
		return status.Errorf(codes.AlreadyExists, "laptop with ID %s already exists", laptop.Id)
	}

	newLaptop := &pb.Laptop{}
	err := copier.Copy(newLaptop, laptop)

	if err != nil {
		return status.Errorf(codes.Internal, "internal error: %v", err)
	}

	store.store[laptop.Id] = newLaptop

	return nil
}
