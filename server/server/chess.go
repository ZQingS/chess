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

		initBoard.SetChesses(initbuffer.Data.Board)

		return &board.Game{
			Chesses:     initbuffer.Data.Chesses,
			ChesseBoard: initBoard,
		}
	} else {
		return newGame()
	}
}

func HandleBoardMessage(sessionId string, belong string, currentBoard [8][8]string, positionCommand string) [8][8]string {
	var resBoard = &board.Board{}

	if len(currentBoard) == 0 {
		game := NewORContinueGame(sessionId)

		resBoard = game.ChesseBoard
	} else {
		resBoard.SetChesses(currentBoard)
	}

	if belong == "White" {
		return resBoard.GetChesses()
	} else {
		return resBoard.GetReverseChesses()
	}
}
