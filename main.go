package main

import (
  "fmt"
  "math/big" // べき乗の計算で使用(桁あふれに備えて)
  "flag"
)

var plaintext = flag.Int64("text", 123, "Input plaintext")
var rangeN = flag.Int64("range", 100, "Input range of random numbers")
var flagEncrypt = flag.Int64("en", -1, "Encryption")
var flagDecrypt = flag.Int64("de", -1, "Decryption")
var inputP = flag.Int64("p", -1, "input pirme number P")
var inputQ = flag.Int64("q", -1, "input pirme number Q")

func main() {
  flag.Parse()
  /*
  fmt.Printf("plaintext is %d\n", *plaintext)
  fmt.Printf("range of prime number is %d\n", *rangeN)
  */

  // Generate Keys
  var N, E, D int64
  N, E, D = generateKey(*rangeN, *inputP, *inputQ)
  fmt.Printf("PublicKey(E, N) = (%d, %d)\n", E, N)
  fmt.Printf("SecretrKey(D, N) = (%d, %d)\n", D, N)

  // Encryption
  encrypt := new(big.Int).Exp(big.NewInt(*plaintext), big.NewInt(E), big.NewInt(N))
  fmt.Printf("encrypt : %d\n", encrypt.Int64())

  // Decryption
  decrypt := new(big.Int).Exp(encrypt, big.NewInt(D), big.NewInt(N))
  fmt.Printf("decrypt : %d\n", decrypt.Int64())
}
