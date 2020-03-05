package car_pooling

import "net/http"

func Init() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	})
}
