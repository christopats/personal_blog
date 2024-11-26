package main

import (
	"fmt"
	"net/http"
)

// DECLARE A HANDLER THAT RESPONDS WITH INFO ABOUT APP STATUS, ENV AND VERSION

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}