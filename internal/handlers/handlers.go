// Package handlers contains the handling logic
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/adrr-dev/tic-tac-toe-app/internal/repository"
)

type ServiceGame interface {
	NewGame() (*repository.Game, error)
	FetchGame() (*repository.Game, error)
	InputCell(cell int) error
	CheckWin() error
	RestartBoard() (*repository.Game, error)
}

type Handling struct {
	Service   ServiceGame
	Tmpls     *template.Template
	Fragments *template.Template
}

func NewHandling(service ServiceGame, tmpls, fragments *template.Template) *Handling {
	newHandling := &Handling{Service: service, Tmpls: tmpls, Fragments: fragments}
	return newHandling
}

func (h Handling) RootHandle(w http.ResponseWriter, r *http.Request) {
	newGame, err := h.Service.NewGame()
	if err != nil {
		notice := fmt.Sprintf("trouble loading game: %e", err)
		http.Error(w, notice, http.StatusBadRequest)
	}

	err = h.Tmpls.ExecuteTemplate(w, "index.html", newGame)
	if err != nil {
		notice := fmt.Sprintf("trouble with template: %e", err)
		http.Error(w, notice, http.StatusNotFound)
	}
}

func (h Handling) CellHandle(w http.ResponseWriter, r *http.Request) {
	cell := r.FormValue("cell")
	intCell, err := strconv.Atoi(cell)
	if err != nil {
		notice := fmt.Sprintf("could not convert string cell to int: %e", err)
		http.Error(w, notice, http.StatusBadRequest)
	}

	err = h.Service.InputCell(intCell)
	if err != nil {
		notice := fmt.Sprintf("something went wrong with service logic: %e", err)
		http.Error(w, notice, http.StatusBadRequest)
	}

	err = h.Service.CheckWin()
	if err != nil {
		notice := fmt.Sprintf("something went wrong check win logic: %e", err)
		http.Error(w, notice, http.StatusBadRequest)

	}

	data, err := h.Service.FetchGame()
	if err != nil {
		notice := fmt.Sprintf("trouble loading game: %e", err)
		http.Error(w, notice, http.StatusBadRequest)
	}

	err = h.Fragments.ExecuteTemplate(w, "reload.html", data)
	if err != nil {
		log.Println(err)
	}
}

func (h Handling) RestartHandle(w http.ResponseWriter, r *http.Request) {
	game, err := h.Service.RestartBoard()
	if err != nil {
		notice := fmt.Sprintf("trouble loading game: %e", err)
		http.Error(w, notice, http.StatusBadRequest)
	}

	err = h.Fragments.ExecuteTemplate(w, "reload.html", game)
	if err != nil {
		log.Println(err)
	}
}
