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

	data,err := os.ReadFile("./index.html")
	if err!=nil{
		panic("could not open the file")
	}




	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,string(data))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",port),nil))

}