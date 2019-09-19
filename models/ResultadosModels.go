package models

import("time")

type Resultados struct{	
IdResultado 	int  `gorm:"column:idResultado"`	
IdMP 			int  `gorm:"column:idMP"`	
Anio 			int  `gorm:"column:Anio"`	
Mes 			int  `gorm:"column:Mes"`	
Semana 			int  `gorm:"column:Semana"`	
Dia 			int  `gorm:"column:Dia"`	
Valor 			float32  `gorm:"column:Valor"`	
FechaModificacion time.Time   `gorm:"column:FechaModificacion"`
}

