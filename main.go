package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"regexp"
)

const homeURL = "https://www.comperdelivery.com.br"
const itemsRegex = "<h3 class=\"shelf-item__title\">.*\n(.*)\n.*</h3>"
const itemURLRegex = "href=\"(.*)\" "
const df = "?sc=3"

type outputData map[string]interface{}


func getContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func getItemURLs(data *[]outputData) {
	var URL string
	homeRegex, _ := regexp.Compile(itemsRegex)
	URLRegex, _ := regexp.Compile(itemURLRegex)
	homeContent := getContent(homeURL)
	items := homeRegex.FindAllString(homeContent, -1)

	for _, item := range items {
		URL = URLRegex.FindStringSubmatch(item)[1] + df
		*data = append(*data, outputData{"url": URL})
	}
}



func main() {
	var data []outputData
	getItemURLs(&data)
	fmt.Println(data)
}