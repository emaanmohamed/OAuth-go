package main

import (
	"github.com/emaanmohamed/OAuth-go/auth"
	"net/http"
)

func main() {
	auth.NewAuth()
	var authInstance = &auth.Auth{}
	router := authInstance.RegisterRoutes()
	http.ListenAndServe(":8080", router)
}
