package main

import (
	"log"
	"net"

	"github.com/bmf-san/golang-grpc-example/pkg/proto/user"
	"github.com/bmf-san/golang-grpc-example/pkg/service"

	grpc "google.golang.org/grpc"
)

func main() {
	var p net.Listener
	var err error
	if p, err = net.Listen("tcp", ":19003"); err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	userService := &service.UserService{}
	user.RegisterUserServer(s, userService)
	s.Serve(p)
}
