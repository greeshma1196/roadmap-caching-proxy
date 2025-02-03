package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type product struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
}

type products struct {
	Products []product `json:"products"`
}

func main() {

	_ = flag.Int("port", 8080, "an int")
	urlPtr := flag.String("origin", "http://dummyjson.com", "a string")
	flag.Parse()

	urlStr := *urlPtr

	path := parseURL(urlStr)

	if path == "/products" {
		prods := fetchProductsDataHandler()
		fmt.Printf("Products: %+v\n", prods)
	}

}

func parseURL(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return u.Path
}

func fetchProductsDataHandler() products {
	var products products
	url := "http://dummyjson.com/products"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &products)
	if err != nil {
		log.Fatal(err.Error())
	}
	return products
}
