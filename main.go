package main

import "fmt"
import "github.com/foxreymann/example/ledger"

func main() {
  token := ledger.Token{"ecToken", 100, "fox"}
	fmt.Println(token)

  contract := ledger.Contract{}
  token = contract.Transfer(token, "Paul")
	fmt.Println(token)
}
