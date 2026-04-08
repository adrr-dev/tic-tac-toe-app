// Package repository contains the model and repo logic
package repository

type Game struct {
	Board   []string `json:"board"`   // 3 by 3 board, len 9
	Current string   `json:"current"` // current player, X/O
	Winner  string   `json:"winner"`  // X, O, Draw
	Count   int      `json:"count"`   // how many squares are full
	Active  bool     `json:"active"`
	XWin    int      `json:"xwin"`
	OWin    int      `json:"owin"`
}

var Wins = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
	{0, 4, 8}, {2, 4, 6}, // Diagonals
}
