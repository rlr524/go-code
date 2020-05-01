package main

import (
	"github.com/rlr524/go-code/controllers"
	"log"
	"net/http"
)

//// we're importing our User struct via the models package; we need to use the fully qualified name on the import here
//// and then we can use dot notation, e.g. packageName.methodBeingCalled to get our struct then use it to add some state
//// the _ is called the "write only" variable and allows us to dump information into it but not use it; it's often used the
//// other way from where we're using it here, e.g. we'd print the port (using p as the var name) and not use the err
//func main() {
//	port := 3000
//	_, err := startWebServer(port, 5)
//	fmt.Println(err)
//	u := models.User{
//		ID:    1,
//		FName: "Madison",
//		LName: "Hu",
//	}
//	fmt.Println("Hello, " + u.FName)
//}
//
//func startWebServer(port, numberOfRetries int) (int, error) {
//	fmt.Println("Server is starting...")
//	// stuff to do
//	fmt.Println("Server has started on port", port, "...")
//	fmt.Println("Number of retries", numberOfRetries)
//	return port, nil
//}

// ListenAndServe takes two params, our address (can just use the port for localhost) and the ServeMux (the multiplexor being used)
// which is usually the default, expressed as nil
func main() {
	controllers.RegisterControllers()
	log.Fatal(http.ListenAndServe(":3000", nil))
}