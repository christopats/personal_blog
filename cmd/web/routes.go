package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// USING THIRDPARTY ROUTER TO ENABLE DIFFERENT HTTP REQUESTS FROM SAME ENDPOINT
// EXTRACTING ROUTING LOGIC FROM MAIN()
// ROUTES() IS CREATED AS A METHOD ON APPLICATION STRUCT


func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodGet, "/", app.Index)
	router.HandlerFunc(http.MethodGet, "/home", app.Home)
	router.HandlerFunc(http.MethodGet, "/about", app.About)
	router.HandlerFunc(http.MethodGet, "/blog", app.Blog)
	router.HandlerFunc(http.MethodGet, "/blog/", app.BlogPost)
	router.HandlerFunc(http.MethodGet, "/projects", app.Projects)
	router.HandlerFunc(http.MethodGet, "/projects/", app.ProjectPostPage)
	router.HandlerFunc(http.MethodGet, "/contact", app.Contact)
	router.Handler(http.MethodGet, "/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	return router

}