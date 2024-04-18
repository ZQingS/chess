package main

import (
	"github.com/ZQingS/chess/internal/board"
	"github.com/ZQingS/chess/internal/chessdb"
)

func newGame() *board.Game {
	game := board.InitGame()
	game.InitChesses()
	game.SetChessInBoard()

	return game
}

func NewORContinueGame(sessionId string) *board.Game {
	storage := chessdb.Storage{
		SessionId: sessionId,
	}

	var bufdata chessdb.BufferData

	initbuffer := &chessdb.Buffer{
		Storage: storage,
		Data:    bufdata,
	}

	if initbuffer.DataExist() {
		initbuffer.GetData()

		var initBoard = &board.Board{}

		resBoard := initBoard.SetChesses(initbuffer.Data.Board)

		return &board.Game{
			Chesses:     initbuffer.Data.Chesses,
			ChesseBoard: resBoard,
		}
	} else {
		return newGame()
	}
}
