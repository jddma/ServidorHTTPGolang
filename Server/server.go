package Server

import (
	"net/http"
)

//struct Server ...
type Server struct{

	port string
	router *Router

}

//función encargada de implementar los middlewares deseados a un handler
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {

	for _, m := range middlewares {
		f = m(f)
	}
	return f

}


//función que se ejecutara para iniciar el servidor
func (s *Server) Listen() error{

	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil{
		return err
	}
	return nil

}

//función usada para definir los Handles del servidor con sus respectias rutas
func (s *Server) Handle(method string, path string, handler http.HandlerFunc){

	_, exist := s.router.rules[path]
	if !exist{
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler

}

//NewServer ...
func NewServer(port string) *Server{

	return &Server{
		port: port,
		router: NewRouter(),
	}

}

