package car_pooling

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Init() {
	http.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		_, _ = fmt.Fprint(writer, "OK")
	})

	http.HandleFunc("/cars", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPut {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		err := json.NewDecoder(request.Body).Decode(&cars)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
	})
}
