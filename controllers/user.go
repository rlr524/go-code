package controllers

import (
	"net/http"
	"regexp"
)

// create a custom type named userController that will be a blank struct for now
type userController struct {
	userIDPattern *regexp.Regexp
}

// here we're using a func to create a method (class of function associated with a data type) named ServeHTTP to
// bind the function to our userController type
// we create a local variable, by convention naming it uc then the type we're going to bind to
// note the use of the uc local var vs using a "this" keyword; in go, we avoid "this" as it doesn't tell you
// anything as we work through the method body
// our ServeHTTP method takes two params, the http.ResponseWriter method (names it as w) and a pointer named as r to the
// http.Request object that are both part of the net/http package
// remember a byte slice []byte is just an alias for a string, like char in C, it is a pointer to the first
// element in an array of chars (a string)
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

// we're creating a constructor function here to control how the regular expression in our userController
// is going to be defined (go doesn't support classic oop constructors as used in oop languages with classes); by convention, we name all
// constructor functions with new followed by the name of the type of object we're constructing
// we're going to return a pointer to a userController object (use pointers instead of returning as it helps avoid
// an unnecessary copy; we're then providing our implementation which is our regex pattern definition applied to our
// userIDPattern field; we use a string literal for our regex definition which is looking for paths that are
// /users followed by / then a number
// because we're using a struct type here, we can immediately take the address of it for our return (& operator)
// go is recognizing we are using the address of a local variable and it will automatically promote that variable up to
// where it needs to be in order to avoid having it overwritten in memory once the function leaves scope
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

