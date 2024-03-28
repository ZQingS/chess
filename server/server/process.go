package main

import (
	"fmt"
	"net"

	"github.com/ZQingS/chess/internal/board"
	"github.com/ZQingS/chess/internal/protocol"
)

type Process struct {
	SessionID   string
	Belong      string
	CurrentSend *protocol.Send
	receive     *protocol.Receive
	ConnPool    net.Conn
}

type GameProcess struct {
	White Process
	Black Process
	Round int // (W/B) (0/1)
	Board [8][8]string
}

func (gp *GameProcess) StartPlay() bool {
	return len(gp.White.Belong) > 0 && len(gp.Black.Belong) > 0 && len(gp.White.SessionID) > 0 && len(gp.Black.Belong) > 0
}

func (gp *GameProcess) HandleBoardMessage(belong string, sessionId string) {
	var resBoard = &board.Board{}

	if len(gp.Board[0][0]) == 0 {
		game := NewORContinueGame(sessionId)

		resBoard = game.ChesseBoard
	} else {
		resBoard = resBoard.SetChesses(gp.Board)
	}

	if belong == "W" {
		gp.White.CurrentSend.Board = resBoard.GetChesses()
	} else {
		gp.Black.CurrentSend.Board = resBoard.GetReverseChesses()
	}
}

func (gp *GameProcess) HandleRoundMessage(finishMy bool, belong string) {
	if finishMy {
		if belong == "W" {
			gp.Round = 1
		} else {
			gp.Round = 0
		}
	}
}

func (gp *GameProcess) InitProcess(p *Process) {
	process := *p

	if p.Belong == "W" {
		gp.White = process
	} else {
		gp.Black = process
	}
}

func (g *Process) HandleReadyChanPool() bool {
	if len(g.SessionID) > 0 && len(g.Belong) > 0 {
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
