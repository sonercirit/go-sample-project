package main

import (
	"fmt"
	"gitlab.com/cabify-challenge/car-pooling-challenge-sonercirit/car_pooling"
	"log"
	"net/http"
	"os"
)

func main() {
	car_pooling.Init()
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), nil)
	if err != nil {
		log.Println(err)
	}
}
