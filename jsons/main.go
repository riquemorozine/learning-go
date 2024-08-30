package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number int     `json:"n"` // "n" is the name that will be used in the JSON
	Wallet float64 `json:"_"` // "_" means that the field will be ignored
}

func main() {
	account := Account{
		Number: 1,
		Wallet: 100.0,
	}

	res, err := json.Marshal(account) // Encoding the JSON
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(account) // Encoding the JSON

	if err != nil {
		panic(err)
	}

	defaultJson := []byte(`{"n": 2, "w": 200.0}`) // "w" is not a field of the struct

	var account2 Account
	err = json.Unmarshal(defaultJson, &account2) // Decoding the JSON

	if err != nil {
		panic(err)
	}

	println(account2.Number)
}
