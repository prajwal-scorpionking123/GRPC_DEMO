package main

import (
	pb "GRPCDEMO/AUTH"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	c := pb.NewAUTHENTICATIONClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Authenticate(ctx, &pb.LoginRequest{Username: "prajwalp", Password: "12345"})

	if err != nil {
		log.Println("could not greet:", err)
	}
	log.Printf("response: %s", r.GetMessage())

}
