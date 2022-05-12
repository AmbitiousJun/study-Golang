package main

import "fmt"

func main() {
	fmt.Println("start....")
	defer fmt.Println("s1")
	defer fmt.Println("s2")
	defer fmt.Println("s3")
	fmt.Println("end....")
}
