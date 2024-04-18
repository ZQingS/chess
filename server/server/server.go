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

func msgSendToClient(process *Process) {
	bytes := process.CurrentSend.ToByte()
	_, err := process.ConnPool.Write(bytes)

	if err != nil {
		fmt.Println("conn msgSendToClient err =", err)
	}
}

func process(conn net.Conn) {
	currentProcess := &Process{}
	currentProcess.ConnPool = conn

	currentProcess.HandleFirstConnectMessage()
	msgSendToClient(currentProcess)

	defer close(conn)

	for {
		receive, err := resolveReceive(conn)
		send := &protocol.Send{}

		if err != nil {
			fmt.Println("conn err =", err)

			break
		}

		currentProcess.receive = receive
		currentProcess.CurrentSend = send

		if err, quit := currentProcess.HandleQuitMessage(); !quit {
			if err != nil {
				break
			}
			if err = currentProcess.HandleOtherMessage(); err != nil {
				break
			}
			if err = currentProcess.HandleSessionMessage(); err != nil {
				break
			}

			if processMap[currentProcess.SessionID] == nil {
				processMap[currentProcess.SessionID] = &GameProcess{}
			}

			if len(receive.Belong) > 0 && len(currentProcess.Belong) == 0 {
				if hasBelong := processMap[currentProcess.SessionID].HasChooseBelongInGame(receive.Belong); !hasBelong {
					if currentProcess.HandleBelongMessage(); err != nil {
						break
					}
				} else {
					if currentProcess.HandlePossessiveMessage(); err != nil {
						break
					}
				}
			}

			if readyChanPool := currentProcess.HandleReadyChanPool(); readyChanPool {
				processMap[currentProcess.SessionID].InitProcess(currentProcess)

				if processMap[currentProcess.SessionID].StartPlay() {
					processMap[currentProcess.SessionID].HandlePositionMessage(currentProcess.Belong, receive.Action, receive.Position)
					processMap[currentProcess.SessionID].HandleBoardMessage(currentProcess.SessionID, receive.Action)
				} else {
					currentProcess.HandleWaitMessage()
				}
			}
		}

	}
}

func handleClientMessage(btyes []byte) *protocol.Receive {
	receive := protocol.InitReceive(btyes)

	return receive
}

func close(conn net.Conn) {
	conn.Close()
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
