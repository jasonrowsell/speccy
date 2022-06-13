package main

import (
	"context"
	"flag"
	"log"

	"github.com/jasonrowsell/speccy/pb"
	"github.com/jasonrowsell/speccy/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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

	response, err := laptopClient.CreateLaptop(context.Background(), &pb.CreateLaptopRequest{
		Laptop: laptop,
	})

	if err != nil {
		log.Fatal("cannot create laptop: ", err)
	} else {
		log.Println("created laptop", response.GetId())
	}
}
