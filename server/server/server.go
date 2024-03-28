package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/ZQingS/chess/internal/protocol"
)

var processMap = make(map[string]*GameProcess)

func resolveReceive(conn net.Conn) (*protocol.Receive, error) {
	receive := &protocol.Receive{}

	reader := bufio.NewReader(conn)

	var buf [128]byte

	n, err := reader.Read(buf[:])

	if err != nil {
		return receive, err
	}

	receive = handleClientMessage(buf[:n])

	return receive, nil
}

func msgSendToOther(process *Process) {
	gameprocess := processMap[process.SessionID]

	var oatherProcess Process

	if process.Belong == "W" {
		oatherProcess = gameprocess.Black
	} else {
		oatherProcess = gameprocess.White
	}

	bytes := oatherProcess.CurrentSend.ToByte()

	fmt.Println(oatherProcess.ConnPool)

	_, err := oatherProcess.ConnPool.Write(bytes)

	if err != nil {
		fmt.Println("conn msgSendToOther err =", err)
	}
}

func process(conn net.Conn) {
	currentProcess := &Process{}
	currentProcess.ConnPool = conn
	fmt.Println(conn)

	defer conn.Close()

	for {
		receive, err := resolveReceive(conn)
		send := &protocol.Send{}

		if err != nil {
			fmt.Println("conn err =", err)

			break
		}

		currentProcess.receive = receive
		currentProcess.CurrentSend = send

		if quit := currentProcess.HandleQuitMessage(); !quit {
			currentProcess.HandleSessionMessage()
			currentProcess.HandleBelongMessage()

			if readyChanPool := currentProcess.HandleReadyChanPool(); readyChanPool {
				if processMap[currentProcess.SessionID] == nil {
					processMap[currentProcess.SessionID] = &GameProcess{}
				}

				processMap[currentProcess.SessionID].InitProcess(currentProcess)

				if processMap[currentProcess.SessionID].StartPlay() {
					processMap[currentProcess.SessionID].HandleBoardMessage(currentProcess.Belong, currentProcess.SessionID)
					msgSendToOther(currentProcess)
				}
			}
		}

		_, err = conn.Write(currentProcess.CurrentSend.ToByte())

		if err != nil {
			fmt.Println("conn write err =", err)

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

func handleClientMessage(btyes []byte) *protocol.Receive {
	receive := protocol.InitReceive(btyes)

	return receive
}
