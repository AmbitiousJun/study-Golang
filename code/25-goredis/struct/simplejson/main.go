// simplejson库的使用
package main

import (
	"fmt"
	"log"

	simplejson "github.com/bitly/go-simplejson"
)

type User struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age int `json:"age"`
	Roles []Role `json:"roles"`
}

type Role struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

func main() {
	// struct -> json
	j := simplejson.New()
	roles := []Role{
		{1, "ADMIN"},
		{2, "MERCHANT"},
	}
	j.Set("user", &User{ID: 1, Username: "张三", Password: "123456", Age: 18, Roles: roles})
	b, err := j.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("string(b): %v\n", string(b))
	// json -> struct
	jsonStr := `{"user":{"id":2,"username":"李四","password":"123123","age":25,"roles":[{"id":1,"name":"ADMIN"},{"id":2,"name":"MERCHANT"}]}}`
	j2, err := simplejson.NewJson([]byte(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	u := User{}
	j2 = j2.Get("user")
	u.ID = uint(j2.Get("id").MustInt())
	u.Username = j2.Get("username").MustString()
	u.Password = j2.Get("password").MustString()
	u.Age = j2.Get("age").MustInt()
	for _, v := range j2.Get("roles").MustArray() {
		fmt.Printf("v: %v\n", v)
	}
	// if u, ok := j2.Get("user").Interface().(User); ok {
	// 	fmt.Printf("u: %v\n", u)
	// }
}