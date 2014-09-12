package main

import (
  // "io"
  "log"
  "os"
  "fmt"
  "encoding/csv"
)

type Chunk struct {
  Header []string
  Lines  [][]string
}

func main() {
  fp, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  defer fp.Close()

  xz  := xzReader(fp)
  csv := csv.NewReader(xz)

  chunks, done := chunkFile(csv)
  done = readChunks(chunks, done)

  select {
    case <- done:
      fmt.Println("Exiting")
  }
}

func chunkFile(csv *csv.Reader) (chunks chan Chunk, done chan bool) {
  chunkSize := 1000
  chunks = make(chan Chunk, 10)
  done   = make(chan bool)

  go func(){
    header, err := csv.Read()
    if err != nil {
      done <- true
      return
    }

    lines := [][]string{}

    for {
      line, err := csv.Read()
      if err != nil {
        break
      }
      lines = append(lines, line)

      if len(lines) >= chunkSize {
        chunks <- Chunk{Header: header, Lines: lines}
        lines = [][]string{}
      }
    }

    if len(lines) >= 0 {
      chunks <- Chunk{Header: header, Lines: lines}
    }
    done <- true
  }()

  return
}

func readChunks(chunks chan Chunk, chunkerDone chan bool) (done chan bool) {
  done = make(chan bool)

  go func(){
    for {
      select {
        case chunk := <-chunks:
          fmt.Println(chunk)
        case _ = <-chunkerDone:
          fmt.Println("Done")
          done <- true
          return
      }
    }
  }()

  return
}
