package main

import (
	"fmt"
	"log"
	"time"

	"github.com/p-gonzo/hashServer/server"
)

//Application entrypoint
//main will start the server and log to the terminal that it is running
func main() {
	srv := server.CreateHTTPServerWithRoutes()
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Sever listening on http://localhost:8080")
	}()
	log.Fatal(srv.ListenAndServe())
}
