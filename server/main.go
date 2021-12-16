package main

import (
	"GRPCDEMO/AUTH"
	pb "GRPCDEMO/AUTH"
	"context"
	"errors"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type server struct {
	pb.UnimplementedAUTHENTICATIONServer
}

func (s *server) Authenticate(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	log.Printf("Received: %v", in.GetUsername())
	return &pb.LoginReply{Message: "Logged in as:" + in.GetUsername()}, nil
}
func (s *server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	log.Printf("Received: %v", in.GetUsername())
	return &pb.RegisterReply{Message: "Register in as:" + in.GetUsername()}, nil
}
func (s *server) Upload(ctx context.Context, stream AUTH.GuploadService_UploadServer) (err error) {
	for {
		_, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				goto END
			}
			err = errors.Unwrap(err)

		}
	}
END:
	err = stream.SendAndClose(&AUTH.UploadStatus{
		Message: "Upload received with success",
		Code:    AUTH.UploadStatusCode_Ok,
	})
	return
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", lis.Addr())
	}
	s := grpc.NewServer()
	pb.RegisterAUTHENTICATIONServer(s, &server{})
	log.Printf("server liserning at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server:%v", err)
	}
}
