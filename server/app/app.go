package app

import (
	"context"
	"log"
	_ "math/rand"

	"github.com/golang/protobuf/proto"
	pb "github.com/yowcow/grpc-http-backdoor/service"
)

var (
	_ pb.DataServer = (*App)(nil)
)

type App struct {
	logger *log.Logger
}

func New(logger *log.Logger) *App {
	return &App{logger}
}

func (app *App) GetPerson(ctx context.Context, v *pb.Void) (*pb.Person, error) {
	app.logger.Println("rpc GetPerson() called")

	//// Randomly return not found error
	//if n := rand.Intn(4); n == 0 {
	//	return nil, errors.New("This time person was not found")
	//}

	p := new(pb.Person)
	p.Id = proto.Int32(123)
	p.Name = proto.String("Hoge Fuga")
	p.Address = proto.String("234 Foo Bar")

	return p, nil
}
