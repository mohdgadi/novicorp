package main

import (
	"fmt"
	"novicorp/item"
	"novicorp/transaction"
)

func main() {
	var (
		itemsMap map[string]item.Item
		err      error
	)

	if itemsMap, err = item.GetItemsData(); err != nil {
		panic(err.Error())
	}

	tr := transaction.New(itemsMap)
	if err = tr.ScanShoppingist(); err != nil {
		panic(err)
	}

	var cost float64
	if cost, err = tr.CalculateCost();err!=nil{
		panic(err)
	}

	fmt.Println("Total Cost for transaction is:", cost)
	return
}
