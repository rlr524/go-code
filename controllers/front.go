package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// note in go route handling /users is not the same as /users/, we need to provide both patterns in order to ensure that
// our userController is used in either case
func RegisterControllers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	_ = enc.Encode(data)
}
