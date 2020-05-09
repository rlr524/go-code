package main

import (
	"github.com/rlr524/go-code/controllers"
	"net/http"
)

// ListenAndServe takes two params, our address (can just use the port for localhost) and the ServeMux (the multiplexor being used)
// which is usually the default, expressed as nil
func main() {
	controllers.RegisterControllers()
	_ = http.ListenAndServe(":3000", nil)
}