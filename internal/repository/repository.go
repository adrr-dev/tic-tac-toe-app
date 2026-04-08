package repository

import (
	"encoding/json"
	"os"
)

type Repo struct {
	DataFile string
}

func NewRepo(dataFile string) *Repo {
	newRepo := &Repo{DataFile: dataFile}
	return newRepo
}

func (r Repo) GetGame() (*Game, error) {
	file, err := os.Open(r.DataFile)
	if err != nil {
		return nil, err
	}
	dec := json.NewDecoder(file)

	defer func() { _ = file.Close() }()

	var data Game
	err = dec.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r Repo) UpdateGame(game *Game) error {
	file, err := os.Create(r.DataFile)
	if err != nil {
		return err
	}

	defer func() { _ = file.Close() }()

	enc := json.NewEncoder(file)
	err = enc.Encode(game)
	if err != nil {
		return err
	}
	return nil
}
