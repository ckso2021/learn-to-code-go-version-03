package main

import (
	"fmt"
	puppy "github.com/ckso2021/go-mod-puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)
}
