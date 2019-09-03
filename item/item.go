package item

type Item struct {
	Name string `json:"name"`
	Pricings []Pricing `json:"pricings"`
}

type Pricing struct{
	MinimumQuantity int `json:"minimumQuantity"`
	AppliedToAll bool `json:"appliedToAll"`
	Price float64 `json:"price"`
}

func(item Item) CalcuateCost(quantity int) (cost float64) {
	for _,pricing:=range item.Pricings{
		var offerCost float64
		quantity,offerCost = pricing.Apply(quantity)
		cost+=offerCost
		if quantity == 0{
			return
		}
	}

	return
}

func(pricing Pricing) Apply(quantity int) (remainingQuantity int, cost float64) {
	for quantity>=pricing.MinimumQuantity {
		if pricing.AppliedToAll{
			cost+= float64(quantity)*pricing.Price
			quantity = 0
		} else {
		cost+= float64(pricing.MinimumQuantity)*pricing.Price
		quantity-=pricing.MinimumQuantity
		}
	}

	remainingQuantity = quantity
	return
}