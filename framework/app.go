package framework

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// App is a struct that we will use to
// prepare the server and router s
type App struct {
	Router *Router
}

// New creates a new App and initializes
// everything we need to start it up
func New() *App {
	loadEnv()
	r := NewRouter()
	return &App{r}
}

// Start will listen to the host and port that we
// defined in our .env variables
func (a *App) Start() {

	appHost := fmt.Sprintf(
		"%s:%s",
		os.Getenv("APP_HOST"),
		os.Getenv("APP_PORT"),
	)

	// Bind to a port and pass our router in
	log.Println("Server listening at " + appHost)
	log.Fatal(http.ListenAndServe(appHost, a.Router))
}

// This will make sure that we have a .env
// file when we start a new app
func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
