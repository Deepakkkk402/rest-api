package main

import "fmt"

func Run() error {
	fmt.Println("strarting the application")
	return nil
}

func main() {
	fmt.Println()
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
