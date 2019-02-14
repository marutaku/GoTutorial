package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //urlが渡すオプションを解析します。POSTに対してはレスポンスパケットのボディを解析します（request body）
	fmt.Println("path", r.Form)
	fmt.Println("URL", r.URL)
	fmt.Println("scehme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key", k)
		fmt.Println("Value", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello, astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		// r.Form[]は事前にr.ParseForm()をコールする必要があるが、
		// r.FormValue()は事前にParseForm()をコールする必要がない
		fmt.Println("username:", r.FormValue("username"))
		fmt.Println("password", r.FormValue("password"))
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Listen and serve:", err)
	}
}
