package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/zedithx/devops_practice/handlers"
)

func main(){
	http.HandleFunc("/", handlers.HelloHandler)
	fmt.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
