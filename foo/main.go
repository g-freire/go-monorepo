package main

import (
	"bar/pkg/model"
	"fmt"
)

func main() {
	model.HelloBar()
	customer := model.Customer{Name: "Lennon", Age: 21}
	fmt.Println(customer)

}
