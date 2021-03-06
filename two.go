//Basic web server in Go
package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

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

type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	c <- n //send the n to the channel
}
func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	//get information form the internet
	var s SitemapIndex

	news_map := make(map[string]NewsMap) //A map having 2 things in the value
	//the  " _ " is a vraible that we don't want to use it
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()

	xml.Unmarshal(bytes, &s)
	queue := make(chan News, 500) //to be passed to our newsRoutine
	//fmt.Println(s.Locations)
	//iterating over the data structure
	for _, Location := range s.Locations {
		//fmt.Printf("\n%s", Location) //we want to visit the location as well rather than printing it.
		wg.Add(1)
		go newsRoutine(queue, Location)
	}
	wg.Wait()
	close(queue)
	for elem := range queue {
		for idx, _ := range elem.Titles {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}
	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}
	t, _ := template.ParseFiles("basictemplate.html")
	// fmt.Println(err)
	fmt.Println(t.Execute(w, p))

}
func index_handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Whoa, Go is neat");		//writes with the response writer

	fmt.Fprintf(w, "<h1>Hey There!!</h1>")
	fmt.Fprintf(w, "<p>Go is Fast!</p>")
	fmt.Fprintf(w, "<p>.....and simple!</p>")
	fmt.Fprintf(w, "<p>You %s even add %s</p>", "can", "<strong>variabel</strong>")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Rashi \n Rashi is a cute girl. She is so adorable. I just love her. ") //writes with the response writer
}

func main() {
	//A handler takes URL figure out path and what function handels the path
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.HandleFunc("/agg", newsAggHandler)
	http.ListenAndServe(":8080", nil)
}
