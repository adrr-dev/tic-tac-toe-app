package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/adrr-dev/tic-tac-toe-app/internal/handlers"
	"github.com/adrr-dev/tic-tac-toe-app/internal/repository"
	"github.com/adrr-dev/tic-tac-toe-app/internal/service"
)

func main() {
	// dataFile:= "data.json" //testing in local file
	dataFile, err := pathSetup()
	if err != nil {
		log.Fatal(err)
	}
	myRepo := repository.NewRepo(dataFile)
	myService := service.NewService(myRepo)

	tmpls, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	fragments, err := template.ParseGlob("templates/fragments/*.html")
	if err != nil {
		log.Fatal(err)
	}
	myHandling := &handlers.Handling{Service: myService, Tmpls: tmpls, Fragments: fragments}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /{$}", myHandling.RootHandle)
	mux.HandleFunc("POST /cell", myHandling.CellHandle)
	mux.HandleFunc("POST /restart", myHandling.RestartHandle)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func pathSetup() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dataPath := filepath.Join(home, ".local", "share", "tic-tac-toe")

	err = os.MkdirAll(dataPath, 0o755)
	if err != nil {
		return "", err
	}

	dbPath := "data.json"
	dataFile := filepath.Join(dataPath, dbPath)

	_, err = os.Stat(dataFile)
	if os.IsNotExist(err) {
		//  File doesn't exist, create it with an empty JSON map
		err = os.WriteFile(dataFile, nil, 0o644)
		if err != nil {
			return "", err
		}
	}

	return dataFile, nil
}
