package main

import (
  // "io"
  "log"
  "os"
  "fmt"
  "encoding/csv"
)

func main() {
  fp, err := os.Open("test_data/test.csv.xz")
  if err != nil {
    log.Fatal(err)
  }
  defer fp.Close()

  xz  := xzReader(fp)
  csv := csv.NewReader(xz)

  header, err := csv.Read()
  fmt.Println(header)

  lines, err := csv.ReadAll()
  fmt.Println(NewUnprocessedTrafficTable(header, lines))
}
