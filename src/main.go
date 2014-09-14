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

func convertLine(csvLines chan *CsvLine) (lines chan *NetflowLine) {
  lines = make(chan *NetflowLine, Buffer)

  go func(){
    var netflowLine *NetflowLine

    for line := range csvLines {
      netflowLine, _ = NewNetflowLine(line)
      lines <- netflowLine
    }
    close(lines)
  }()

  return
}

func printStream(lines chan *NetflowLine) (done chan int) {
  done = make(chan int)

  go func(){
    i := 0
    for line := range lines {
      fmt.Println(line)
      i++
      // fmt.Println(i)
    }
    done <- i
  }()

  return
}
