package main

import (
	"fmt"

	"github.com/Riverfount/go-rest-api/src/routes"
)

func main() {
	fmt.Println("Inicializando Servidor REST em Go ...")
	routes.HandlerRequests()
}
