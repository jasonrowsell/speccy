package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const timeout = 500 * time.Millisecond

func main() {
	serverAddress := flag.String("server", "", "server address")

	flag.Parse()
	log.Println("dialing server", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	} else {
		log.Println("dialed server")
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	request := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	response, err := laptopClient.CreateLaptop(ctx, request)

	if err != nil {
		log.Fatal("cannot create laptop: ", err)
	} else {
		log.Println("created laptop", response.GetId())
	}
}
