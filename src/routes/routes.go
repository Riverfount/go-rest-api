package routes

import (
	"log"
	"net/http"

	"github.com/Riverfount/go-rest-api/src/controllers"
)

func HandlerRequests() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/presonalidades", controllers.GetAllPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
