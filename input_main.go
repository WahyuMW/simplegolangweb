package main

import (
		//"encoding/json"
        "fmt"
        "net/http"
        "html/template"
  	  	"log"
  	  	"regexp"
  	  	"strconv"

)

/*
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
*/
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
		num := r.FormValue("phone_num")
		s := strconv.Itoa(regex(num))
		fmt.Fprintf(w, "Country Code: ")
		fmt.Fprintf(w, r.FormValue("code"))
		fmt.Fprintf(w, "\nPhone Number: ")
		fmt.Fprintf(w, num)
		fmt.Fprintf(w, "\nRespons: ")
		fmt.Fprintf(w, s)	

    }
}

func regex(num string) int {
	 match, _ := regexp.MatchString("^08[0-9]{9,11}$", num )
    fmt.Println("regexp: ", match)
    if match == true {
    	respons := 200
    	return respons
    } else {
    	respons := 400
    	return respons
    }
}

func main() {  
//	http.HandleFunc("/", welcome) // setting router rule
    http.HandleFunc("/", input)
	    err := http.ListenAndServe(":9090", nil ) // setting listening port
	    if err != nil {
	        log.Fatal("ListenAndServe: ", err)
	    }
}

//func validation ()