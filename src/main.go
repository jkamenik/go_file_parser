package main

import (
	// "io"
	"log"
	"os"
	"fmt"
)


func main() {
	done := make(chan bool)

	fp, err := os.Open("test_data/test.csv.xz")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	r := xzReader(fp)
	log.Print(r)
	s := stringGenerator(r, done)
	log.Print(s)

	loop(done, s)

	println("after select")
}

func loop(done chan bool, input <- chan string) {
	for {
		select {
			case line := <- input:
				fmt.Printf("line: %s\n", line)
			case <- done:
				return
		}
	}
}
