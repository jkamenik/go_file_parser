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

type CsvLine struct {
  Header []string
  Line   []string
}

type Line CsvLine

const Buffer = 20


func main() {
  fp, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  defer fp.Close()

  xz  := xzReader(fp)
  csv := csv.NewReader(xz)

  lines, done := streamFile(csv)
  convertedLines, done := convertLine(lines,done)
  done = printStream(convertedLines,done)

  // halt until I am told we are done
  <- done
  fmt.Println("Done")
  os.Exit(0)
}

func streamFile(csv *csv.Reader) (lines chan CsvLine, done chan bool) {
  lines = make(chan CsvLine, Buffer)
  done  = make(chan bool)

  go func(){
    // get Header
    header, err := csv.Read()
    if err != nil {
      done <- true
      return
    }

    for {
      line, err := csv.Read()

      if len(line) > 0 {
        lines <- CsvLine{Header: header, Line: line}
      }

      if err != nil {
        done <- true
        return
      }
    }
  }()

  return
}

func convertLine(csvLines chan CsvLine,finished chan bool) (lines chan Line, done chan bool) {
  lines = make(chan Line, Buffer)
  done = make(chan bool)

  process := func(processor int){
    for {
      select {
        case line := <- csvLines:
          // simulate a IO conversion
          time.Sleep(1 * time.Millisecond)
          fmt.Printf("%d: %s\n", processor, line)
          lines <- Line(line)
        case <- finished:
          done <- true
          return
      }
    }
  }

  go func(){
    var ws sync.WaitGroup
    for i := 0; i < Buffer; i++ {
      fmt.Println(i)
      ws.Add(1)
      go process(i)
    }

    ws.Wait()
    done <- true
  }()

  return
}

func printStream(lines chan Line,finished chan bool) (done chan bool) {
  done = make(chan bool)

  go func(){
    for {
      select {
        case line := <- lines:
          fmt.Println(line)
        case <- finished:
          done <- true
          return
      }
    }
  }()

  return
}
