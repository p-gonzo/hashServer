package util

import (
	"crypto/sha512"
	"encoding/base64"
	"time"

	"github.com/p-gonzo/hashServer/state"
)

//EncryptPassword Helper function to encrypt a password
func EncryptPassword(password string) string {

	//Create a timestamp to determine how long the hashing function takes to complete
	start := time.Now()

	//Hash the password
	h := sha512.New()
	h.Write([]byte(password))
	b := h.Sum(nil)
	hashedPass := base64.StdEncoding.EncodeToString(b)

	//Calculate the elapsed time to complete the complete the hashing
	//And add the hashing statistics to the application state
	t := time.Now()
	elapsed := t.Sub(start)
	go state.SubmitHashTimeAndUpdateAverage(elapsed.Nanoseconds() / int64(time.Microsecond))

	return hashedPass
}
