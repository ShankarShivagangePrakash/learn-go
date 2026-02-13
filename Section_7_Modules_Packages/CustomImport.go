package main

// before executing this code, run go mod tidy to add the dependency to go.mod file
import (
	"fmt"
	// importing a custom package from github
	"github.com/pborman/uuid"
)

func main () {

	uuid := uuid.New()

	fmt.Println(uuid)
}