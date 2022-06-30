package service_test

import (
	"context"
	"testing"

	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample"
	service "github.com/jasonrowsell/speccy/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLaptopServer_CreateLaptop(t *testing.T) {
	t.Run("should create a new laptop and save it to the store", func(t *testing.T) {
		t.Parallel()
		// given
		laptop := sample.NewLaptop()
		laptopStore := service.NewInMemoryLaptopStore()
		server := service.NewLaptopServer(laptopStore, nil)

		// when
		request := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}

		response, err := server.CreateLaptop(context.Background(), request)

		// then
		require.NoError(t, err)
		assert.Equal(t, laptop.Id, response.Id)
	})

	t.Run("should return an error if the laptop already exists", func(t *testing.T) {
		t.Parallel()
		// given
		laptop := sample.NewLaptop()
		store := service.NewInMemoryLaptopStore()
		server := service.NewLaptopServer(store, nil)
		store.SaveLaptop(laptop)

		// when
		request := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}

		_, err := server.CreateLaptop(context.Background(), request)

		// then
		require.Error(t, err)
		assert.Equal(t, codes.AlreadyExists, status.Code(err))
		assert.Equal(t, "laptop with ID: "+laptop.Id+" already exists", status.Convert(err).Message())
	})

	t.Run("should return an error if the laptop ID is invalid", func(t *testing.T) {
		t.Parallel()
		// given
		laptop := sample.NewLaptop()
		laptop.Id = "invalid"
		store := service.NewInMemoryLaptopStore()
		server := service.NewLaptopServer(store, nil)

		// when
		request := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}

		_, err := server.CreateLaptop(context.Background(), request)

		// then
		require.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		assert.Equal(t, "invalid ID: parsing \"invalid\" failed: invalid UUID length: 7", status.Convert(err).Message())
	})

	t.Run("should assign a new ID if the laptop ID is empty", func(t *testing.T) {
		t.Parallel()
		// given
		emptyID := ""
		laptop := sample.NewLaptop()
		laptop.Id = emptyID
		store := service.NewInMemoryLaptopStore()
		server := service.NewLaptopServer(store, nil)

		// when
		request := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}

		response, err := server.CreateLaptop(context.Background(), request)

		// then
		require.NoError(t, err)
		assert.NotEqual(t, emptyID, response.Id)
	})
}
