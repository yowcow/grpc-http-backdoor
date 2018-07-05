package main

import (
	"context"
	"flag"
	_ "fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/yowcow/test/grpc-test/service"
	"google.golang.org/grpc"
)

type Server struct {
	logger *log.Logger
}

func (s *Server) GetPerson(ctx context.Context, v *pb.Void) (*pb.Person, error) {
	s.logger.Println("rpc GetPerson() called")

	p := new(pb.Person)
	p.Id = proto.Int32(123)
	p.Name = proto.String("Hoge Fuga")
	p.Address = proto.String("234 Foo Bar")

	return p, nil
}

var (
	_ pb.DataServer = (*Server)(nil)
)

var network string
var address string
var logger *log.Logger

func init() {
	flag.StringVar(&network, "network", "tcp", "network to listen")
	flag.StringVar(&address, "address", ":9999", "address to listen")
	flag.Parse()
}

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	ln, err := net.Listen(network, address)
	if err != nil {
		logger.Fatalln(err)
	}

	grpcsv := grpc.NewServer()
	sv := &Server{logger}
	pb.RegisterDataServer(grpcsv, sv)

	logger.Println("Starting gRPC server")

	grpcsv.Serve(ln)
}
