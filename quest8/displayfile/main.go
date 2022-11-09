package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) == 1 {
		file, err := ioutil.ReadFile("quest8.txt")
		if err != nil {
			err.Error()
		}
		fmt.Printf("%s", file)

	} else {
		fmt.Println("Too many arguments")
	}
}
