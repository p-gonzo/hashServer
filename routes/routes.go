package routes

import (
	"encoding/json"
	"hashServer/state"
	"hashServer/util"
	"io"
	"net/http"
	"time"
)

//the stats struct is used to send a JSON object back to the client on the '/stats' endpoint
type stats struct {
	Total   int `json:"total"`
	Average int `json:"average"`
}

//StatsHandlerFunc is the function for handling GET requests to the '/stats' endpoint
func StatsHandlerFunc(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		a, b := state.GetHashStats()
		w.Header().Set("Content-Type", "application/json")
		s := stats{Total: a, Average: b}
		json.NewEncoder(w).Encode(&s)
	} else {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404: Not Found")
	}
}

//HashHandlerFunc is the function for handling POST requests to the '/hash' endpoint
func HashHandlerFunc(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := req.ParseForm()
		if err != nil {
			panic(err)
		}
		password := req.Form.Get("password")

		//time.sleep is okay to use here
		//the HTTP handler is run in a separate goroutine for each request so it is non-blocking
		//This is verified in the TestHashHandlerFunc in ./routes_test.go
		time.Sleep(5 * time.Second)
		io.WriteString(w, util.EncryptPassword(password))
	} else {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "404: Not Found")
	}
}
