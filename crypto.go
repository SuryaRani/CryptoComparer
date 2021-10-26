package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Orders []struct {
	BaseSymbol  string `json:"baseSymbol"`
	QuoteSymbol string `json:"quoteSymbol"`
	OrderBooks  []struct {
		Exchange  string `json:"exchange"`
		OrderBook struct {
			Asks []struct {
				Price    string `json:"price"`
				Quantity string `json:"quantity"`
			} `json:"asks"`
			Bids []struct {
				Price    string `json:"price"`
				Quantity string `json:"quantity"`
			} `json:"bids"`
		} `json:"orderBook"`
	} `json:"orderBooks"`
}

func main() {
	//key = cf6e7628-edb0-4a17-9ff8-3e77daadc266

	/*resp, err := http.Get("https://api.blockchain.com/v3/exchange/l3/BTC-USD")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)



	//bids and asks on blockhains website


	/*fmt.Println("Size: " + strconv.Itoa(len(buys.Buys)))
	fmt.Println("Buy price:", buys.Buys[0].Price)
	fmt.Println("Buy price: ", buys.Buys[0].Quantity)
	fmt.Println("Buy price: " + strconv.Itoa(buys.Buys[0].Number))
	fmt.Println("SEll price:", sells.Sells[0].Price*/

	//buy1, _ = strconv.ParseFloat(buyString1, 32)
	//sell1, _ = strconv.ParseFloat(sellString1, 32)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://dev-api.shrimpy.io/v1/orderbooks?exchange=bittrex&baseSymbol=BTC&quoteSymbol=USD&limit=10")

	if err != nil {
		log.Fatalln(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var exchange1 Orders

	/*var buy1 float32
	var sell1 float32
	var buy2 float32
	var sell2 float32
	*/
	json.Unmarshal(body, &exchange1)

	sb := string(body)
	fmt.Println(sb)

	//this is because for some reason sometimes i dont get the asks or bids array to be filled
	// so i have to do it in a while loop
	for len(exchange1[0].OrderBooks[0].OrderBook.Asks) <= 0 || len(exchange1[0].OrderBooks[0].OrderBook.Bids) <= 0 {

		resp, err := http.Get("https://dev-api.shrimpy.io/v1/orderbooks?exchange=bittrex&baseSymbol=BTC&quoteSymbol=USD&limit=10")

		if err != nil {
			log.Fatalln(err)
		}

		body, _ = ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &exchange1)

	}

	buyString1 := exchange1[0].OrderBooks[0].OrderBook.Asks[0]
	sellString1 := exchange1[0].OrderBooks[0].OrderBook.Bids[0]

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
