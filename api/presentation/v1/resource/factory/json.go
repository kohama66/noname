package factory

import (
	"encoding/json"
	"net/http"

	"github.com/myapp/noname/api/presentation/v1/resource"
)

func JSON(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ErrorJSON(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&resource.Error{Message: err})
}
