package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// Reads the entire body of the HTTP request
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		// If the read is successful without errors

		// Parses the read JSON content into the object pointed to by x
		if err := json.Unmarshal([]byte(body), x); err != nil {
			// If an error occurs parsing the JSON, exit the function
			return
		}
	}
	// If an error occurs reading the request body, or parsing the JSON,
	// or if the request body is empty, then no further action is taken.
	// This function is complete.
}
