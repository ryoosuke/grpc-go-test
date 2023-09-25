package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-test/rock-paper-scissors/pb"
	"log"
	"time"
)

var (
	addr       = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	name       = flag.String("name", "John", "The name of player")
	playerHand = flag.Int64("hand", 0, "The hand of player")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRPSClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DoGame(ctx, &pb.DoGameRequest{Name: *name, PlayerHand: *playerHand})
	if err != nil {
		log.Fatalf("could not start game: %v", err)
	}
	log.Print("The result: ", r.GetResult())
}
