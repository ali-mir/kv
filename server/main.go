package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "kv/proto"
	store "kv/server/store"
	"log"
	"net"
	"os"
)

var port *int = flag.Int("port", 20020, "port")

// RPC handlers
type server struct {
	pb.UnimplementedKVServer
}

func (s *server) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertReply, error) {
	ok := store.Insert(req.GetKey(), req.GetValue())
	return &pb.InsertReply{Success: ok}, nil
}

func (s *server) Lookup(ctx context.Context, req *pb.LookupRequest) (*pb.LookupReply, error) {
	val := store.Lookup(req.GetKey())
	return &pb.LookupReply{Value: val}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	ok := store.Delete(req.GetKey())
	return &pb.DeleteReply{Success: ok}, nil
}

func main() {
	flag.Parse()
	log.SetOutput(os.Stdout)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to start listening: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKVServer(s, &server{})

	store.Initialize()

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
