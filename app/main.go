package main

import (
	"fmt"

	"github.com/gordjw/resume-server/server"
)

func main() {
	fmt.Println("Starting server")
	server.NewServer(8090)
}
