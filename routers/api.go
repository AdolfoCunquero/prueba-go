package routers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	m "prueba-go/models"
)

//Total de registros unicos que necesitamos
var totalRecords = 15

//Tamaño de bloques o solicitudes simultaneas para no crear demasiados hilos
var batchZize = 15

func TestAPI(rw http.ResponseWriter, r *http.Request) {
	allResults := make([]m.Item, 0)
	count := 0
	c := make(chan bool)

	if batchZize > totalRecords {
		batchZize = totalRecords
	}

	for count < totalRecords {
		countBatch := 0
		for i := 1; i <= batchZize; i++ {
			go getItem(&countBatch, &count, &allResults, c)
		}
		v := <-c
		if v {
			fmt.Println("Bloque finalizado")
		}
	}

	jsonResponse(rw, allResults, 200)
}

func getItem(conteoBatch *int, conteo *int, results *[]m.Item, c chan bool) {
	var result m.Item
	filtered := make([]m.Item, 0)

	if *conteo == totalRecords || *conteoBatch == batchZize {
		return
	}

	resp, err := http.Get("https://api.chucknorris.io/jokes/random")

	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	for _, item := range *results {
		if item.Id == result.Id {
			filtered = append(filtered, item)
		}
	}

	if len(filtered) == 0 {

		*results = append(*results, result)

		*conteoBatch += 1
		*conteo += 1
		if *conteoBatch == batchZize {
			c <- true
		}
	}
}

func jsonResponse(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
