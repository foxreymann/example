package main

import "fmt"

type Token struct {
  name string
  value int
  owner string
}

func changeOwner(token Token,newOwner string) Token {
  newToken := Token{}
  newToken.name = token.name
  newToken.value = token.value
  newToken.owner = newOwner
  return newToken
}

type Contract struct {
}

func (c* Contract) transfer(token Token,newOwner string) Token {
  return changeOwner(token, newOwner)
}

func main() {
  token := Token{"ecToken", 100, "fox"}
	fmt.Println(token)

  contract := Contract{}
  token = contract.transfer(token, "Paul")
	fmt.Println(token)
}
