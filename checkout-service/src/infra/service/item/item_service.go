package item_services

/*
 * Author      : jodyalmaida (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : checkout-service
 */
import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	dto "warung-pintar/checkout-service/src/app/dtos/checkout"

	_ "github.com/joho/godotenv/autoload"
)

type ItemServices interface {
	ListItems(sku []string) ([]*dto.ItemDTO, error)
	ListPromos(sku []string) ([]*dto.PromoDTO, error)
}
type itemService struct {
}

func NewItemService() ItemServices {
	return &itemService{}
}

func (s *itemService) ListItems(sku []string) ([]*dto.ItemDTO, error) {

	itemListUrl := os.Getenv("ITEM_LIST_URL")
	var responseData *dto.ItemsRespDTO

	var dataItemReqDTO []*dto.DataItemReqDTO

	for _, val := range sku {
		dataSKU := &dto.DataItemReqDTO{
			Sku: val,
		}
		dataItemReqDTO = append(dataItemReqDTO, dataSKU)
	}

	reqDTO := &dto.ItemReqDTO{
		Data: dataItemReqDTO,
	}

	body, _ := json.Marshal(reqDTO)
	req, err := http.NewRequest("POST", itemListUrl, bytes.NewBuffer(body))
	req.Header["x-api-key"] = []string{"warung-pintar"}
	req.Header["Content-Type"] = []string{"application/json"}
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to create http request List Item: %v", err)
	}

	client := &http.Client{
		Timeout: time.Second * 50,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err, "here error")
		return nil, err
	}
	log.Println("Success execute  : ", itemListUrl)

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("failed decode get items response: %v", err)
	}

	if len(responseData.Data) < 1 {
		return nil, errors.New("data not found")
	}

	return responseData.Data, nil
}

func (s *itemService) ListPromos(sku []string) ([]*dto.PromoDTO, error) {

	itemListUrl := os.Getenv("PROMO_LIST_URL")
	var responseData *dto.PromoRespDTO

	var dataItemReqDTO []*dto.DataItemReqDTO

	for _, val := range sku {
		dataSKU := &dto.DataItemReqDTO{
			Sku: val,
		}
		dataItemReqDTO = append(dataItemReqDTO, dataSKU)
	}

	reqDTO := &dto.ItemReqDTO{
		Data: dataItemReqDTO,
	}

	body, _ := json.Marshal(reqDTO)
	req, err := http.NewRequest("POST", itemListUrl, bytes.NewBuffer(body))
	req.Header["x-api-key"] = []string{"warung-pintar"}
	req.Header["Content-Type"] = []string{"application/json"}
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to create http request Promo Item: %v", err)
	}

	client := &http.Client{
		Timeout: time.Second * 50,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err, "here error")
		return nil, err
	}
	log.Println("Success execute  : ", itemListUrl)

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("failed decode get promo response: %v", err)
	}

	if len(responseData.Data) < 1 {
		return nil, errors.New("data not found")
	}

	return responseData.Data, nil
}
