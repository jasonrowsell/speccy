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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
