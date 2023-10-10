package bd

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/CloudAcademy/tools"

	//"strconv"
	"github.com/CloudAcademy/models"
	_ "github.com/go-sql-driver/mysql"
)

func InsertCursos(p models.Cursos) (int64, error) {
	fmt.Println("Inicializando funcion  db.InsertCursos")

	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO cursos (Curso_Tittle"

	if len(p.Curso_Description) > 0 {
		sentencia += ", Curso_Description"
	}

	if p.Curso_Price > 0 {
		sentencia += ", Curso_Price"
	}

	if p.Curso_CategoryId > 0 {
		sentencia += ", Curso_CategoryId"
	}

	if len(p.Curso_Path) > 0 {
		sentencia += ", Curso_Path"
	}

	sentencia += ") Values ('" + tools.EscapeString(p.Curso_Tittle) + "'"

	if len(p.Curso_Description) > 0 {
		sentencia += ", '" + tools.EscapeString(p.Curso_Description) + "'"
	}

	if p.Curso_Price > 0 {
		sentencia += ", " + strconv.FormatFloat(p.Curso_Price, 'e', -1, 64)
	}

	if p.Curso_CategoryId > 0 {
		sentencia += ", " + strconv.Itoa(p.Curso_CategoryId)
	}

	if len(p.Curso_Path) > 0 {
		sentencia += ", '" + tools.EscapeString(p.Curso_Path) + "'"
	}

	sentencia += ")"

	var result sql.Result

	fmt.Println("Vamos a ejecutar la sentancia: ")
	fmt.Println(sentencia)
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
