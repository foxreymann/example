package main

import "fmt"

type Token struct {
  name string
  value int
  owner string
}

func changeOwner(token Token,newOwner string) Token {
  var newToken = Token{}
  newToken.name = token.name
  newToken.value = token.value
  newToken.owner = newOwner
  return newToken
}

func main() {
  var token = Token{"ecToken", 100, "fox"}
	fmt.Println(token)
  token = changeOwner(token, "Paul")
	fmt.Println(token)
}
