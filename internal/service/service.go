// Package service contains the service logic
package service

import (
	"log"

	"github.com/adrr-dev/tic-tac-toe-app/internal/repository"
)

type RepoGame interface {
	GetGame() (*repository.Game, error)
	UpdateGame(game *repository.Game) error
}

type Service struct {
	Repo RepoGame
}

func NewService(repo RepoGame) *Service {
	newService := &Service{Repo: repo}
	return newService
}

func (s Service) NewGame() (*repository.Game, error) {
	newBoard := make([]string, 9)
	newCurrent := "X"
	newGame := &repository.Game{Board: newBoard, Current: newCurrent, Active: true}
	err := s.Repo.UpdateGame(newGame)
	if err != nil {
		return nil, err
	}

	return newGame, nil
}

func (s Service) RestartBoard() (*repository.Game, error) {
	game, err := s.Repo.GetGame()
	if err != nil {
		return nil, err
	}
	board := make([]string, 9)
	game.Board = board
	game.Current = "X"
	game.Active = true
	game.Count = 0
	err = s.Repo.UpdateGame(game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (s Service) FetchGame() (*repository.Game, error) {
	game, err := s.Repo.GetGame()
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (s Service) InputCell(cell int) error {
	game, err := s.Repo.GetGame()
	if err != nil {
		return err
	}

	if !game.Active {
		log.Println("game not active")
		return nil
	}

	if game.Board[cell] == "" {
		game.Board[cell] = game.Current
		game.Count++

		if game.Current == "X" {
			game.Current = "O"
		} else {
			game.Current = "X"
		}
		err = s.Repo.UpdateGame(game)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) CheckWin() error {
	game, err := s.Repo.GetGame()
	if err != nil {
		return err
	}

	if !game.Active {
		log.Println("game not active")
		return nil
	}

	board := game.Board

	if game.Count >= 9 {
		game.Winner = "draw"
		game.Active = false

		err = s.Repo.UpdateGame(game)
		if err != nil {
			return err
		}
		return nil
	}
	for _, line := range repository.Wins {
		a, b, c := line[0], line[1], line[2]
		if board[a] != "" && board[a] == board[b] && board[a] == board[c] {
			game.Winner = board[a]
			game.Active = false

			if game.Current == "X" {
				game.OWin++
			} else {
				game.XWin++
			}

			err = s.Repo.UpdateGame(game)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}
