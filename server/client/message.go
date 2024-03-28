package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/ZQingS/chess/internal/protocol"
)

func HandleServerMessage(conn net.Conn) error {
	var buf [512]byte

	n, err := conn.Read(buf[:])

	if err != nil {
		return err
	}

	send := protocol.InitSend(buf[:n])

	fmt.Println(send.Msg)

	if len(send.Board[0][0]) > 0 {
		for _, v := range send.Board {
			for _, c := range v {
				fmt.Printf("|%s ", c)
			}
			fmt.Println("|")
		}
	}

	return nil
}

func handleClientMessage(conn net.Conn, line string, sessionId string, belong string) error {
	var serverReceive *protocol.Receive

	if strings.Contains(line, "Room") {
		serverReceive = &protocol.Receive{
			Action:    "session",
			Belong:    "",
			SessionID: strings.Replace(line, "Room", "", 1),
			Position:  "",
		}
	} else if line == "quit" || line == "exit" {
		serverReceive = &protocol.Receive{
			Action:    "quit",
			Belong:    "",
			SessionID: sessionId,
			Position:  "",
		}
	} else if (line == "W" || line == "B") && (belong == "White" || belong == "Black") && len(sessionId) > 0 {
		serverReceive = &protocol.Receive{
			Action:    "belong",
			Belong:    line,
			SessionID: sessionId,
			Position:  "",
		}
	} else if len(sessionId) > 0 && len(belong) > 0 {
		serverReceive = &protocol.Receive{
			Action:    "position",
			Belong:    line,
			SessionID: sessionId,
			Position:  line,
		}
	} else {
		serverReceive = &protocol.Receive{
			Action:    "other",
			Belong:    "",
			SessionID: "",
			Position:  line,
		}
	}

	_, err := conn.Write(serverReceive.ToByte())

	return err
}

func handleLineMessage(reader *bufio.Reader) (string, bool) {
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("reading string err =", err)
	}

	line = strings.Trim(line, "\r\n")

	return line, true
}

func handleSessionIdMessage(line string, sessionId string) string {
	if len(sessionId) > 0 {
		return sessionId
	}

	if strings.Contains(line, "Room") {
		return strings.Replace(line, "Room", "", 1)
	}

	return sessionId
}

func handleBelongIdMessage(line string, belong string, sessionId string) string {
	if len(sessionId) == 0 {
		return belong
	}

	if len(belong) > 0 {
		return belong
	}

	if strings.Contains(line, "W") {
		return "White"
	} else if strings.Contains(line, "B") {
		return "Black"
	}

	return belong
}
