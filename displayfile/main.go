package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		os.Stdout.WriteString("File name missing")
	} else if len(os.Args) > 2 {
		os.Stdout.WriteString("Too many arguments")
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)
}
