package main

import (
	"fmt"

	"github.com/ZQingS/chess/internal/board"
	"github.com/ZQingS/chess/internal/protocol"
)

type Process struct {
	SessionID   string
	Belong      string
	Board       [8][8]string
	Round       int // (W/B) (0/1)
	Start       int // (start/stop) (1/0)
	CurrentSend *protocol.Send
	receive     *protocol.Receive
}

func (g *Process) HandleReadyChanPool() string {
	if len(g.SessionID) > 0 && len(g.Belong) > 0 {
		return fmt.Sprintf("%s:%s:connect", g.SessionID, g.Belong)
	}

	return ""
}

func (g *Process) HaveOtherReadyChanPool(chanPool string) bool {
	var belong string

	if g.Belong == "W" {
		belong = "B"
	} else {
		belong = "W"
	}

	if chanPool == fmt.Sprintf("%s:%s:connect", g.SessionID, belong) {
		return true
	}

	return false
}

func (g *Process) HandleQuitMessage() bool {
	if g.receive.Action == "quit" {
		g.CurrentSend.OK = "false"
		g.CurrentSend.Msg = "Quit the Game"

		return true
	}

	return false
}

func (g *Process) HandleBelongMessage() {
	if g.receive.Action == "belong" {
		g.Belong = g.receive.Belong
		g.CurrentSend.OK = "true"
		g.CurrentSend.Msg = fmt.Sprintf("Choose Belong %s", g.Belong)
	}
}

func (g *Process) HandleSessionMessage() {
	if g.receive.Action == "session" {
		g.SessionID = g.receive.SessionID
		g.CurrentSend.OK = "true"
		g.CurrentSend.Msg = fmt.Sprintf("Join The Game %s, Wait Other", g.SessionID)
	}
}

// 需要双方准备好了后才可以进行开启

func (g *Process) HandleSatrtMessage(start bool) {
	if start {
		g.Start = 1
	} else {
		g.Start = 0
	}
}

func (g *Process) HandleRoundMessage(finishMy bool) {
	if finishMy {
		if g.Belong == "W" {
			g.Round = 1
		} else {
			g.Round = 0
		}
	}
}

func (g *Process) HandleBoardMessage() {
	if g.Start == 1 {
		var resBoard = &board.Board{}

		if len(g.Board[0][0]) == 0 {
			game := NewORContinueGame(g.SessionID)

			resBoard = game.ChesseBoard
		} else {
			resBoard = resBoard.SetChesses(g.Board)
		}

		if g.Belong == "W" {
			g.Board = resBoard.GetChesses()
			g.CurrentSend.Board = resBoard.GetChesses()
		} else {
			g.Board = resBoard.GetReverseChesses()
			g.CurrentSend.Board = resBoard.GetReverseChesses()
		}
	}
}
