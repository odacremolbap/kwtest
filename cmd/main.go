package main

import (
	"fmt"
	"log"

	"github.com/odacremolbap/kwtest/pkg/kwtest"
)

type options struct {
	port int
}

func main() {

	// TODO parse params
	op := options{
		port: 80,
	}

	fmt.Printf("Starting kwtest")
	config := kwtest.Config{
		Port: op.port,
	}

	kwt, err := kwtest.NewServer(config)
	if err != nil {
		log.Fatalf("Error creating server: %s", err.Error())
	}
	kwt.Run()

}
