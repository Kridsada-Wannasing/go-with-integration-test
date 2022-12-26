package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{ID: 1, Name: "kridsada", Age: 26}

	b, err := json.Marshal(u)
	fmt.Printf("type : %T \n", b)
	fmt.Printf("byte : %s \n", b)
	fmt.Println(err)
}
