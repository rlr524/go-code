package controllers

import (
	"encoding/json"
	"github.com/rlr524/go-code/models"
	"net/http"
	"regexp"
	"strconv"
)

// create a custom type named userController that will be a blank struct for now
type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
		} else {
			matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
			if len(matches) == 0 {
				w.WriteHeader(http.StatusNotFound)
			}
			id, err := strconv.Atoi(matches[1])
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}
			switch r.Method {
			case http.MethodGet:
				uc.get(id, w)
			case http.MethodPut:
				uc.put(id, w, r)
			case http.MethodDelete:
				uc.delete(id, w)
			default:
				w.WriteHeader(http.StatusNotImplemented)
			}
		}
	}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not parse user object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not parse user object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}




// in func ServeHTTP
// we're using a func to create a method (class of function associated with a data type) named ServeHTTP to
// bind the function to our userController type
// we create a local variable, by convention naming it uc then the type we're going to bind to
// note the use of the uc local var vs using a "this" keyword; in go, we avoid "this" as it doesn't tell you
// anything as we work through the method body our ServeHTTP method takes two params, the http.ResponseWriter
// method (names it as w) and a pointer named as r to the
// http.Request object that are both part of the net/http package
// remember a byte slice []byte is just an alias for a string, like char in C, it is a pointer to the first
// element in an array of chars (a string)

// in func newUserController we're creating a constructor function here to control how the regular
// expression in our userController
// is going to be defined (go doesn't support classic oop constructors as used in oop languages with classes);
// by convention, we name all
// constructor functions with new followed by the name of the type of object we're constructing
// we're going to return a pointer to a userController object (use pointers instead of returning as it helps avoid
// an unnecessary copy; we're then providing our implementation which is our regex pattern definition applied to our
// userIDPattern field; we use a string literal for our regex definition which is looking for paths that are
// /users followed by / then a number
// because we're using a struct type here, we can immediately take the address of it for our return (& operator)
// go is recognizing we are using the address of a local variable and it will automatically promote that variable up to
// where it needs to be in order to avoid having it overwritten in memory once the function leaves scope