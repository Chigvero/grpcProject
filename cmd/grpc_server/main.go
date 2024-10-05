package main

import (
	"context"
	"fmt"
	desc "github.com/Chigvero/grpcProject/pkg/note_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const (
	hostStr  = "localhost"
	grpcPort = 50051
)

type server struct {
	desc.UnimplementedNoteV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	id := req.GetId()
	log.Printf("note id: %v\n", id)
	return &desc.GetResponse{
		Note: &desc.Note{
			Id: id,
			Info: &desc.NoteInfo{
				Title:    gofakeit.BeerName(),
				Context:  gofakeit.BeerName(),
				Author:   gofakeit.Name(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Error with listening port:%v\n", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})
	log.Println("server listening at:", grpcPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
