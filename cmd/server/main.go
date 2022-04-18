package main

import (
	"log"
)

func main() {
	srv := server.newHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
