package main

import (
	"fmt"

	"github.com/Jameszhanqingsheng/chess/internal/board"
)

var chesses map[string]*board.ChessMan = make(map[string]*board.ChessMan)
var chesseBoard *board.Board

// := board.InitBoard()

func InitChess(symbol string, position []int, belong string, icon string, row int) {
	// [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	for index, value := range position {

		chess, err := board.InitChess(symbol, belong, 0, [2]int{row, value}, icon)

		if err != nil {
			fmt.Println(err)
			break
		}

		chesses[fmt.Sprintf("%s%s%s", belong, symbol, string(index))] = chess
	}
}

func InitChesses() {
	InitChess("P", []int{1, 2, 3, 4, 5, 6, 7, 8}, "W", "♙", 7)
	InitChess("P", []int{1, 2, 3, 4, 5, 6, 7, 8}, "B", "♟", 2)
	InitChess("R", []int{1, 8}, "W", "♖", 8)
	InitChess("R", []int{1, 8}, "B", "♜", 1)
	InitChess("N", []int{2, 7}, "W", "♘", 8)
	InitChess("N", []int{2, 7}, "B", "♞", 1)
	InitChess("B", []int{3, 6}, "W", "♗", 8)
	InitChess("B", []int{3, 6}, "B", "♝", 1)
	InitChess("Q", []int{4}, "W", "♕", 8)
	InitChess("Q", []int{4}, "B", "♛", 1)
	InitChess("K", []int{5}, "W", "♔", 8)
	InitChess("K", []int{5}, "B", "♚", 1)
}

func SetChessInBoard() {
	for _, v := range chesses {
		chesseBoard.SetChess(v.Position, v.Icon)
	}
}
