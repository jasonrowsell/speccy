package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jasonrowsell/speccy/pb"
	service "github.com/jasonrowsell/speccy/services"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "port to listen on")
	flag.Parse()
	log.Println("server listening on port", *port)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	laptopServer := service.NewLaptopServer(laptopStore, imageStore)

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	grpcServer.Serve(listener)
}
