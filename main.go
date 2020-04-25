package main

import (
	"fmt"
	"github.com/rlr524/go-code/models"
)
// we're importing our User struct via the models package; we need to use the fully qualified name on the import here
// and then we can use dot notation, e.g. packageName.methodBeingCalled to get our struct then use it to add some state
func main() {
	u := models.User{
		ID:    2,
		FName: "Madison",
		LName: "Hu",
	}
	fmt.Println(u)
}
