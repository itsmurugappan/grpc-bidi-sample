package main

import (
	"io"
	"net"
	"time"

	pp "github.com/itsmurugappan/grpc-bidi-sample/pp"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

type server struct {
	pp.UnimplementedPingPongServer
}

func main() {
	log.Println("Starting the server...")

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Unable to listen on port 3000: %v", err)
	}

	s := grpc.NewServer()
	pp.RegisterPingPongServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) PingPong(stream pp.PingPong_PingPongServer) error {
	log.Println("CreateUser Function")

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
		resp := &pp.PP{}
		req.UnmarshalTo(resp)
		log.Println(resp.Data)

		// Build and send response to the client
		r := &pp.PP{Data: "pong"}
		pReq, _ := anypb.New(r)
		stream.Send(pReq)
	}
}
