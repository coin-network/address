package main

// see http://safecurves.cr.yp.to/
// see https://github.com/vsergeev/btckeygenie
// see https://godoc.org/github.com/decred/dcrd/dcrec/secp256k1

import (
  "fmt"
  "github.com/coin-network/curve"
)

func main() {

  KoblitzCurve := curve.S256() // see https://godoc.org/github.com/btcsuite/btcd/btcec#S256
  fmt.Println(KoblitzCurve)
}
