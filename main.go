package main

import (
  "fmt"
  "math/big" // using for power
  "flag"
)

var plaintext = flag.Int64("text", 123, "Input plaintext")
var rangeN = flag.Int64("range", 100, "Input range of random numbers")

// only encryption
var flagEncrypt = flag.Int64("encrypt", -1, "Encrypt")
// only decryption 
var flagDecrypt = flag.Int64("decrypt", -1, "Decrypt")

func main() {
  flag.Parse()
  var N, E, D int64

  switch {
  case *flagEncrypt > 0:
    E = 65537
    N = 3085280579
    encrypt := new(big.Int).Exp(big.NewInt(*flagEncrypt), big.NewInt(E), big.NewInt(N))
    fmt.Printf("encrypt : %d\n", encrypt.Int64())
  case *flagDecrypt > 0:
    D = 16382177
    N = 3085280579
    decrypt := new(big.Int).Exp(big.NewInt(*flagDecrypt), big.NewInt(D), big.NewInt(N))
    fmt.Printf("decrypt : %d\n", decrypt.Int64())
  default:
    fmt.Printf("plaintext is %d\n", *plaintext)
    fmt.Printf("range of prime number is %d\n", *rangeN)

    N, E, D = generateKey(*rangeN)
    fmt.Printf("PublicKey(E, N) = (%d, %d)\n", E, N)
    fmt.Printf("SecretrKey(D, N) = (%d, %d)\n", D, N)

    encrypt := new(big.Int).Exp(big.NewInt(*plaintext), big.NewInt(E), big.NewInt(N))
    fmt.Printf("encrypt : %d\n", encrypt.Int64())

    decrypt := new(big.Int).Exp(encrypt, big.NewInt(D), big.NewInt(N))
    fmt.Printf("decrypt : %d\n", decrypt.Int64())
  }
}
