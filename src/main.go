package main

import (
  // "io"
  "log"
  "os"
  "fmt"
  "encoding/csv"
  // "time"
  // "sync"
)


type Line CsvLine

const Buffer = 20000


func main() {
  fp, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  defer fp.Close()

  xz  := xzReader(fp)
  csv := csv.NewReader(xz)

  lines := streamCsv(csv, Buffer)
  convertedLines := convertLine(lines)
  completed := printStream(convertedLines)

  // halt until I am told we are done
  x := <- completed
  fmt.Printf("Done %d lines\n", x)
  os.Exit(0)
}

func convertLine(csvLines chan CsvLine) (lines chan Line) {
  lines = make(chan Line, Buffer)

  go func(){
    for line := range csvLines {
      // TODO: Put conversion here
      lines <- Line(line)
    }
    close(lines)
  }()

  return
}

func printStream(lines chan Line) (done chan int) {
  done = make(chan int)

  go func(){
    i := 0
    for _ = range lines {
      // fmt.Println(line)
      i++
      // fmt.Println(i)
    }
    done <- i
  }()

  return
}
