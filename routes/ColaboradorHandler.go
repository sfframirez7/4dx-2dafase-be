package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"
	"github.com/gorilla/mux"
)

func ColaboradoresPorArea(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	AreaId := vars["AreaId"]

	var subAreas, err = services.ColaboradoresPorArea(AreaId)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(&subAreas)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetColaboradoresSubArea(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Application/json")

	vars := mux.Vars(r)
	idSubArea := vars["idSubArea"]

	var usuarioModel, erro = services.GetColaboradoresSubArea(idSubArea)

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
func GetColaboradoresAdmins(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Expose-Headers: Content-Length", "X-JSON")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-type", "Application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Application/json")

	var usuarioModel, erro = services.ColaboradoresAdmins()

	if erro != nil {
		fmt.Println(erro)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, erro)
		return
	}

	response, _ := json.Marshal(&usuarioModel)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
