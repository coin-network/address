package main

// see http://safecurves.cr.yp.to/
// see https://github.com/vsergeev/btckeygenie
// see https://godoc.org/github.com/decred/dcrd/dcrec/secp256k1

import (
  "github.com/coin-network/curve"
  "crypto/rand"
  "crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
)

type PrivateKey ecdsa.PrivateKey

func main() {

  KoblitzCurve := curve.S256() // see https://godoc.org/github.com/btcsuite/btcd/btcec#S256

  privkey, err := curve.NewPrivateKey(KoblitzCurve)

  if err != nil {
		panic("Error")
	}

  fmt.Println(" Private Key ")
  fmt.Println(privkey.D)

  fmt.Println(" Public Key ")
  pubkey := (privkey.PublicKey)
  fmt.Println(pubkey.X)
  fmt.Println(pubkey.Y)

  fmt.Println("-------")
}
