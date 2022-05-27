package handlers

import (
	"log"
	"net/http"
	"os"
	r "prueba-go/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()
	router.HandleFunc("/test", r.TestAPI).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3500"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
