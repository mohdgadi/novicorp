package main
import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"./item"
)

func main(){
	fmt.Println("hello world")
	getPricings()
}

func getPricings() {
	data, err := ioutil.ReadFile("./items.json")
	var items []item.Item
	err = json.Unmarshal(data, &items)
	fmt.Println(err, items)
}

func calculate(itemNames []string) (total float64) {
	itemNamesQuantityMap := make(map [string]int)
	for _,name := range itemNames {
		itemNamesQuantityMap[name]++
	}

	return
}