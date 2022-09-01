package main

import (
	"fmt"
	"foo/pkg/model"
)

type Customer struct {
	Name string
	Age  int
}

func main() {
	model.HelloFoo()
	foo := model.Foo{Name: "John"}
	fmt.Println(foo)
}
