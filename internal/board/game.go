package board

import (
	"fmt"
)

type Game struct {
	Chesses     map[string]*ChessMan `json:"chesses"`
	ChesseBoard *Board               `json:"board"`
}

func InitGame() *Game {
	board := InitBoard()
	chesses := make(map[string]*ChessMan)

	return &Game{
		ChesseBoard: board,
		Chesses:     chesses,
	}
}

func (g *Game) InitChess(symbol string, position []int, belong string, icon string, row int) {
	// [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	for index, value := range position {

		chess, err := InitChessMan(symbol, belong, 0, [2]int{row, value}, icon)

		if err != nil {
			fmt.Println(err)
			break
		}

		g.Chesses[fmt.Sprintf("%s%s%s", belong, symbol, string(index))] = chess
	}
}

func (g *Game) InitChesses() {
	g.InitChess("P", []int{1, 2, 3, 4, 5, 6, 7, 8}, "W", "♙", 7)
	g.InitChess("P", []int{1, 2, 3, 4, 5, 6, 7, 8}, "B", "♟", 2)
	g.InitChess("R", []int{1, 8}, "W", "♖", 8)
	g.InitChess("R", []int{1, 8}, "B", "♜", 1)
	g.InitChess("N", []int{2, 7}, "W", "♘", 8)
	g.InitChess("N", []int{2, 7}, "B", "♞", 1)
	g.InitChess("B", []int{3, 6}, "W", "♗", 8)
	g.InitChess("B", []int{3, 6}, "B", "♝", 1)
	g.InitChess("Q", []int{4}, "W", "♕", 8)
	g.InitChess("Q", []int{4}, "B", "♛", 1)
	g.InitChess("K", []int{5}, "W", "♔", 8)
	g.InitChess("K", []int{5}, "B", "♚", 1)
}

func (g *Game) SetChessInBoard() {
	for _, v := range g.Chesses {
		g.ChesseBoard.SetChess(v.Position, v.Icon)
	}
}
