package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

// App ...
type App struct {
	Router *mux.Router
	DB     *pg.DB
}

// Start ...
func Start() {
	app := App{}
	app.DB = startDatabase()
	defer app.DB.Close()

	app.Router = mux.NewRouter()
	app.initializeRoutes()

	srv := &http.Server{
		Handler:      app.Router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/", app.ListURLHandler).Methods("GET")
	app.Router.HandleFunc("/url", app.CreateURLHandler).Methods("POST")
	app.Router.HandleFunc("/url/{hash:[a-z0-9]+}", app.RedirectHandler).Methods("GET")
}

// startDatabase ...
func startDatabase() *pg.DB {
	cfg, err := readConfig("config")
	if err != nil {
		panic("read config failed " + err.Error())
	}

	cnf := &pg.Options{
		Addr:     cfg.GetString("dbHost"),
		User:     cfg.GetString("dbUser"),
		Database: cfg.GetString("dbDatabase"),
		Password: cfg.GetString("dbPassword"),
	}

	return pg.Connect(cnf)
}
