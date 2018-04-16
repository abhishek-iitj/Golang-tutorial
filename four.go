//News Reader Website
//Getting XML Document
//Parsing XML

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"` //slice of location type
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword  string
	Location string
}

//[5]type == array
//[5 5] int==2D array 5x5
//[]type == slice

func main() {
	//get information form the internet
	var s SitemapIndex
	var n News
	news_map := make(map[string]NewsMap) //A map having 2 things in the value
	//the  " _ " is a vraible that we don't want to use it
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)

	//fmt.Println(s.Locations)
	//iterating over the data structure
	for _, Location := range s.Locations {
		//fmt.Printf("\n%s", Location) //we want to visit the location as well rather than printing it.
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}
	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}

	fmt.Println("\n")
}
