package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port:=os.Args[1]
	fmt.Println(port)

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"<h1>hello world</h1>")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port),nil))

}