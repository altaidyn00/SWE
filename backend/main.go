package main

import (
	"github.com/whym9/hospital/internal/server"
)

func main() {
	addr := "localhost:8080"

	server.StartServer(addr)
}
