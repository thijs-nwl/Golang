package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var urlData = [4]string{"BTC", "XMR", "1512086400", "1512088400"}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func buildURL(first string, sec string, start string, end string) (string, string) {
	return fmt.Sprintf("https://poloniex.com/public?command=returnChartData&currencyPair=%v_%v&start=%v&end=%v&period=300", first, sec, start, end),
		fmt.Sprintf("../datastore/%v_%v_%v_%v_", first, sec, start, end)
}

func main() {
	url, _ := buildURL(urlData[0], urlData[1], urlData[2], urlData[3])
	err := getData(url)
	if err == nil {
		fmt.Printf("succes")
	}
}

func getData(url string) error {
	res, err := http.Get(url)
	check(err)

	body, err := ioutil.ReadAll(res.Body)
	check(err)
	b := []byte(body)

	_, path := buildURL(urlData[0], urlData[1], urlData[2], urlData[3])
	err = ioutil.WriteFile(path, b, 0644)
	check(err)

	return err
}
