package Server

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type IndexContext struct{
	Title string
}