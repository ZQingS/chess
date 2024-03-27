package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/ZQingS/chess/internal/protocol"
)

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

func process(conn net.Conn, bufferPools chan string) {
	currentProcess := &Process{}

	defer conn.Close()

	for {
		receive, err := resolveReceive(conn)

		if err != nil {
			fmt.Println("conn err =", err)

			break
		}

		currentProcess.receive = receive

		if quit := currentProcess.HandleQuitMessage(); !quit {
			currentProcess.HandleSessionMessage()
			currentProcess.HandleBelongMessage()

			readyChanPool := currentProcess.HandleReadyChanPool()

			if len(readyChanPool) > 0 {
				if currentProcess.Start == 0 {
					// buffer := []string{}

					// for date := range bufferPools {
					// 	buffer
					// }

					for date := range bufferPools {
						bufferPools <- date

						if currentProcess.HaveOtherReadyChanPool(date) {
							currentProcess.HandleSatrtMessage(true)
						}
					}

					bufferPools <- readyChanPool
				}

				currentProcess.HandleBoardMessage()
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
	bufferPools := make(chan string)

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

		go process(conn, bufferPools)
	}
}

func handleClientMessage(btyes []byte) *protocol.Receive {
	receive := protocol.InitReceive(btyes)

	return receive
}
