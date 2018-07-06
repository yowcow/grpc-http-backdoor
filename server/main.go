package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var grpcNetwork string
var grpcAddress string
var httpNetwork string
var httpAddress string
var logger *log.Logger

func init() {
	flag.StringVar(&grpcNetwork, "grpc-network", "tcp", "network to listen for grpc")
	flag.StringVar(&grpcAddress, "grpc-address", ":9999", "address to listen for grpc")
	flag.StringVar(&httpNetwork, "http-network", "tcp", "network to listen for http")
	flag.StringVar(&httpAddress, "http-address", ":9998", "address to listen for http")
	flag.Parse()
}

func init() {
	logger = log.New(os.Stdout, fmt.Sprintf("[%d] ", os.Getpid()), log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	grpcln, err := net.Listen(grpcNetwork, grpcAddress)
	if err != nil {
		logger.Fatalln(err)
	}

	httpln, err := net.Listen(httpNetwork, httpAddress)
	if err != nil {
		logger.Fatalln(err)
	}

	app := &App{logger}
	sv := app.NewServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		<-c
		logger.Println("Stopping server gracefully")
		sv.GracefulStop()
	}()

	logger.Println("Server starting")
	sv.Serve(grpcln, httpln)
	logger.Println("Server stopped")
}
