package main

import "os"

func main() {
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "01" || os.Args[i] == "galaxy" || os.Args[i] == "galaxy 01" {
			os.Stdout.WriteString("Alert!!!\n")
			return
		}
	}
}
