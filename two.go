//Basic web server in Go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amzing News Aggregator", News: "Some news"}
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
