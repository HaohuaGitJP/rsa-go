/* RSA暗号の生成 (https://qiita.com/YutaKase6/items/cd9e26d723809dc85928)を参考 */
package main

import (
  "math/rand"
  "time"
  "fmt"
)

// Check if it is prime numbers 
func checkPrime(num int64) bool {
  var flagPrime bool
  switch num {
  case 0:
    return false
  case 1:
    return false
  case 2:
    return true
  default:
    flagPrime = true
  }
  var i int64
  for i = 2; i < num ; i = i + 1 {
    if num % i == 0 {
      flagPrime = false
      break
    }
  }
  return flagPrime
}

// Generate Prime numbers
func generatePrime(n int64) (int64, int64) {
  t := time.Now().Unix()
  rand.Seed(t)
  var p, q int64
  p = rand.Int63n(n)
  q = rand.Int63n(n)

  // p, q == Prime number && p != q
  for !checkPrime(p) || !checkPrime(q) {
    for p = rand.Int63n(n); !checkPrime(p) || p < 10; p = rand.Int63n(n) {}
    for q = rand.Int63n(n); !checkPrime(q) || q < 10; q = rand.Int63n(n) {}
    if p == q {
      for p == q {
        for q = rand.Int63n(n); !checkPrime(q) || q < 10; q = rand.Int63n(n) {}
      }
    }
  }
  return p, q
}

// generateCommonDevide
func generateCommonDevide(a int64, b int64) int64 {
  if b == 0{
    return a
  }  
  return generateCommonDevide(b, a % b)
}

// generateCommonMult
func generateCommonMult(a int64, b int64) int64 {
  return a * b / generateCommonDevide(a, b)
}

// generateKey
func generateKey(n int64) (int64, int64, int64) {
  p, q := generatePrime(n)
  var N, E, D int64
  N = p * q
  L := generateCommonMult((p - 1), (q - 1))
  for E = rand.Int63n(n); generateCommonDevide(E, L) != 1; E = rand.Int63n(n) {}
  for D = 2; (E * D) % L != 1 || E == D; D = D + 1 {}
  fmt.Printf("p = %d, q = %d\nN = %d, L = %d\nE = %d, D = %d\n", p, q, N, L, E, D)
  return N, E, D
}
