package routers

import (
	"encoding/json"
	"strconv"

	"github.com/CloudAcademy/bd"
	"github.com/CloudAcademy/models"
)

func InsertCursos(body string, User string) (int, string) {
	var t models.Cursos

	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.Curso_Tittle) == 0 {
		return 400, "Debe especificar el Nombre (Title) del curso"
	}

	if len(t.Curso_Description) == 0 {
		return 400, "Debe especificar la descripcion (Ruta) del curso"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCursos(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el registro de la categoria " + t.Curso_Tittle + " > " + err2.Error()
	}

	return 200, "{ Curso_Id: " + strconv.Itoa(int(result)) + "}"
}
