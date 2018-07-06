package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/yowcow/test/grpc-test/server/httpserver"
	"google.golang.org/grpc"
)

type Server struct {
	grpcsv     *grpc.Server
	httpsv     *httpserver.Server
	httpsvDone chan struct{}
	logger     *log.Logger
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
