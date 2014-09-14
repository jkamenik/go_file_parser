package main

import (
  "encoding/csv"
  "fmt"
  "strings"
)

type CsvLine struct {
  Header []string
  Line   []string
}

// streamCsv
//  Streams a CSV Reader into a returned channel.  Each CSV row is streamed along with the header.
//  "true" is sent to the `done` channel when the file is finished.
//
// Args
//  csv    - The csv.Reader that will be read from.
//  buffer - The "lines" buffer factor.  Send "0" for an unbuffered channel.
func streamCsv(csv *csv.Reader, buffer int) (lines chan *CsvLine) {
  lines = make(chan *CsvLine, buffer)

  go func(){
    // get Header
    header, err := csv.Read()
    if err != nil {
      close(lines)
      return
    }

    i := 0

    for {
      line, err := csv.Read()

      if len(line) > 0 {
        i++
        lines <- &CsvLine{Header: header, Line: line}
      }

      if err != nil {
        fmt.Printf("Sent %d lines\n",i)
        close(lines)
        return
      }
    }
  }()

  return
}

func (self *CsvLine)Get(key string) (value string){
  x := -1
  for i, value := range self.Header {
    if value == key {
      x = i
      break
    }
  }

  if x == -1 {
    return ""
  }

  return strings.TrimSpace(self.Line[x])
}
