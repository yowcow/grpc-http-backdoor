package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	pb "github.com/yowcow/test/grpc-test/service"
	"google.golang.org/grpc"
)

var address string
var logger *log.Logger

func init() {
	flag.StringVar(&address, "address", "127.0.0.1:9999", "server address")
	flag.Parse()
}

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewDataClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	person, err := client.GetPerson(ctx, &pb.Void{})
	if err != nil {
		log.Fatalln(err)
	}
	logger.Println(person)
	cancel()
}
