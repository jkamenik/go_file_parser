package main

type UnprocessedTrafficTable []UnprocessedTraffic

func NewUnprocessedTrafficTable(header []string, rows [][]string) (table UnprocessedTrafficTable, err error) {
  var traffic *UnprocessedTraffic
  table = make(UnprocessedTrafficTable, len(rows))

  for _, row := range(rows) {
    traffic, err = NewUnprocessedTraffic(header, row)
    table = append(table, *traffic)
  }

  return
}
