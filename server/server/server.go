package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/ZQingS/chess/internal/protocol"
)

func process(conn net.Conn) {
	var sessionId string
	var belong string
	var board [8][8]string

	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)

		var buf [128]byte

		n, err := reader.Read(buf[:])

		if err != nil {
			fmt.Println("read from conn failed, err: ", err)
			break
		}

		res, ok, currentSessionId, currentBelong, currentBoard := handleClientMessage(buf[:n], sessionId, belong, board)
		sessionId = currentSessionId
		belong = currentBelong
		board = currentBoard

		_, err = conn.Write(res)

		if !ok {
			break
		}

		if err != nil {
			fmt.Println("Write from conn failed, err: ", err)
			break
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")

	if err != nil {
		fmt.Println("Listen Failed, Err: ", err)

		return
	}

	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("Accept Failed, Err: ", err)

			continue
		}

		go process(conn)
	}
}

func handleClientMessage(btyes []byte, sessionId, belong string, board [8][8]string) ([]byte, bool, string, string, [8][8]string) {
	receive := protocol.InitReceive(btyes)
	var msg string
	var connect bool = true
	var ok string = "true"

	switch receive.Action {
	case "quit":
		msg = "Quit the Game"
		connect = false
		ok = "false"

	case "belong":
		belong = receive.Belong

		board = HandleBoardMessage(sessionId, belong, board, "")

		msg = "Choose Belong"
	case "session":
		sessionId = receive.SessionID

		msg = "Get Session"
	case "other":
		msg = "Error Command"
	case "position":
		board = HandleBoardMessage(sessionId, belong, board, receive.Position)

		msg = fmt.Sprintf("Last Command is %s", receive.Position)
	default:
		msg = "Error Command"
	}

	res := handleSendMessage(msg, ok, board)

	return res, connect, sessionId, belong, board
}

func handleSendMessage(msg string, ok string, Board [8][8]string) []byte {
	send := &protocol.Send{
		OK:    ok,
		Msg:   msg,
		Board: Board,
	}

	return send.ToByte()
}
