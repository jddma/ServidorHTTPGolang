package Server

/*
*agregar aqui los handles con dos parametros (http.ResponseWriter, *http.*Request)
*
*HandleRoot y HandlerInicio son solamente ejemplos pueden ser eliminados o
*modificados
*/

import (
	"fmt"
	"net/http"
)

//Handle principal por defecto
func HandleRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Root")

}