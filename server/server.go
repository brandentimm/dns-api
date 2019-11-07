package main

import (
  "context"
  pb "github.com/brandentimm/dns-api/grpc"
  "google.golang.org/grpc"
  "log"
  "net"
)

type server struct {}

func (s *server) NewRecord(ctx context.Context, in *pb.NewRecordRequest) (*pb.NewRecordReply, error) {
  reply := &pb.NewRecordReply{
    ResponseCode:         0,
    ErrorMessage:         "",
  }
  log.Printf("Created DNS record with hostname: %s, zone: %s, ttl: %d", in.Host, in.Zone, in.Ttl)

  return reply, nil
}

func main() {
  lis, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalf("Error binding to tcp port: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterDNSServer(s, &server{})

  if err := s.Serve(lis); err != nil {
    log.Fatalf("server failure: %v", s.Serve(lis))
  }
}