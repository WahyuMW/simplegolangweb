package main

import (

        "fmt"
        "net/http"
        "html/template"
  	  	"log"
    	"strings"

)


func welcome(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() 
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "welocme to input page") // write data to response
}

func input(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("input.html")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        // logic part of input
        fmt.Println("country code:", r.Form["code"])
        fmt.Println("phone number:", r.Form["phone_num"])
        fmt.Fprintf(w, "Response: 200")
    }
}

func main() {  
	http.HandleFunc("/", welcome) // setting router rule
    http.HandleFunc("/input", input)
	    err := http.ListenAndServe(":9090", nil ) // setting listening port
	    if err != nil {
	        log.Fatal("ListenAndServe: ", err)
	    }
}

//func validation ()