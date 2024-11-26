package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// CONFIG STRUCT HOLDS ALL APP SETTINGS
type config struct {
	port int
	env string
}

// APPLICATION STRUCT HOLDS DEPENDENCIES REQUIRED APP WIDE
// WE ATTACH OUR METHODS AND HELPERS TO THIS
type application struct {
	config config
	logger *log.Logger
}

func main() {

	// INIT CONFIG VAR
	var cfg config

	// READ CL FLAGS TO SET CFG VALUES, SET DEFAULTS FOR DEV
	flag.IntVar(&cfg.port, "port", 42069, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")


	// INIT LOGGER
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	// INIT INSTANCE OF APPLICATION AND GIVE IT CFG AND LOGGER 
	app := *&application{
		config: cfg,
		logger: logger,
	}
	// DECLARE OUR SERVEMUX AND HANDLERFUNCS
	// HANDLEFUNCS MOVED TO ROUTES.GO, USE HTTPROUTER INSTEAD OF SERVEMUX



	// INIT OUR CUSTOM HTTP SERVER 
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// START THE SERVER

	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)
	
	
}