package main

import (
	"context"
	"errors"
	"log"
	"math/rand"

	"github.com/golang/protobuf/proto"
	"github.com/yowcow/test/grpc-test/server/httpserver"
	pb "github.com/yowcow/test/grpc-test/service"
	"google.golang.org/grpc"
)

var (
	_ pb.DataServer = (*App)(nil)
)

type App struct {
	logger *log.Logger
}

func (app *App) GetPerson(ctx context.Context, v *pb.Void) (*pb.Person, error) {
	app.logger.Println("rpc GetPerson() called")

	// Randomly return not found error
	if n := rand.Intn(4); n == 0 {
		return nil, errors.New("This time person was not found")
	}

	p := new(pb.Person)
	p.Id = proto.Int32(123)
	p.Name = proto.String("Hoge Fuga")
	p.Address = proto.String("234 Foo Bar")

	return p, nil
}

func (app *App) NewServer() *Server {
	// setup grpc server
	grpcsv := grpc.NewServer()
	pb.RegisterDataServer(grpcsv, app)

	// setup http server
	httpsv := httpserver.New(app.logger)
	httpsv.RegisterDataServer(app)

	return &Server{
		grpcsv: grpcsv,
		httpsv: httpsv,
		logger: logger,
	}
}
