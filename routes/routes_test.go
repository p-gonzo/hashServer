package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

//TestHashHandlerFunc
func TestHashHandlerFunc(t *testing.T) {

	//create a separate goroutine for 10 concurrent POST requests
	for i := 0; i < 10; i++ {
		go func() {
			form := url.Values{}
			form.Add("password", "angryMonkey")

			// Create a request to pass to our handler with the form value
			req, err := http.NewRequest("POST", "/hash", strings.NewReader(form.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			if err != nil {
				t.Fatal(err)
			}
			//Create a ResponseRecorder and define our handler function
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(HashHandlerFunc)

			//call ServeHTTP method and pass in our Request and ResponseRecorder.
			handler.ServeHTTP(rr, req)
			// Check the status code is what we expect.
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			// Check the response body is what we expect.
			expected := `ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		}()
	}
	//Provide time for each ~5s request goroutine to complete
	time.Sleep(time.Second * 6)
}

func TestStatsHandlerFunc(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatsHandlerFunc)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the JSON response is what we expect.
	var expected float64
	expected = 10
	var data map[string]interface{}
	err2 := json.Unmarshal([]byte(rr.Body.String()), &data)
	if err2 != nil {
		panic(err2)
	}
	//access the 'total' property on our JSON object
	value := data["total"].(float64)
	if value != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			value, expected)
	}

}
