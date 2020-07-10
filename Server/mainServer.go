package Server

import (
	"log"
)

/**
*	la esta función permite iniciar la ejecución del servidor
*/
func RunServer()  {

	//a la función NewServer() se le debe enviar el puerto a usar
	server := NewServer(":8081")

	/**
	*	definir en este punto los Handlers correspondientes a las rutas
	*	teniendo en cuante que cada handler solo admite un tipo de solicitud
	*/
	server.Handle("GET" ,"/", server.AddMiddleware(HandleRoot, Log()))

	log.Fatal(server.Listen())
	
}