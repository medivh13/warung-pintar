package infra_helper

import (
	dtoCheckout "warung-pintar/checkout-service/src/app/dtos/checkout"
)

func CheckPromo(dataReq *dtoCheckout.CheckoutReqDTO, items []*dtoCheckout.ItemDTO, promos []*dtoCheckout.PromoDTO) (float64, error) {
	// example SKU
	// 	mac    = "43N23P"
	// 	google = "120P90"
	// 	alexa  = "A304SD"
	// raspbery = "234234"
	var total float64 = 0
	getItems := make(map[string]int64)
	getPromosPercentage := make(map[string]float64)
	getPromosMinQty := make(map[string]int64)

	for _, val := range dataReq.Data {
		if _, ok := getItems[val.Sku]; !ok {
			getItems[val.Sku] = val.Qty
		}
	}

	for _, val := range promos {
		if _, ok := getPromosPercentage[val.SKU]; !ok {
			getPromosPercentage[val.SKU] = val.ValuePercentage
		}
		if _, ok := getPromosMinQty[val.SKU]; !ok {
			getPromosMinQty[val.SKU] = val.MinimumQty
		}
	}

	for _, item := range items {
		if _, ok := getItems[item.SKU]; ok {
			if _, ok := getPromosMinQty[item.SKU]; ok {
				if getItems[item.SKU] >= getPromosMinQty[item.SKU] {
					total += (item.Price * float64(getItems[item.SKU])) - (item.Price * float64(getItems[item.SKU]) * getPromosPercentage[item.SKU])
				} else {
					total += (item.Price * float64(getItems[item.SKU]))
				}
			} else {
				total += (item.Price * float64(getItems[item.SKU]))
			}
		}
	}

	return total, nil
}
