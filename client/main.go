package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/yowcow/test/grpc-test/service"
	"google.golang.org/grpc"
)

var grpcAddress string
var httpAddress string
var logger *log.Logger

func init() {
	flag.StringVar(&grpcAddress, "grpc-address", "127.0.0.1:9999", "grpc server address")
	flag.StringVar(&httpAddress, "http-address", "127.0.0.1:9998", "http server address")
	flag.Parse()
}

func init() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	runGRPC()
	runHTTP()
}

func runGRPC() {
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
	if err != nil {
		logger.Fatalln(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := pb.NewDataClient(conn)
	person, err := client.GetPerson(ctx, &pb.Void{})
	if err != nil {
		log.Println("grpc GetPerson():", err)
		return
	}

	logger.Println("grpc GetPerson():", person)
}

func runHTTP() {
	uri := &url.URL{
		Scheme: "http",
		Host:   httpAddress,
		Path:   "/GetPerson/",
	}

	resp, err := http.Get(uri.String())
	if err != nil {
		logger.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		logger.Println("http GetPerson(): Not found")
		return
	} else if resp.StatusCode != http.StatusOK {
		logger.Println("http GetPerson():", resp.StatusCode)
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Fatalln(err)
	}

	person := new(pb.Person)
	if err := proto.Unmarshal(b, person); err != nil {
		logger.Fatalln(err)
	}

	logger.Println("http GetPerson():", person)
}
