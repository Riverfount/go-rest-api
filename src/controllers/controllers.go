package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Riverfount/go-rest-api/src/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func GetAllPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
