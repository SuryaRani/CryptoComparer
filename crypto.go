package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Buys struct {
	Buys []Buy `json:"bids"`
}

//struct to hold buys
type Buy struct {
	Price    float32 `json:"price"`
	Quantity float32 `json:"quantity"`
}

type Sells struct {
	Sells []Sell `json:"asks"`
}

//struct to hold sells
type Sell struct {
	Price    float32 `json:"price"`
	Quantity float32 `json:"quantity"`
}

func main() {
	//key = cf6e7628-edb0-4a17-9ff8-3e77daadc266

	/*resp, err := http.Get("https://api.blockchain.com/v3/exchange/l3/BTC-USD")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)



	//bids and asks on blockhains website


	fmt.Println("Size: " + strconv.Itoa(len(buys.Buys)))
	fmt.Println("Buy price:", buys.Buys[0].Price)
	fmt.Println("Buy price: ", buys.Buys[0].Quantity)
	fmt.Println("Buy price: " + strconv.Itoa(buys.Buys[0].Number))
	fmt.Println("SEll price:", sells.Sells[0].Price)*/

	var sells Sells
	var buys Buys

	resp, err := http.Get("https://dev-api.shrimpy.io/v1/orderbooks?exchange=bittrex&baseSymbol=BTC&limit=10")

	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &buys)
	json.Unmarshal(body, &sells)

	fmt.Println("Size: " + strconv.Itoa(len(buys.Buys)))
	fmt.Println("Buy price:", buys.Buys[0].Price)

}
