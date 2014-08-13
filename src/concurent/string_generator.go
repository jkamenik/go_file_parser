package main

import (
	"io"
	"bufio"
	"strings"
)

func stringGenerator(r io.Reader, done chan bool) <- chan string {
	buffer := bufio.NewReader(r)
	output := make(chan string)

	go func() {
		var out string
		var stripped string
		var err error

		for {
			out, err = buffer.ReadString('\n')
			stripped = strings.Trim(out, " \n")

			if len(stripped) > 0 {
				output <- stripped
			}

			if err != nil {
				done <- true
			}
		}
	}()

	return output
}
