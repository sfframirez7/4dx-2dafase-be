package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func BrujulaPorMPCreate(w http.ResponseWriter, r *http.Request) {

	var newBrujula models.BrujulaLista
	err := json.NewDecoder(r.Body).Decode(&newBrujula)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var Brujula, err2 = services.BrujulaPorMPCreate(newBrujula)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Brujula)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func BrujulasPorMPGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	codigoEmpleado := vars["codigoEmpleado"]

	if codigoEmpleado == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var Brujulas, err2 = services.BrujulasPorMPGet(codigoEmpleado)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Brujulas)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func BrujulaPorMPUpdate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	var newBrujula models.Brujula
	err := json.NewDecoder(r.Body).Decode(&newBrujula)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var Brujula, err2 = services.BrujulaEstadoUpdate(newBrujula)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Brujula)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func BrujulaEstadosGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	var Brujulas, err2 = services.BrujulaEstadoGet()

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Brujulas)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func BrujulaActividadesPorMP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	idMP := vars["idMP"]
	mes := vars["mes"]

	if idMP == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var Brujulas, err2 = services.BrujulaActividadesPorMP(idMP, mes)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Brujulas)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func BrujulaActividadesPorColaborador(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	codigoEmpleado := vars["codigoEmpleado"]
	idEstado := vars["idEstado"]
	esLider := vars["esLider"]

	fmt.Println(idEstado)

	if codigoEmpleado == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var Brujulas, err2 = services.BrujulaActividadesPorColaborador(codigoEmpleado, idEstado, esLider)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Brujulas)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}

func GetBrujulaCantidad(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "Application/json")

	vars := mux.Vars(r)
	idColaborador := vars["idColaborador"]
	fmt.Println(idColaborador)
	if idColaborador == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Información incorrectas")
		return
	}

	var Brujulas, err2 = services.GetFBrujulaCantidad(idColaborador)

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No se ha podido obtener la data")
		return
	}

	response, _ := json.Marshal(Brujulas)

	responseString := string(response)

	fmt.Fprint(w, responseString)
}
