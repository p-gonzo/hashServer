package server

import (
	"io"
	"net/http"

	"github.com/p-gonzo/hashServer/middleware"

	"github.com/p-gonzo/hashServer/routes"
)

//CreateHTTPServerWithRoutes returns a reference so caller can invoke methods on the server instance
func CreateHTTPServerWithRoutes() *http.Server {

	srv := &http.Server{Addr: ":8080"} // pointer to instance

	hashHandler := http.HandlerFunc(routes.HashHandlerFunc)
	statsHandler := http.HandlerFunc(routes.StatsHandlerFunc)

	http.Handle("/hash", middleware.RecoveryMiddleware((hashHandler)))
	http.Handle("/stats", middleware.RecoveryMiddleware((statsHandler)))

	//a GET request to the '/shutdown' route will invoke a graceful shutdown
	//without interrupting any active connections
	//https://golang.org/pkg/net/http/#Server.Shutdown
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Shutdown Request Received ")
		go func() {
			srv.Shutdown(nil)
		}()
	})

	return srv
}
