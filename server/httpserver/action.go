package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	pb "github.com/yowcow/test/grpc-test/service"
)

func (s *Server) GetHello(w http.ResponseWriter, req *http.Request) {
	s.logger.Println("http GetHello() called")

	time.Sleep(10 * time.Second)
	fmt.Fprintln(w, "Hello world")
}

func (s *Server) GetPerson(w http.ResponseWriter, req *http.Request) {
	s.logger.Println("http GetPerson() called")

	h := w.Header()
	h.Set("Content-Type", "application/vnd.google.protobuf")

	ctx := context.Background()
	p, err := s.app.GetPerson(ctx, new(pb.Void))
	if err != nil {
		s.logger.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := proto.Marshal(p)
	if err != nil {
		s.logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.Set("Content-Length", strconv.Itoa(len(b)))
	w.Write(b)
}
