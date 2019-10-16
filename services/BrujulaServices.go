package services

import (
	"time"

	"../config"
	"../models"
)

func BrujulaPorMPCreate(Modelo models.Brujula) (models.Brujula, error) {

	Modelo.IdBrujula = 0
	Modelo.FechaCreada = time.Now()
	Modelo.FechaModificada = time.Now()

	db := config.ConnectDB()
	defer db.Close()

	db.Create(&Modelo)

	return Modelo, nil
}

func BrujulasPorMPGet(codigoEmpleado string, idResultado string) ([]models.BrujulaConEstado, error) {

	var brujulas []models.BrujulaConEstado

	db := config.ConnectDB()
	defer db.Close()

	db.Raw("SELECT * FROM dbBrujulaPorMP b INNER JOIN dbEstadosBrujula e ON e.idEstado = b.idEstado where b.idColaborador = ? and b.IdResultadoMP = ? order by b.fechaCreada desc", codigoEmpleado, idResultado).Scan(&brujulas)

	return brujulas, nil
}

func BrujulaEstadoUpdate(Modelo models.Brujula) (models.Brujula, error) {

	var updatedBrujula models.Brujula

	//  if Modelo.Autorizado == true{
	db := config.ConnectDB()
	defer db.Close()

	db.Where("IdBrujula = ?", Modelo.IdBrujula).Find(&updatedBrujula)

	db.Model(&updatedBrujula).Where("IdBrujula= ?", Modelo.IdBrujula).Update(models.Brujula{IdEstado: Modelo.IdEstado, FechaModificada: time.Now()})

	return updatedBrujula, nil
}

func BrujulaEstadoGet() ([]models.BrujulaEstado, error) {

	var BrujulaEstados []models.BrujulaEstado

	db := config.ConnectDB()
	defer db.Close()

	db.Find(&BrujulaEstados)

	return BrujulaEstados, nil
}
