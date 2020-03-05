package car_pooling

import (
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
}
