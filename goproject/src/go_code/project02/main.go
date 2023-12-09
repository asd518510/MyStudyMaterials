package main

import (
	"fmt"

	"github.com/blockcypher/gobcy"
)

func main() {
	btc := gobcy.API{"YOURTOKEN", "btc", "main"}
	txs, err := btc.GetChain()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", txs)
}
