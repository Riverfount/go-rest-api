package main

import (
	"fmt"

	"github.com/Riverfount/go-rest-api/src/routes"
	"github.com/Riverfount/go-rest-api/src/models"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{
			Nome:     "Albert Einstein",
			Historia: "Físico teórico que desenvolveu a teoria da relatividade.",
		},
		{
			Nome:     "Marie Curie",
			Historia: "Pioneira na pesquisa sobre radioatividade.",
		},
	}
	fmt.Println("Inicializando Servidor REST em Go ...")
	routes.HandlerRequests()
}
