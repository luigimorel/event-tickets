package models

import "testing"

func CheckForValidation(t *testing.T) {
	u := &User{}

	err := u.ValidateJSON()

	if err != nil {
		t.Fatal(err)
	}
}
