package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

type Candle struct {
	Date   int32
	High   float32
	Low    float32
	Open   float64
	Close  float64
	Change int
}

var candles []Candle

func read() {
	b, err := ioutil.ReadFile("../datastore/BTC_XMR_1512086400_1512087400_")

	err = json.Unmarshal(b, &candles)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func main() {
	read()
	fmt.Println(candles[0])
	Change := math.Dim(candles[0].Open, candles[0].Close)
	fmt.Println(Change)
}
