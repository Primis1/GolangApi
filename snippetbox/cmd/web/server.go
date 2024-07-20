package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

func echoInUpperCase(w io.Writer, r io.Reader) {

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Fprintf(w, "%s\n", strings.ToUpper(line))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error in %v: %v", scanner.Scan(), err)
	}


	type User struct {
		ID int `json:"oiadfgdfgdkbvfk"`
	}
}