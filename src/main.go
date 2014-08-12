package main

import (
  "io"
  "log"
  "os"
)


func main() {
	fp, err := os.Open("test_data/test.csv.xz")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	r := xzReader(fp)

	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Printf("copied %d bytes with err: %v", n, err)
	} else {
		log.Printf("copied %d bytes", n)
	}
}
