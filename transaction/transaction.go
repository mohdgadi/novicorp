package transaction

import (
	"fmt"
	"novicorp/item"
)

type Transaction struct {
	// itemList contains the data of Items and possible pricing stratergy.
	itemList     map[string]item.Item
	shoppingList map[string]int
}

// New creates a new transaction.
func New(itemList map[string]item.Item) (transaction Transaction) {
	return Transaction{
		itemList:     itemList,
		shoppingList: make(map[string]int),
	}
}

// Scan each Item in the shopping list until -1 is encountered.
func (transaction Transaction) ScanShoppingist() (err error) {
	fmt.Println("Enter items:")
	for {
		var itemName string
		fmt.Scanln(&itemName)
		if itemName == "-1" {
			return
		}

		var (
			ok   bool
			item item.Item
		)

		if item, ok = transaction.itemList[itemName]; !ok {
			err = fmt.Errorf("invalid item %s", itemName)
			return
		}

		transaction.shoppingList[item.Name]++
	}
}

// CalculateCost of the transaction i.e all the Items.
func (trasaction Transaction) CalculateCost() (cost float64, err error) {
	for itemName, quantity := range trasaction.shoppingList {
		var item item.Item
		item, _ = trasaction.itemList[itemName]
		cost += item.CalcuateCost(quantity)
	}

	return
}
