package main

import (
	"flag"
	"log"
)

var(
    port int
    endpoint string
)

func main() {
	flag.IntVar(&port, "p", 8080, "Port on which server will start. Defaults to 8080")
	flag.StringVar(&endpoint, "e", "/", "Endpoint on which server will listen, preceded by '/'. Defaults to '/'")
	srv := Server{port: port, endpoint: endpoint}
	srv.Start()
	log.Println("Listening on http://127.0.0.1:" + endpoint)
}