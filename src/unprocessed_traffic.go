package main

import (
  "fmt"
)

type UnprocessedTraffic struct {
    SourceIpAddress uint32
    // source_port integer,
    // destination_ip_address numeric(38,0),
    // destination_port integer,
    // file_id integer,
    // input_bytes bigint,
    // output_bytes bigint,
    // start_time timestamp without time zone,
    // end_time timestamp without time zone,
    // net_protocol character varying(64),
    // tcp_flags character varying(64),
    // input_packets bigint,
    // output_packets bigint,

    headers []string
}

func NewUnprocessedTraffic(headers []string, row []string) (*UnprocessedTraffic, error) {
  var key, value string
  var i int
  var err error

  up := UnprocessedTraffic{headers: headers}

  for i, key = range headers {
    value = row[i]
    fmt.Println(i, key, value)

    err = up.Set(key, value)
  }

  return &up, err
}

func (self *UnprocessedTraffic) String() string {
  return fmt.Sprintf("{sa: %d}", self.SourceIpAddress)
}

func (self *UnprocessedTraffic) Set(key string, value interface{}) error {
  // There must be a better way
  switch key {
    case "sa":
      self.SourceIpAddress = value.(uint32)
    default:
      return fmt.Errorf("Unknown column %s", key)
  }

  return nil
}
