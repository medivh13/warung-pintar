package infra_helper

import (
	"math"
	dtoCheckout "warung-pintar/item-service/src/app/dtos/items"
	"warung-pintar/item-service/src/infra/models"
)

func CheckPromo(dataReq *dtoCheckout.CheckoutReqDTO, items []*models.Items) (float64, error) {
	// example SKU
	// 	mac    = "43N23P"
	// 	google = "120P90"
	// 	alexa  = "A304SD"
	// raspbery = "234234"
	var total float64 = 0
	getItems := make(map[string]int64)

	for _, val := range dataReq.Data {
		if _, ok := getItems[val.Sku]; !ok {
			getItems[val.Sku] = val.Qty
		}
	}

	for _, item := range items {
		if _, ok := getItems[item.SKU]; ok {
			if item.SKU == "A304SD" {
				if getItems[item.SKU] >= 3 {
					total += (item.Price * float64(getItems[item.SKU])) - (item.Price * float64(getItems[item.SKU]) * 10 / 100)
				} else {
					total += (item.Price * float64(getItems[item.SKU]))
				}
			} else if item.SKU == "120P90" && getItems[item.SKU] >= 1 {
				for i := 1; i <= int(getItems[item.SKU]); i++ {
					if i%3 != 0 {
						total += item.Price
					}
				}
			} else if item.SKU == "43N23P" {
				_, ok := getItems["234234"]
				if ok {
					total += (item.Price * float64(getItems[item.SKU])) + ((math.Abs(float64(getItems["234234"]) - float64(getItems[item.SKU]))) * 30)
				} else {
					total += (item.Price * float64(getItems[item.SKU]))
				}
			} else if item.SKU == "234234" {
				if _, ok := getItems["43N23P"]; !ok {
					total += (item.Price * float64(getItems[item.SKU]))
				}
			}
		}
	}

	return total, nil
}
