package main

import (
  //"github.com/btcsuite/btcd/btcec"
  "crypto/rand"
  "crypto/ecdsa"
	"crypto/elliptic"
	//"crypto/sha256"
	//"encoding/base64"
	"fmt"
  "bytes"
)

func GeneratePrivateKey(pubkeyCurve elliptic.Curve) (*ecdsa.PrivateKey, error) {
  //key := new(ecdsa.PrivateKey)
	key, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return (*ecdsa.PrivateKey)(key), nil
}

func GetPublicKey(privatekey *ecdsa.PrivateKey) ecdsa.PublicKey{
  return (privatekey.PublicKey)
}

func main() {

 	pubkeyCurve := elliptic.P256() //see http://golang.org/pkg/crypto/elliptic/#P256
  //pubkeyCurve := btcec.S256() // see https://godoc.org/github.com/btcsuite/btcd/btcec#S256

  privkey, err := GeneratePrivateKey(pubkeyCurve)

  if err != nil {
		panic("Error")
	}

  fmt.Println(" Private Key ")
  fmt.Println(privkey.D)

  d := privkey.D.Bytes()

  /* Relleno con 0x00 */
  padded_d := append(bytes.Repeat([]byte{0x00}, 32-len(d)), d...)

  fmt.Println("padded_d:")
  fmt.Println(padded_d)
  fmt.Println("-------")

  fmt.Println(" Public Key ")
  pubkey := GetPublicKey(privkey)
  fmt.Println(pubkey.X)
  fmt.Println(pubkey.Y)


  fmt.Println("-------")

}
