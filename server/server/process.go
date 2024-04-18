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
	return len(gp.White.Belong) > 0 && len(gp.Black.Belong) > 0 &&
		len(gp.White.SessionID) > 0 && len(gp.Black.Belong) > 0
}

func (gp *GameProcess) OtherProcess(currentProcess Process) Process {
	if currentProcess.Belong == "W" {
		return gp.Black
	}

	return gp.White
}

func (gp *GameProcess) HandleBoardMessage(sessionId string, action string) (whiteErr error, BlackErr error) {
	var resBoard = &board.Board{}

	if len(gp.Board[0][0]) == 0 {
		game := NewORContinueGame(sessionId)

		resBoard = game.ChesseBoard
	} else {
		resBoard = resBoard.SetChesses(gp.Board)
	}

	gp.Board = resBoard.GetChesses()

	whiteErr = gp.White.HandleBoardMessage(gp.GetRound(), resBoard.GetChesses())
	BlackErr = gp.Black.HandleBoardMessage(gp.GetRound(), resBoard.GetReverseChesses())

	return whiteErr, BlackErr
}

func (gp *GameProcess) HandlePositionMessage(belong string, action string, position string) {
	if action == "position" && belong == gp.GetRound() {
		var resBoard = &board.Board{}

		resBoard = resBoard.SetChesses(gp.Board)

		gp.Board = resBoard.HandlePostionCommand(position)

		gp.HandleRoundMessage(true, belong)
	}
}

func (gp *GameProcess) GetRound() string {
	if gp.Round == 1 {
		return "B"
	} else {
		return "W"
	}
}

func (gp *GameProcess) HasChooseBelongInGame(belong string) bool {
	if belong == "W" {
		return len(gp.White.Belong) > 0
	} else {
		return len(gp.Black.Belong) > 0
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

func (g *Process) HandleFirstConnectMessage() {
	send := &protocol.Send{}
	send.OK = "true"
	send.Msg = "Choose Your SessionID, Please Insert Like `Room1234`"
	send.EndPoint = "NoSessionID My > "

	g.CurrentSend = send
}

func (g *Process) HandleQuitMessage() (error, bool) {
	if g.receive.Action == "quit" {
		g.CurrentSend.OK = "false"
		g.CurrentSend.Msg = "Quit the Game"

		_, err := g.ConnPool.Write(g.CurrentSend.ToByte())

		if err != nil {
			fmt.Println("HandleQuitMessage conn write err =", err)

			return err, true
		}

		return nil, true
	}

	return nil, false
}

func (g *Process) HandleBoardMessage(round string, board [8][8]string) error {
	if g.Belong == round {
		g.CurrentSend.Msg = "Your Turn"
		g.CurrentSend.EndPoint = g.handleEndPoint()
		g.CurrentSend.Board = board
	} else {
		g.CurrentSend.Msg = "Other Turn, Please Wait"
		g.CurrentSend.EndPoint = ""
	}

	_, err := g.ConnPool.Write(g.CurrentSend.ToByte())

	if err != nil {
		fmt.Println("HandlePossessiveMessage conn write err =", err)

		return err
	}

	return nil
}

func (g *Process) HandleBelongMessage() error {
	if g.receive.Action == "belong" {
		g.Belong = g.receive.Belong
		g.CurrentSend.OK = "true"
		g.CurrentSend.Msg = fmt.Sprintf("Choose Belong %s", g.Belong)

		_, err := g.ConnPool.Write(g.CurrentSend.ToByte())

		if err != nil {
			fmt.Println("HandleBelongMessage conn write err =", err)

			return err
		}
	}

	return nil
}

func (g *Process) HandleOtherMessage() error {
	if g.receive.Action == "other" {
		g.CurrentSend.OK = "true"
		g.CurrentSend.Msg = "UnKnow Action"
		g.CurrentSend.EndPoint = g.handleEndPoint()

		_, err := g.ConnPool.Write(g.CurrentSend.ToByte())

		if err != nil {
			fmt.Println("HandleOtherMessage conn write err =", err)

			return err
		}
	}

	return nil
}

func (g *Process) handleEndPoint() string {
	pre := "NoSessionID"
	after := "My"

	if len(g.SessionID) > 0 {
		pre = fmt.Sprintf("SessionID %s", g.SessionID)
	}

	if len(g.Belong) > 0 {
		after = g.Belong
	}

	return fmt.Sprintf("%s %s > ", pre, after)
}

func (g *Process) HandlePossessiveMessage() error {
	if g.receive.Action == "belong" {
		g.Belong = ""
		g.CurrentSend.OK = "true"
		g.CurrentSend.Msg = fmt.Sprintf("Belong %s Has Been Possessive", g.receive.Belong)
		g.CurrentSend.EndPoint = fmt.Sprintf("SessionID %s My > ", g.SessionID)
		g.CurrentSend.Action = "revert:belong"

		_, err := g.ConnPool.Write(g.CurrentSend.ToByte())

		if err != nil {
			fmt.Println("HandlePossessiveMessage conn write err =", err)

			return err
		}
	}

	return nil
}

func (g *Process) HandleSessionMessage() error {
	if g.receive.Action == "session" {
		g.SessionID = g.receive.SessionID
		g.CurrentSend.OK = "true"
		g.CurrentSend.Msg = fmt.Sprintf("Join The Game %s\nChoose Your Belong, Please W OR B, White OR BLACK? (W/B)", g.SessionID)
		g.CurrentSend.EndPoint = fmt.Sprintf("SessionID %s My > ", g.SessionID)

		_, err := g.ConnPool.Write(g.CurrentSend.ToByte())

		if err != nil {
			fmt.Println("HandleSessionMessage conn write err =", err)

			return err
		}
	}

	return nil
}

func (g *Process) HandleWaitMessage() error {
	g.CurrentSend.OK = "true"
	g.CurrentSend.Msg = fmt.Sprintln("Wait Other...")

	_, err := g.ConnPool.Write(g.CurrentSend.ToByte())

	if err != nil {
		fmt.Println("HandleWaitMessage conn write err =", err)

		return err
	}

	return nil
}
