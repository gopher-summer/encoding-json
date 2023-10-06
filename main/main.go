package main

import (
	"fmt"
	json "github.com/gopher-summer/encoding-json"
)

func main() {

	bytes, err := json.Marshal(User{})
	fmt.Println(string(bytes), err)
}

type User struct {
	Id int
	id int
}
