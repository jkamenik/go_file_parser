package main

import (
  // "io"
  "log"
  "os"
  "fmt"
  "encoding/csv"
  "time"
  "sync"
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

  process := func(processor int, group sync.WaitGroup){
    defer group.Done()

    for line := range csvLines {
      // simulate a IO conversion
      time.Sleep(1 * time.Millisecond)
      // fmt.Printf("%d: %s\n", processor, line)
      lines <- Line(line)
    }
    fmt.Printf("Processor %d is done\n", processor)
  }

  go func(){
    var ws sync.WaitGroup
    ws.Add(1000)
    for i := 0; i < 1000 ; i++ {
      // ws.Add(1)
      go process(i,ws)
    }

    ws.Wait()
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
