package Server

import (
	"log"
	"time"
	"fmt"
	"net/http"
)

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