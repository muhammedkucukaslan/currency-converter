package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

//* The API I used for this project does not send "400 Bad Request" or "402 Payment Required"
//* status code in case user inputs are invalid symbol types(e.g "XYZ" or "T13"). Instead, they send
//* the error in the response body with "200 OK" status code :)

//* If i tried to handle them together, I would had two options.
//* 1- Adding `Error` field to CurrencyData. (I wrote them with comment lines)
//* 2- Implemeting two response type. `SuccessCurrencyData` and `ErrorCurrencyData`
//* Option 2 would be better. Maybe it is a really correct way to handle both body in error and success case.
//* However, i avoided both options against the API to save the sake of REST Principles.

//* Of course, i had to validate the inputs. So, i used "https://api.apilayer.com/fixer/symbols" to check that which
//* variables are useable and saved their data with .csv format(symbols.csv in repo). Then, i read them to map because
//* using map for validation is easier than slices. You can go `validation.go` to see the parsing algorithm.
//* I preferred  csv to json since csv is faster for the data i needed.

//* Check https://apilayer.com/marketplace/fixer-api for api doc details

type CurrencyData struct {
	Success bool    `json:"success"`
	Result  float64 `json:"result"`
	Info    struct {
		Rate float64 `json:"rate"`
	} `json:"info"`

	//* The addition I mentioned above
	//* Error struct {
	//*	Info string `json:"info"`
	//* } `json:"error"
}

func fetchCurrencyAmount(base string, target string, amount float64) (CurrencyData, error) {
	url := fmt.Sprintf("https://api.apilayer.com/fixer/convert?from=%s&to=%s&amount=%f", strings.ToUpper(base), strings.ToUpper(target), amount)

	//* since we want to manipulate req.Header
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", os.Getenv("API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		return CurrencyData{}, fmt.Errorf("Client Request Error: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CurrencyData{}, fmt.Errorf("Invalid authentication credentials")
	}

	var data CurrencyData
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return CurrencyData{}, fmt.Errorf("Reading Body Error: %s", err)
	}

	return data, nil
}
