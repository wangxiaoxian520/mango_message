package main

import (
	"log"
	"os"
)

func main() {
	//message 调用
	//go run main.go message.go boltdb.go 0 tcp://localhost:45454
	//go run main.go message.go boltdb.go 1 tcp://localhost:45454
	if len(os.Args) <= 2 {
		log.Printf("Usage: %s 0|1 <url>\n", os.Args[0])
	} else {
		node = os.Args[1]
		runNode(os.Args[2])
	}

	//pub sub 调用
	//go run main.go pubsub.go boltdb.go server tcp://localhost:45454
	// if len(os.Args) > 2 && os.Args[1] == "server" {
	// 	server(os.Args[2])
	// 	os.Exit(0)
	// }
	// if len(os.Args) > 3 && os.Args[1] == "client" {
	// 	client(os.Args[2], os.Args[3])
	// 	os.Exit(0)
	// }
	// fmt.Fprintf(os.Stderr, "Usage: pubsub server|client <URL> <ARG>\n")
	// os.Exit(1)
}
