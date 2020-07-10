package Server

/*
*Agregar en este fichero los handles con dos parametros (http.ResponseWriter, *http.Request)
*/

import (
	"net/http"
	"html/template"
	"path"
)

//Handle principal por defecto
func HandleRoot(w http.ResponseWriter, r *http.Request) {

	context := IndexContext{
		Title: "Index",
	}

	//Renderizar archivo html
	templatePath:= path.Join("views/", "index.html")
	template, _ := template.ParseFiles(templatePath)
	template.Execute(w, context)

}