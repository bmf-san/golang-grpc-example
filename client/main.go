package main

import (
	context "context"
	"fmt"
	"log"

	"github.com/bmf-san/golang-grpc-example/pkg/proto/user"
	grpc "google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	var err error
	if conn, err = grpc.Dial("127.0.0.1:19003", grpc.WithInsecure()); err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	c := user.NewUserClient(conn)
	req := &user.GetUserRequest{
		Type: "admin",
	}
	res, err := c.GetUser(context.TODO(), req)
	fmt.Printf("result:%#v \n", res.Name)
	fmt.Printf("error::%#v \n", err)
}
