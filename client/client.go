package main

import (
    "context"
    "log"
    "os"

    pb "github.com/ry-itto/gRPC_sample"
    "google.golang.org/grpc"
)

func main() {
    addr := "localhost:8080"
    conn, err := grpc.Dial(addr, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewHelloClient(conn)

    name := os.Args[1]

    ctx := context.Background()
    r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not hello: %v", err)
    }
    log.Printf("Hello: %s", r.Message)
}
