// JSON与struct之间的转化
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Roles    []Role `json:"roles"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// struct -> json
	roles := []Role{
		{1, "ADMIN"},
		{2, "MERCHANT"},
	}
	b, err := json.Marshal(&User{ID: 1, Username: "张三", Age: 18, Roles: roles})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b): %v\n", string(b))
	// json -> struct
	jsonStr := `{"id":1,"username":"张三","age":18,"roles":[{"id":1,"name":"ADMIN"},{"id":2,"name":"MERCHANT"}]}`
	u := User{}
	err = json.Unmarshal([]byte(jsonStr), &u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("u: %v\n", u)
}
