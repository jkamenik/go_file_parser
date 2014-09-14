package main

import (
  "fmt"
)

type NetflowLine struct {
  ts string
  te string
  pr string
  sa string
  da string
  sp string
  dp string
  ibyt string
  obyt string
}

type NetflowTable []*NetflowLine


func NewNetflowLine(csv *CsvLine) (*NetflowLine,error) {
  self := NetflowLine{}

  // TODO: Find a better way
  self.ts = csv.Get("ts")
  self.te = csv.Get("te")
  self.pr = csv.Get("pr")
  self.sa = csv.Get("sa")
  self.da = csv.Get("da")
  self.sp = csv.Get("sp")
  self.dp = csv.Get("dp")
  self.ibyt = csv.Get("ibyt")
  self.obyt = csv.Get("obyt")

  return &self, nil
}

func (self *NetflowTable)Send() {
  // code to send to the database here.
  fmt.Printf("----\nSending %d lines\n%s", len(*self), *self)
}
