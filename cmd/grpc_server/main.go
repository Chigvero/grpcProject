package main

import (
	"context"
	"flag"
	"github.com/Chigvero/grpcProject/internal/config"
	"github.com/Chigvero/grpcProject/internal/config/env"
	desc "github.com/Chigvero/grpcProject/pkg/note_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

var configPath string

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

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Error with loading configPath:%v\n", err)
	}
	grpcConfig, err := env.NewGRPCCongfig()
	if err != nil {
		log.Fatalf("failed to get grpcConfig:%v", err)
	}
	_, err = env.NewPGConfig()
	//need to add pgConfig
	if err != nil {
		log.Fatalf("failed to get grpcConfig:%v", err)
	}
	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("Error with listening port:%v\n", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterNoteV1Server(s, &server{})
	log.Println("server listening at:", grpcConfig)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
