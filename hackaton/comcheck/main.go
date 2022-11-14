package main

import "os"

func main() {
	args := os.Args[1:]
	p := 0
	for _, v := range args {
		if v == "galaxy" || v == "01" || v == "galaxy 01" {
			p = 1
		}
	}
	if p != 0 {
		output := []byte("Alert!!!\n")
		os.Stdout.Write(output)
	}
}
