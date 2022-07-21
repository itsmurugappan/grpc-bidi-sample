package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pp "github.com/itsmurugappan/grpc-bidi-sample/pp"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	fmt.Println("Client Stream...")
	fmt.Println()

	opts := grpc.WithInsecure()
	con, err := grpc.Dial("localhost:3000", opts)
	if err != nil {
		log.Fatalf("Error connecting: %v \n", err)
	}

	defer con.Close()
	c := pp.NewPingPongClient(con)
	play(c)
}

func play(c pp.PingPongClient) {
	// Get the stream and a possible error from the CreateUser function
	stream, err := c.PingPong(context.Background())
	if err != nil {
		log.Fatalf("Error when getting stream object: %v", err)
		return
	}

	for {
		r := &pp.PP{Data: "ping"}
		pReq, _ := anypb.New(r)
		stream.Send(pReq)

		// Get response and possible error message from the stream
		resp, err := stream.Recv()

		// Break for loop if there are no more response messages
		if err == io.EOF {
			break
		}

		// Handle a possible error
		if err != nil {
			log.Fatalf("Error when receiving response: %v", err)
		}
		res := &pp.PP{}
		resp.UnmarshalTo(res)
		// Log the response
		fmt.Println(res.Data)
	}

}
