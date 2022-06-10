package service_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample"
	"github.com/jasonrowsell/speccy/serializer"
	service "github.com/jasonrowsell/speccy/services/laptop"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestLaptopClient_CreateLaptop(t *testing.T) {
	t.Run("should create a new laptop and save it to the store", func(t *testing.T) {
		t.Parallel()

		// given
		laptopStore := service.NewInMemoryLaptopStore()
		serverAddress := startServer(t, laptopStore)
		laptopClient := newClient(t, serverAddress)

		// when
		laptop := sample.NewLaptop()
		request := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}

		response, err := laptopClient.CreateLaptop(context.Background(), request)

		// then
		require.NoError(t, err)
		assert.Equal(t, laptop.Id, response.Id)

		// and
		savedLaptop, err := laptopStore.FindLaptop(laptop.Id)
		require.NoError(t, err)
		assertSameLaptop(t, laptop, savedLaptop)
	})

	t.Run("should return an error if the laptop ID is invalid", func(t *testing.T) {
		t.Parallel()

		// given
		laptopStore := service.NewInMemoryLaptopStore()
		serverAddress := startServer(t, laptopStore)
		laptopClient := newClient(t, serverAddress)

		// when
		laptop := sample.NewLaptop()
		laptop.Id = "invalid"
		request := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}

		_, err := laptopClient.CreateLaptop(context.Background(), request)

		// then
		require.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
	})

	t.Run("should return an error if the laptop already exists", func(t *testing.T) {
		t.Parallel()

		// given
		laptopStore := service.NewInMemoryLaptopStore()
		serverAddress := startServer(t, laptopStore)
		laptopClient := newClient(t, serverAddress)

		// when we create a laptop
		laptop := sample.NewLaptop()
		request := &pb.CreateLaptopRequest{
			Laptop: laptop,
		}

		_, err := laptopClient.CreateLaptop(context.Background(), request)
		require.NoError(t, err)

		// when we try to create a laptop with the same ID
		_, err = laptopClient.CreateLaptop(context.Background(), request)

		// then
		require.Error(t, err)
		assert.Equal(t, codes.AlreadyExists, status.Code(err))

		// and
		savedLaptop, err := laptopStore.FindLaptop(laptop.Id)
		require.NoError(t, err)
		assertSameLaptop(t, laptop, savedLaptop)
	})
}

func startServer(t *testing.T, laptopStore service.LaptopStore) string {
	laptopServer := service.NewLaptopServer(laptopStore)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0") // random available port
	require.NoError(t, err)

	go grpcServer.Serve(listener)

	log.Println("server started on:", listener.Addr().String())

	return listener.Addr().String()
}

func newClient(t *testing.T, address string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	log.Println("client connected to:", address)
	return pb.NewLaptopServiceClient(conn)
}

func assertSameLaptop(t *testing.T, expected *pb.Laptop, actual *pb.Laptop) {
	laptopExpectJson, err := serializer.ProtobufToJSON(expected)

	if err != nil {
		t.Errorf("Error serializing laptop: %v", err)
	}

	laptopActualJson, err := serializer.ProtobufToJSON(actual)

	if err != nil {
		t.Errorf("Error serializing laptop: %v", err)
	}

	assert.Equal(t, laptopExpectJson, laptopActualJson)
}
