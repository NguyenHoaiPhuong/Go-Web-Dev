package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/033_CryptoCurrency/config"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/033_CryptoCurrency/models"
)

const (
	// CMCEndpoint : get price of BTC, ETH, TOKO
	CMCEndpoint string = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?symbol=ETH,BTC,TOKO,USDT"
)

func getAllMarketsPrice(conf *config.Config) {
	request, err := http.NewRequest("GET", CMCEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("X-CMC_PRO_API_KEY", conf.CMCApiKey)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data *models.Response
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
	fmt.Printf("Status: %+v\n", data.Status)
	fmt.Printf("Data: %+v\n", data.Data)
	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("BTC Data: %+v\n", data.Data.BTC)
	fmt.Printf("BTC Quote: %+v\n", data.Data.BTC.Quote.USD)

	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("ETH Data: %+v\n", data.Data.ETH)
	fmt.Printf("BTC Quote: %+v\n", data.Data.ETH.Quote.USD)

	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("USDT Data: %+v\n", data.Data.USDT)
	fmt.Printf("BTC Quote: %+v\n", data.Data.USDT.Quote.USD)

	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("TOKO Data: %+v\n", data.Data.TOKO)
	fmt.Printf("BTC Quote: %+v\n", data.Data.TOKO.Quote.USD)
}

func main() {
	// get config
	conf := config.GetConfig()

	getAllMarketsPrice(conf)
}
