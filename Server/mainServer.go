package Server

import (
	"fmt"
)

func RunServer()  {

	server := NewServer(":8081")

	//definir los Handlers correspondientes a las rutas

	server.Handle("GET" ,"/", server.AddMiddleware(HandleRoot, Log()))

	fmt.Println(server.Listen())
	
}