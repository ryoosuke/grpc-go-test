package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-go-test/rock-paper-scissors/pb"
	"log"
	"math/rand"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port number")
)

type server struct {
	pb.UnimplementedRPSServer
}

type hand error

/**
0 = rock
1 = paper
2 = scissors
**/

func (s *server) DoGame(ctx context.Context, in *pb.DoGameRequest) (*pb.Results, error) {
	playerHand := in.GetPlayerHand()
	if playerHand < 0 || playerHand > 2 {
		log.Fatalf("Invalid hand: %v", playerHand)
		return nil, errors.New("Invalid hand")
	}
	computerHand := rand.Int63n(3)
	log.Print("Player(", in.GetName(), "): ", in.GetPlayerHand(), ", Computer: ", computerHand)
	if playerHand == computerHand {
		return &pb.Results{Result: "TIE"}, nil
	} else if playerHand == 0 && computerHand == 1 || playerHand == 1 && computerHand == 2 || playerHand == 2 && computerHand == 0 {
		return &pb.Results{Result: "LOSE"}, nil
	} else {
		return &pb.Results{Result: "WIN"}, nil
	}
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRPSServer(s, &server{})
	log.Print("Starting server on port: ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
