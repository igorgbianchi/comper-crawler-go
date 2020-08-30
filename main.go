package main

import (
	"net/http"
	"io/ioutil"
	"regexp"
	"encoding/json"
	"strings"
)

const homeURL = "https://www.comperdelivery.com.br"
const itemsPattern = "<h3 class=\"shelf-item__title\">.*\n(.*)\n.*</h3>"
const itemURLPattern = "href=\"(.*)\" "
const pageErrorPattern = "\"pageUrl\":\"http://www.comperdelivery.com.br/Sistema/404\""
const dataLayerPattern = "skuJson_0 = (.*});"
const df = "?sc=3"

type productInfo struct {
	Name string `json:"name"`
	URL string `json:"url"`
	ImgURL string `json:"img_url"`
	Price float64 `json:"price"`
	Availability bool `json:"availability"`
}

// Get the web page content on string format
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

// List items URL from Comper home page
func listItemURLs() []string {
	var URL string
	var urls []string

	homeRegex, _ := regexp.Compile(itemsPattern)
	URLRegex, _ := regexp.Compile(itemURLPattern)
	homeContent := getContent(homeURL)
	items := homeRegex.FindAllString(homeContent, -1)

	for _, item := range items {
		URL = URLRegex.FindStringSubmatch(item)[1] + df
		urls = append(urls, URL)
	}

	return urls
}

// Format in a legible way the price crawled
func formatPrice(price float64, availability bool) float64{
	if availability{
		return price/100
	}
	return 0
}

// Crawl relevant info for each product passed in a URLs list
func crawlItemData(urls []string) []productInfo{
	var content string
	var dataLayerStr string
	var dataLayerMap map[string]interface{}
	var outputData []productInfo

	errorRegex, _ := regexp.Compile(pageErrorPattern)
	dataLayerRegex, _ := regexp.Compile(dataLayerPattern)
	numItems, itemIdx := 0, 0

	for numItems < 3 {
		content = getContent(urls[itemIdx])
		if !errorRegex.MatchString(content){
			numItems++
			dataLayerStr = dataLayerRegex.FindStringSubmatch(content)[1]
			json.Unmarshal([]byte(dataLayerStr), &dataLayerMap)
			skus := dataLayerMap["skus"].([]interface{})
			skuInfo := skus[0].(map[string]interface{})
			product := &productInfo{
				Name: dataLayerMap["name"].(string),
				ImgURL: skuInfo["image"].(string),
				Price: formatPrice(skuInfo["bestPrice"].(float64), skuInfo["available"].(bool)),
				Availability: skuInfo["available"].(bool),
				URL: strings.Replace(urls[itemIdx], "?sc=3", "",1),
			} 
			outputData = append(outputData, *product)
		}
		
		itemIdx++
	}

	return outputData
}

func main() {
	urls := listItemURLs()
	data := crawlItemData(urls)
	dataFile, _ := json.Marshal(data)
	_ = ioutil.WriteFile("output.json", dataFile, 0644)
}