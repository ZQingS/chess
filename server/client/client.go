package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/ZQingS/chess/internal/protocol"
)

var connClose int = 0 // 1/0

func main() {
	receive := &protocol.Receive{}

	conn, err := net.Dial("tcp", "127.0.0.1:9090")

	if err != nil {
		fmt.Println("client dial err =", err)
		return
	}

	defer conn.Close()
	defer func() {
		connClose = 1
	}()

	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			if connClose == 0 {
				if err := HandleServerMessage(conn, receive); err != nil {
					fmt.Println("HandleServerMessage err ", err)

					connClose = 1

					break
				}
			}
		}
	}()

	for {
		lineDoStatus := true

		if connClose == 1 {
			break
		}

		if line, ok := handleLineMessage(reader); !ok {
			break
		} else {
			lineDoStatus = false

			if len(line) == 0 {
				continue
			}

			quit := receive.HandleClientQuitMessage(line)

			if !quit {
				lineDoStatus = receive.HandleClientSessionIdMessage(line, lineDoStatus)
				lineDoStatus = receive.HandleClientBelongMessage(line, lineDoStatus)
				lineDoStatus = receive.HandleClientPostionMessage(line, lineDoStatus)
				receive.HandleClientOtherMessage(line, lineDoStatus)
			}

			_, err := conn.Write(receive.ToByte())

			if err != nil {
				fmt.Println("Err Client Write, ", err)
			}

			if quit {
				break
			}
		}
	}
}
