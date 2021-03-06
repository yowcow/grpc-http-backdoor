package server

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/yowcow/grpc-http-backdoor/server/httpserver"
	pb "github.com/yowcow/grpc-http-backdoor/service"
	"google.golang.org/grpc"
)

type Server struct {
	grpcsv     *grpc.Server
	httpsv     *httpserver.Server
	httpsvDone chan struct{}
	logger     *log.Logger
}

func New(app pb.DataServer, logger *log.Logger) *Server {
	grpcsv := grpc.NewServer()
	pb.RegisterDataServer(grpcsv, app)

	httpsv := httpserver.New(logger)
	httpsv.RegisterDataServer(app)

	return &Server{
		grpcsv: grpcsv,
		httpsv: httpsv,
		logger: logger,
	}
}

func (sv *Server) Serve(grpcln net.Listener, httpln net.Listener) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		sv.logger.Println("Started grpc server")
		if err := sv.grpcsv.Serve(grpcln); err != nil {
			sv.logger.Println(err)
		}
		sv.logger.Println("Stopped grpc server")
	}()

	sv.httpsvDone = make(chan struct{})
	go func() {
		defer wg.Done()
		sv.logger.Println("Started http server")
		if err := sv.httpsv.Serve(httpln); err != nil {
			sv.logger.Println(err)
		}
		<-sv.httpsvDone
		sv.logger.Println("Stopped http server")
	}()

	wg.Wait()
}

func (sv *Server) GracefulStop() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		sv.logger.Println("Stopping grpc server gracefully")
		sv.grpcsv.GracefulStop()
	}()

	go func() {
		defer wg.Done()
		sv.logger.Println("Stopping http server gracefully")
		if err := sv.httpsv.GracefulStop(context.Background()); err != nil {
			panic(err)
		}
		close(sv.httpsvDone)
	}()

	wg.Wait()
}
