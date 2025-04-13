package routes

import (
	"log"
	"net/http"

	"github.com/Riverfount/go-rest-api/src/controllers"
)

func HandlerRequests() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
