package main

import (
	"errors"
	"fmt"

	"github.com/lilosir/LeetCode/kit"
)

// ExecuteAllMiddlewares middleware
func ExecuteAllMiddlewares(funcs []func() error) error {

	for len(funcs) > 0 {
		err := funcs[0]()
		if err != nil {
			return err
		}
		funcs = funcs[1:]
	}
	return nil
}

func add() error {
	fmt.Println("1+2=3")
	return nil
}
func minus() error {
	fmt.Println("5-3=2")
	return nil
}
func multiply() error {
	fmt.Println("4*4=13")
	return errors.New("multiply error")
}
func dive() error {
	fmt.Println("12/2=6")
	return nil
}

func mapTest() map[string]int {
	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["2"] = 3

	return m
}

func main() {
	// result := ExecuteAllMiddlewares([]func() error{add, minus, multiply, dive})
	a := []int{1, 10, 5, 6, 7, 8}
	b := []int{10, 1, 6, 5, 8, 7}
	aInterface := make([]interface{}, len(a))
	for _, v := range a {
		aInterface = append(aInterface, v)
	}
	bInterface := make([]interface{}, len(b))
	for _, v := range b {
		bInterface = append(bInterface, v)
	}
	fmt.Println(kit.CheckTwoSliceHaveSameIntElements(aInterface, bInterface))
}
