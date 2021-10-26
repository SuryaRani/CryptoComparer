package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Sells struct {
	Sells []Sell `json:"bids"`
}

//struct to hold buys
type Sell struct {
	Price    float64 `json:"px"`
	Quantity float64 `json:"qty"`
	Number   int     `json:"num"`
}

type Buys struct {
	Buys []Buy `json:"asks"`
}

//struct to hold sells
type Buy struct {
	Price    float64 `json:"px"`
	Quantity float64 `json:"qty"`
	Number   int     `json:"num"`
}

type Binance struct {
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type Format struct {
	Btc1     float64
	Btc2     float64
	BtcSell1 float64
	BtcSell2 float64
	Eth1     float64
	Eth2     float64
	EthSell1 float64
	EthSell2 float64
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://api.blockchain.com/v3/exchange/l3/BTC-USD")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var sells Sells
	var buys Buys

	json.Unmarshal(body, &buys)
	json.Unmarshal(body, &sells)

	blockhainBtcBuy := buys.Buys[0].Price
	blockhainBtcSell := sells.Sells[0].Price

	//fmt.Fprintf(w, "This is buy price from blockchain for bitcoin %f, and sell price %f\n", blockhainBtcBuy, blockhainBtcSell)

	resp, err = http.Get("https://api.blockchain.com/v3/exchange/l3/ETH-USD")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ = ioutil.ReadAll(resp.Body)

	var sellsEth Sells
	var buysEth Buys

	json.Unmarshal(body, &buysEth)
	json.Unmarshal(body, &sellsEth)

	blockchainEthBuy := buysEth.Buys[0].Price
	blockchainEthSell := sellsEth.Sells[0].Price

	//fmt.Fprintf(w, "This is buy price from blockchain for bitcoin %f, and sell price %f\n", blockchainEthBuy, blockchainEthSell)

	//getting bitcoin price from binance
	resp, err = http.Get("https://api.binance.com/api/v3/depth?symbol=BTCUSDT")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ = ioutil.ReadAll(resp.Body)
	var btcBinance Binance

	json.Unmarshal(body, &btcBinance)

	binanceBtcBuy, _ := strconv.ParseFloat(btcBinance.Asks[0][0], 64)
	binanceBtcSell, _ := strconv.ParseFloat(btcBinance.Bids[0][0], 64)
	//fmt.Fprintf(w, "This is buy price from Binance for bitcoin %f, and sell price %f\n", binanceBtcBuy, binanceBtcSell)

	//getting ethereum price from binance
	resp, err = http.Get("https://api.binance.com/api/v3/depth?symbol=ETHUSDT")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ = ioutil.ReadAll(resp.Body)
	var ethBinance Binance

	json.Unmarshal(body, &ethBinance)

	binanceEthBuy, _ := strconv.ParseFloat(ethBinance.Asks[0][0], 64)
	binanceEthSell, _ := strconv.ParseFloat(ethBinance.Bids[0][0], 64)
	//fmt.Fprintf(w, "This is buy price from Binance for etheruem %f, and sell price %f\n", binanceEthBuy, binanceEthSell)

	if blockhainBtcBuy < binanceBtcBuy {
		//fmt.Fprintf(w, "You should buy Bitcoin from Blockchain\n")
	} else {
		//fmt.Fprintf(w, "You should buy Bitcoin from Binance\n")
	}

	if blockchainEthBuy < binanceEthBuy {
		//fmt.Fprintf(w, "You should buy Ethereum from Blockchain\n")
	} else {
		//fmt.Fprintf(w, "You should buy Ethereum from Binance\n")
	}

	if blockhainBtcSell > binanceBtcSell {
		//fmt.Fprintf(w, "You should sell Bitcoin on Blockchain\n")
	} else {
		//fmt.Fprintf(w, "You should sell Bitcoin on Binance\n")
	}

	if blockchainEthSell > binanceEthSell {
		//fmt.Fprintf(w, "You should sell Ethereum on Blockchain\n")
	} else {
		//fmt.Fprintf(w, "You should sell Ethereum on Binance\n")
	}

	templ, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	data := Format{
		Btc1:     blockhainBtcBuy,
		Btc2:     binanceBtcBuy,
		BtcSell1: blockhainBtcSell,
		BtcSell2: binanceBtcSell,
		Eth1:     blockchainEthBuy,
		Eth2:     binanceEthBuy,
		EthSell1: blockchainEthSell,
		EthSell2: binanceEthSell,
	}
	templ.Execute(w, data)

}
