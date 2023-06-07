package service

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"example/pkg/config"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetDigiflazzEnv(cmd string) (username string, sign string) {
	DigiflazzApikey := config.GetEnv("DIGIFLAZZ_APIKEY")
	username = config.GetEnv("DIGIFLAZZ_USERNAME")
	sign = fmt.Sprintf("%x", md5.Sum([]byte(username+DigiflazzApikey+cmd)))
	return username, sign
}

type balanceCheckRequest struct {
	Cmd      string `json:"cmd"`
	Username string `json:"username"`
	Sign     string `json:"sign"`
}

type BalanceCheckResponse struct {
	Data struct {
		Deposit int `json:"deposit"`
	} `json:"data"`
}

func HitBalanceSCheck() (result BalanceCheckResponse, err error) {
	username, sign := GetDigiflazzEnv("depo")
	url := "https://api.digiflazz.com/v1/cek-saldo"
	var balanceCheckRequest balanceCheckRequest
	var balanceCheckResponse BalanceCheckResponse

	balanceCheckRequest.Cmd = "deposit"
	balanceCheckRequest.Username = username
	balanceCheckRequest.Sign = sign
	marshalled, err := json.Marshal(balanceCheckRequest)

	req, err := http.Post(url, "application/json", bytes.NewReader(marshalled))
	if err != nil {
		return BalanceCheckResponse{}, err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &balanceCheckResponse)

	result = balanceCheckResponse
	return result, err
}

type depositRequest struct {
	Cmd       string `json:"cmd"`
	Username  string `json:"username"`
	Amount    int    `json:"amount"`
	Bank      string `json:"bank"`
	OwnerName string `json:"ownerName"`
	Sign      string `json:"sign"`
}

type DepositResponse struct {
	Data struct {
		RC     string `json:"rc"`
		Amount int    `json:"amount"`
		Notes  string `json:"notes"`
	} `json:"data"`
}

func HitDeposit() (result DepositResponse, err error) {
	username, sign := GetDigiflazzEnv("deposit")
	url := "https://api.digiflazz.com/v1/deposit"

	var depositRequest depositRequest
	var depositResponse DepositResponse

	depositRequest.Cmd = "deposit"
	depositRequest.Username = username
	depositRequest.Sign = sign
	depositRequest.Bank = "BCA"
	depositRequest.OwnerName = "TEGAR SUBHAN FAUZI"
	depositRequest.Amount = 200000
	marshalled, err := json.Marshal(depositRequest)

	req, err := http.Post(url, "application/json", bytes.NewReader(marshalled))
	if err != nil {
		return DepositResponse{}, err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &depositResponse)

	result = depositResponse
	return result, err
}

type priceListRequest struct {
	Cmd      string `json:"cmd"`
	Username string `json:"username"`
	Code     string `json:"code"`
	Sign     string `json:"sign"`
}

type PriceListResponse struct {
	ProductName         string `json:"product_name"`
	Category            string `json:"category"`
	Brand               string `json:"brand"`
	Type                string `json:"type,omitempty"`
	SellerName          string `json:"seller_name"`
	Admin               int    `json:"admin,omitempty"`
	Commission          int    `json:"commission,omitempty"`
	Price               int    `json:"price,omitempty"`
	BuyerSkuCode        string `json:"buyer_sku_code"`
	BuyerProductStatus  bool   `json:"buyer_product_status"`
	SellerProductStatus bool   `json:"seller_product_status"`
	UnlimitedStock      bool   `json:"unlimited_stock,omitempty"`
	Stock               int    `json:"stock,omitempty"`
	Multi               bool   `json:"multi,omitempty"`
	StartCutOff         string `json:"start_cut_off,omitempty"`
	EndCutOff           string `json:"end_cut_off,omitempty"`
	Desc                string `json:"desc"`
}

type WrapPriceListResponse struct {
	Data []PriceListResponse `json:"data"`
}

func HitPriceList(priceType, code string) (result WrapPriceListResponse, err error) {
	if priceType == "" {
		err = config.CustomError{Message: "price type cannot be empty"}
		return WrapPriceListResponse{}, err
	} else if priceType != "prepaid" && priceType != "pasca" {
		err = config.CustomError{Message: "the price type option you selected is not available"}
		return WrapPriceListResponse{}, err
	}

	username, sign := GetDigiflazzEnv("pricelist")
	url := "https://api.digiflazz.com/v1/price-list"

	var priceListRequest priceListRequest
	var wrapPriceListResponse WrapPriceListResponse

	priceListRequest.Cmd = priceType
	priceListRequest.Username = username
	priceListRequest.Sign = sign
	if code != "" {
		priceListRequest.Code = code
	}
	marshalled, err := json.Marshal(priceListRequest)

	req, err := http.Post(url, "application/json", bytes.NewReader(marshalled))
	if err != nil {
		return WrapPriceListResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(req.Body)

	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &wrapPriceListResponse)

	result = wrapPriceListResponse
	return result, err
}

func PriceListGrouped(data []PriceListResponse, groupBy string) (result map[string][]PriceListResponse, err error) {
	groupedData := make(map[string][]PriceListResponse)

	if groupBy == "category" {
		for _, d := range data {
			replacedStr := strings.ReplaceAll(d.Category, " ", "_")
			key := strings.ToLower(replacedStr)
			value := d
			groupedData[key] = append(groupedData[key], value)
		}
	} else if groupBy == "brand" {
		for _, d := range data {
			replacedStr := strings.ReplaceAll(d.Brand, " ", "_")
			key := strings.ToLower(replacedStr)
			value := d
			groupedData[key] = append(groupedData[key], value)
		}
	} else {
		err = config.CustomError{Message: "the group by option you selected is not available"}
	}

	result = groupedData
	return result, err
}
