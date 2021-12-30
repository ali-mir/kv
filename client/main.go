package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "kv/proto"
	"log"
	"time"
)

var conn *grpc.ClientConn
var client pb.KVClient

func Connect(port *int) *grpc.ClientConn {
	log.Printf("connecting to server at port %d", *port)

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *port), grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to connect to server")
	}

	client = pb.NewKVClient(conn)

	return conn
}

func Insert(key, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Insert(ctx, &pb.InsertRequest{Key: key, Value: value})
	if err != nil || r.GetSuccess() != true {
		log.Fatalf("insert failed: %v", err)
	}
}

func Lookup(key string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Lookup(ctx, &pb.LookupRequest{Key: key})
	if err != nil {
		log.Fatalf("lookup failed: %v", err)
	}
	return r.GetValue()
}

func Delete(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Delete(ctx, &pb.DeleteRequest{Key: key})
	if err != nil || r.GetSuccess() != true {
		log.Fatalf("delete failed: %v", err)
	}
}
