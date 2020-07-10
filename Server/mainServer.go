package Server

import (
	"log"
	"github.com/gorilla/securecookie"
	"net/http"
)

/**
*	variable para definir las cookies de sesión
*
*	---NO MODIFICARLA EN NINGUN PUNTO DE EJECUCIÓN---
*/
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

/**
*	la esta función permite iniciar la ejecución del servidor
*/
func RunServer()  {

	//a la función NewServer() se le debe enviar el puerto a usar
	server := NewServer(":8081")

	//estas dos líneas sirven el directorio que contiene los archivos estaticos a usar
	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	/**
	*	definir en este punto los Handlers correspondientes a las rutas
	*	teniendo en cuante que cada handler solo admite un tipo de solicitud
	*/
	server.Handle("GET" ,"/", server.AddMiddleware(HandleRoot, Log()))

	log.Fatal(server.Listen())
	
}