package server

import (
	"hashServer/middleware"
	"hashServer/routes"
	"io"
	"net/http"
)

//CreateHTTPServerWithRoutes returns a reference so caller can invoke methods on the server instance
func CreateHTTPServerWithRoutes() *http.Server {

	srv := &http.Server{Addr: ":8080"} // pointer to instance

	hashHandler := http.HandlerFunc(routes.HashHandlerFunc)
	statsHandler := http.HandlerFunc(routes.StatsHandlerFunc)

	http.Handle("/hash", middleware.RecoveryMiddleware((hashHandler)))
	http.Handle("/stats", middleware.RecoveryMiddleware((statsHandler)))

	//a GET request to the '/shutdown' route will invoke a graceful shutdown
	//without interupting any active connections
	//https://golang.org/pkg/net/http/#Server.Shutdown
	http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Shutdown Request Received ")
		go func() {
			srv.Shutdown(nil)
		}()
	})

	return srv
}
