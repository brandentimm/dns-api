package main

import (
  "context"
  pb "github.com/brandentimm/dns-api/grpc"
  "google.golang.org/grpc"
  "log"
  "net"
)

type server struct {
  requests []*pb.NewRecordRequest
}

func (s *server) NewRecord(ctx context.Context, in *pb.NewRecordRequest) (*pb.NewRecordReply, error) {
  reply := &pb.NewRecordReply{
    ResponseCode:         0,
    ErrorMessage:         "",
  }
  log.Printf("Created DNS record with hostname: %s, zone: %s, ttl: %d", in.Host, in.Zone, in.Ttl)

  s.requests = append(s.requests, in)
  return reply, nil
}

func (s *server) RequestStream(st *pb.NewRequestStream, stream pb.DNS_RequestStreamServer) error {
  ctx := stream.Context()

  var requestIndex int

  for {
    select {
    case <- ctx.Done():
      return ctx.Err()
      default:
    }

    if len(s.requests) > requestIndex-1 {
      for _, request := range s.requests[requestIndex:] {
        err := stream.Send(request)
        if err != nil {
          log.Fatalf("Error streaming requests to client: %v", err)
        }
        requestIndex++
      }
    }
  }
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