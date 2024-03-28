package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ch := make(chan int, 1)

	conn, err := net.Dial("tcp", "127.0.0.1:9090")

	if err != nil {
		fmt.Println("client dial err =", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	var sessionId string
	var belong string

	go func() {
		for {
			<-ch

			if err := HandleServerMessage(conn); err != nil {
				fmt.Println("HandleServerMessage err ", err)
			}
		}
	}()

	for {
		var pre string
		var after string

		if len(sessionId) == 0 {
			pre = "NoSessionID "
			fmt.Println("Choose Your SessionID, Please Insert Like `Room1234`")
		} else {
			pre = fmt.Sprintf("SessionID %s ", sessionId)
		}

		if len(belong) == 0 && len(sessionId) > 0 {
			after = "My> "
			fmt.Println("Choose Your Belong, Please W OR B, White OR BLACK? (W/B)")
		} else {
			after = fmt.Sprintf("%s> ", belong)
		}

		if len(sessionId) > 0 && len(belong) > 0 {
			ch <- 1
		}

		fmt.Print(pre, after)

		if line, ok := handleLineMessage(reader); !ok {
			break
		} else {
			if len(line) == 0 {
				continue
			}
			currentSessionId := sessionId
			sessionId = handleSessionIdMessage(line, sessionId)
			belong = handleBelongIdMessage(line, belong, currentSessionId)

			if err := handleClientMessage(conn, line, sessionId, belong); err != nil {
				fmt.Println("conn write err =", err)
				break
			}

			if err := HandleServerMessage(conn); err != nil {
				fmt.Println("HandleServerMessage err ", err)
				break
			}
		}
	}
}
