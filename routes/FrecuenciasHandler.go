package routes

import (	
	"encoding/json"
	"fmt"
	"net/http"

	//"../models"
	"../services"
)

func GetFrecuencias(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")	
	
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")

	fmt.Println("GeMediciones Handler")

	fmt.Println(r)

	var usuarioModel, erro = services.GetFrecuencias()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	fmt.Fprintf(w, string(response))
}