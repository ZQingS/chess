package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/ZQingS/chess/internal/protocol"
)

func HandleServerMessage(conn net.Conn, receive *protocol.Receive) error {
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

	fmt.Printf(send.EndPoint)

	receive.HandleCallBackMessage(send.Action)

	return nil
}

func handleLineMessage(reader *bufio.Reader) (string, bool) {
	line, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("reading string err =", err)
	}

	line = strings.Trim(line, "\r\n")

	return line, true
}

func handleBelongIdMessage(line string, belong string, sessionId string) string {
	if len(sessionId) == 0 {
		return belong
	}

	if len(belong) > 0 {
		return belong
	}

	if strings.Contains(line, "W") {
		return "W"
	} else if strings.Contains(line, "B") {
		return "B"
	}

	return belong
}
