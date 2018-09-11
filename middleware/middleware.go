package middleware

import (
	"errors"
	"log"
	"net/http"
	"time"
)

//RecoveryMiddleware is used to recover from an HTTP panic
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() {
			r := recover()
			if r != nil {
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				//In a production environment
				//We may log the error here
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

//TimeLogMiddleware logs the time it takes to complete a request
//TESTING ONLY, NOT IMPLEMENTED IN PRODUCTION
func TimeLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() //timestamp the beginning of the hashing function
		next.ServeHTTP(w, r)
		t := time.Now()
		elapsed := t.Sub(start) // calculate elapsed time
		log.Println("Time to serve request: " + elapsed.String())
	})
}
