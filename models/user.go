package models

import (
	"errors"
	"fmt"
)

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

func GetUsers() []*User {
	return users
}
// create a function named AddUser; it uses a reference u to the User struct and returns params of a User object and
// nil or a blank User object and an error message if the user ID is not zero
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include an id or it must be set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}

func UpdateUser (u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", u.ID)
}

func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			// performing a splice on the slice; taking a slice of the users up to but not including i and append
			// to it all the users from the index after the match and on, so replacing the i with the object after
			// i and all objects there after in js this would be like slice.splice(i, 1) but in js with splice all
			// we are doing is dropping the element at the index place specified
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID '%v' not found", id)
}