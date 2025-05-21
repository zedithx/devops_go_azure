package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/zedithx/devops_practice/routes"
	"github.com/zedithx/devops_practice/internal/metrics"
)

func main(){
	metrics.InitMetrics()
	r := routes.SetupRouter()
	fmt.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", r))
}
