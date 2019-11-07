package main

import (
  "context"
  "fmt"
  pb "github.com/brandentimm/dns-api/grpc"
  "google.golang.org/grpc"
  "log"
  "os"
  "strconv"
)

func main() {
  if len(os.Args) < 4 {
    fmt.Println("Must provide hostname, zone, ttl")
  }

  host := os.Args[1]
  zone := os.Args[2]

  ttlInt, err := strconv.Atoi(os.Args[3])
  if err != nil {
    log.Fatalf("Cannot convert %d to numeric TTL", os.Args[3])
  }
  ttl := int32(ttlInt)

  conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Error dialing service: %v", err)
  }
  defer conn.Close()

  client := pb.NewDNSClient(conn)

  request := &pb.NewRecordRequest{
    Host:                 host,
    Zone:                 zone,
    Ttl:                  ttl,
  }

  _, err = client.NewRecord(context.Background(), request)
  if err != nil {
    log.Fatalf("Error creating DNS record: %v", err)
    return
  }

  fmt.Printf("Successfully created dns record %s.%s with TTL %d\n", host, zone, ttl)
}