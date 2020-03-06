package car_pooling

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
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

		sort.SliceStable(cars, func(i, j int) bool {
			return cars[i].Seats < cars[j].Seats
		})
		log.Println(cars)
	})

	http.HandleFunc("/journey", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var group Group
		err := json.NewDecoder(request.Body).Decode(&group)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		groups = append(groups, group)
		log.Println(groups)
	})
}
