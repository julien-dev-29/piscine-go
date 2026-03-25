package main

import (
	"log"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		data, err := os.ReadFile(os.Args[i])
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.Write(data)
	}
}
