package main

import (
	"fmt"
	"server/02/server"
)

func main() {
	//Structs can be passed in this way when they need to accept multiple parameters
	opts := []server.Option{
		server.WithAddress("127.0.0.1:8000"),
		server.WithNetwork("tcp"),
		server.WithNetwork("json"),
	}
	s := server.NewServer(opts...)

	fmt.Println("New server", s)
}
