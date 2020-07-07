package Server

import (
	"net/http"
	"fmt"
)

func RunServer()  {

	server := NewServer(":8081")

	//Servir el directorio que contiene los archivos estaticos
	fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	//definir los Handlers correspondientes a las rutas
	server.Handle("GET" ,"/", server.AddMiddleware(HandleRoot, Log()))

	fmt.Println(server.Listen())
	
}