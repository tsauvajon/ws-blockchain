package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/tsauvajon/ws-blockchain/server/sportmonks"

	"github.com/gorilla/mux"
)

// App : application
type App struct {
	Router *mux.Router
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (app *App) initializeRoutes() {
	static := "../client/dist/static/"

	// Adds the /api/ prefix to every REST request
	api := app.Router.PathPrefix("/api").Subrouter()

	// Adds the /soccer/ prefix for soccer requests
	soccer := api.PathPrefix("/soccer").Subrouter()

	// Handles some functions
	soccer.HandleFunc("/teams/{id:[0-9]+}", app.getTeamByID).Methods("GET")
	soccer.HandleFunc("/fixtures/{deb}/{fin}", app.getFixtureByDate).Methods("GET")
	soccer.HandleFunc("/bets/{id:[0-9]+}", app.getBetByFixture).Methods("GET")

	// Adds the /static prefix to serve the static files generated by the Vue app
	app.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(static))))

	// Add a handler for the root (i.e. "/"). Will serve the front-end
	app.Router.HandleFunc("/", app.indexHandler).Methods("GET")
}

func (app *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	entry := "../client/dist/index.html"

	// open and parse a template text file
	if tpl, err := template.New("index").ParseFiles(entry); err != nil {
		log.Fatal(err)
	} else {
		tpl.Lookup("index").ExecuteTemplate(w, "index.html", nil)
	}
}

// Initialize : initializes the app
// Connection to the databases etc.
func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRoutes()

	// Initialize the SportMonks package
	if err := sportmonks.Initialize(); err != nil {
		panic(err)
	}
}

// Run : Runs the application
func (app *App) Run(addr string) {
	fmt.Println("Running at ", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
