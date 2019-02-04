package ledger

type Token struct {
  Name string
  Value int
  Owner string
}

func changeOwner(token Token,newOwner string) Token {
  newToken := Token{}
  newToken.Name = token.Name
  newToken.Value = token.Value
  newToken.Owner = newOwner
  return newToken
}

type Contract struct {
}

func (c* Contract) Transfer(token Token,newOwner string) Token {
  return changeOwner(token, newOwner)
}
