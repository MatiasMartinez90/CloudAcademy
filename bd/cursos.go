package bd

import (
	"database/sql"
	"fmt"

	//"strconv"
	"github.com/CloudAcademy/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertCursos(c models.Cursos) (int64, error) {
	fmt.Println("Inicializando funcion  db.InsertCursos")

	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO cursos2 (Curso_Tittle, Curso_Description ) Values('" + c.Curso_Tittle + "','" + c.Curso_Description + "')"

	fmt.Println("ejecutnado sentencia")
	fmt.Println(sentencia)
	var result sql.Result

	result, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		fmt.Println(err.Error())
		return 0, err2
	}

	fmt.Println("Insert Cursos > Ejecucion Exitosa")
	return LastInsertId, nil

}
