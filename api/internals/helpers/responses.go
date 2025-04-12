package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if err != nil {
		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Printf("Error encoding response: %v\n", err)
		}
	}
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("Error encoding response: %v\n", err)
	}
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("content-type header is not application/json")
	}

	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return nil
}
