package models

type User struct {
	ID int
	FName string
	LName string
}

// we're creating a variable of type users that is a slice ([]) holding pointers (*) to User objects
var (
	users []*User
	nextID = 1
)
