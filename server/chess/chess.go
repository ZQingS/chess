package main

import (
	"fmt"

	"github.com/Jameszhanqingsheng/chess/internal/board"
)

func main() {
	var nextCommand string
	chesseBoard = board.InitBoard()

	InitChesses()
	SetChessInBoard()

	for _, v := range chesseBoard.GetChesses() {
		for _, c := range v {
			fmt.Printf("|%s ", c)
		}
		fmt.Println("|")
	}

	for {
		fmt.Printf("White>")
		fmt.Scanf("%s", &nextCommand)

	}
}
