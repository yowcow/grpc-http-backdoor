package httpserver

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/yowcow/grpc-http-backdoor/server/app"
	pb "github.com/yowcow/grpc-http-backdoor/service"
)

func TestHello(t *testing.T) {
	logbuf := new(bytes.Buffer)
	logger := log.New(logbuf, "", 0)
	s := New(logger)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	ready := make(chan struct{})
	finish := make(chan struct{})
	done := make(chan struct{})

	go func() {
		close(ready)
		if err := s.Serve(ln); err != http.ErrServerClosed {
			t.Fatal("expected nil but got", err)
		}
		<-finish
		close(done)
	}()

	<-ready // server is ready

	resp, err := http.Get(fmt.Sprintf("http://%s/hello/", ln.Addr().String()))
	if err != nil {
		t.Fatal("expected nil but got", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Error("expected 200 but got", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	s.GracefulStop(context.Background())
	close(finish)
	<-done
	ln.Close()

	if "Hello world" != string(body) {
		t.Error("expected 'Hello world' but got", string(body))
	}
}

func TestGetPerson(t *testing.T) {
	logbuf := new(bytes.Buffer)
	logger := log.New(logbuf, "", 0)
	s := New(logger)
	s.RegisterDataServer(app.New(logger))

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	ready := make(chan struct{})
	finish := make(chan struct{})
	done := make(chan struct{})

	go func() {
		close(ready)
		if err := s.Serve(ln); err != http.ErrServerClosed {
			t.Fatal("expected nil but got", err)
		}
		<-finish
		close(done)
	}()

	<-ready // server is ready

	resp, err := http.Get(fmt.Sprintf("http://%s/GetPerson/", ln.Addr().String()))
	if err != nil {
		t.Fatal("expected nil but got", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("expected nil but got", err)
	}

	s.GracefulStop(context.Background())
	close(finish)
	<-done
	ln.Close()

	p := new(pb.Person)
	if err := proto.Unmarshal(body, p); err != nil {
		t.Error("expected nil but got", err)
	}
	if p.GetId() != 123 {
		t.Error("expected 123 but got", p.GetId())
	}
	if p.GetName() != "Hoge Fuga" {
		t.Error("expected 'Hoge Fuga' but got", p.GetName())
	}
	if p.GetAddress() != "234 Foo Bar" {
		t.Error("expected '234 Foo Bar' but got", p.GetAddress())
	}
}
