package main

import (
	"fmt"
	"io"
	"net"
	"time"
	pp "github.com/itsmurugappan/grpc-bidi-sample/pp"

	"log"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	log.Info("Starting the server...")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	pp.RegisterUsersServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// CreateUser function
func (*server) PingPong(stream pp.PingPong_PingPongServer) error {
	log.Info("CreateUser Function")

	for {
		time.Sleep(30 * time.Second)
		// Receive the request and possible error from the stream object
		req, err := stream.Recv()

		// If there are no more requests, we return
		if err == io.EOF {
			return nil
		}

		// Handle error from the stream object
		if err != nil {
			log.Fatalf("Error when reading client request stream: %v", err)
		}

		log.Println(req.Data)

		// Build and send response to the client
		stream.Send(&pp.PP{
			Data:  "pong",
		})
	}
}