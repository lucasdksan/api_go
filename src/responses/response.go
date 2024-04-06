package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status_code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status_code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func ERR(w http.ResponseWriter, status_code int, err error) {
	JSON(w, status_code, struct {
		Erro string `json:"erro"`
	}{
		Erro: err.Error(),
	})
}
