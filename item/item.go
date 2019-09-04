package item

import (
	"io/ioutil"
	"encoding/json"
) 

// Item denotes the entity Item having a name and different pricing stratergy.
type Item struct {
	Name     string    `json:"name"`
	Pricings []Pricing `json:"pricings"`
}

// Pricing of an Item.
type Pricing struct {
	// Minimum Quantity of an Item for an applicable offer.
	MinimumQuantity int `json:"minimumQuantity"`

	// If this is true then the price is applied to all the items.
	AppliedToAll bool `json:"appliedToAll"`

	// Price of the combo.
	Price float64 `json:"price"`
}

// CalcuateCost for individual Item.
func (item Item) CalcuateCost(quantity int) (cost float64) {
	for _, pricing := range item.Pricings {
		var offerCost float64
		quantity, offerCost = pricing.Apply(quantity)
		cost += offerCost
		if quantity == 0 {
			return
		}
	}

	return
}

// Apply a pricing stratergy to a specific Item.
func (pricing Pricing) Apply(quantity int) (remainingQuantity int, cost float64) {
	for quantity >= pricing.MinimumQuantity {
		if pricing.AppliedToAll {
			cost += float64(quantity) * pricing.Price
			quantity = 0
		} else {
			cost += pricing.Price
			quantity -= pricing.MinimumQuantity
		}
	}

	remainingQuantity = quantity
	return
}

// GetItemsData scans for the items and its prices for the store. 
func GetItemsData() (itemsMap map[string]Item, err error) {
	var data []byte
	if data, err = ioutil.ReadFile("./items.json"); err != nil {
		return
	}

	var items []Item
	if err = json.Unmarshal(data, &items); err != nil {
		return
	}

	itemsMap = make(map[string]Item)
	for _, it := range items {
		itemsMap[it.Name] = it
	}
	return
}