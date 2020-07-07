package Server

import (
	"log"
	"time"
	"fmt"
	"net/http"
)

/**
*Este middleware permite verificar si un usuario esta autenticado
*en caso de que no lo este sera redireccinado a 
*/
func CheckAuth(URLToRedirect string) Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			redirect := true
			if redirect{
				http.Redirect(w, r, URLToRedirect, 307)
				return
			}
			f(w, r)
		}

	}

}

//esta función registra los logs a los handlers en los que se implementa
func Log() Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()
			defer func() { 

				fmt.Println("-------------------")
				log.Println("\nPATH=", r.URL.Path)
				fmt.Println("HOST=", r.RemoteAddr)
				fmt.Println("Method=", r.Method)
				fmt.Println(time.Since(start))
				fmt.Println("-------------------")

			}()
			f(w, r)

		}

	}
}