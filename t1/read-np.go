// Copyright (C) Philip Schlump, 2015-2016.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//
// On MacOS/OSX - to create
//
// mkfifo bob
//

func main() {

	// file, err := os.OpenFile("bob", os.O_RDONLY, os.ModeNamedPipe)
	file, err := os.OpenFile("bob", os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	lim := 200

	// loop limited to 200
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("%s\n", line)
			if strings.HasPrefix(string(line), "quit") {
				os.Exit(0)
			}
		}
		lim--
		if lim < 0 {
			fmt.Printf("limit reached\n")
			os.Exit(0)
		}
	}

}
