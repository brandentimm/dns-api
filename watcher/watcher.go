package main

import (
  "context"
  "fmt"
  pb "github.com/brandentimm/dns-api/grpc"
  "google.golang.org/grpc"
  "log"
)

func main() {
  conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Error dialing service: %v", err)
  }
  defer conn.Close()

  client := pb.NewDNSClient(conn)


  stream, err := client.RequestStream(context.Background(), &pb.NewRequestStream{})
  if err != nil {
    log.Fatalf("Error streaming requests: %v", err)
    return
  }

  for {
    req, err := stream.Recv()
    if err != nil {
      fmt.Printf("Received error reading stream: %v", err)
      break
    }
    log.Printf("Received stream message: %v", req)
  }
}