package main

import (
	"log"
	"prueba-go/handlers"
)

func main() {
	log.Println("Escuchando en el puerto 3500")
	handlers.Handlers()

}
