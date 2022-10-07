// Filename: cmd/api/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//Virsion of app
const version = "1.0"

//App Config
type config struct {
	port int
	env  string
}

//Dependencies injections
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	//Read flags that are need to build the config
	flag.IntVar(&cfg.port, "port", 4000, "API port")
	flag.StringVar(&cfg.env, "env", "dev", "(dev | stg | prds")
	flag.Parse()

	//Create a logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//Create and install app
	app := &application{
		config: cfg,
		logger: logger,
	}

	//Create HTTP Server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server at %s", cfg.env, srv.Addr)
	//Start the server
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
