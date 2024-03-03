package main

import (
	"log"

	"github.com/nahkar-golang/todo"
)

func main() {
	srv := new(todo.Server)
	err := srv.Run("8085")

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
