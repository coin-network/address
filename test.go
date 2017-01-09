// Copyright 2017 The coin-network developers
// License: MIT

package main

import (
  "fmt"
  "github.com/coin-network/curve"
)

func main() {

  KoblitzCurve := curve.S256() // see https://godoc.org/github.com/btcsuite/btcd/btcec#S256
  fmt.Println(KoblitzCurve)
}
