package main

// import (
// 	"fmt"

// 	"github.com/ZQingS/chess/internal/board"
// )

// func main() {
// var nextCommand string
// var chooseBelong string
// var chessesBoardWithBelong [8][8]string

// 	for {
// 		fmt.Printf("White or Black? ")
// 		fmt.Scanf("%s", &chooseBelong)

// 		if chooseBelong == "White" || chooseBelong == "Black" {
// 			break
// 		}
// 	}

// 	chesseBoard := board.InitBoard()

// 	board.InitChesses()
// 	board.SetChessInBoard()

// 	if chooseBelong == "White" {
// 		chessesBoardWithBelong = chesseBoard.GetChesses()
// 	} else {
// 		chessesBoardWithBelong = chesseBoard.GetReverseChesses()
// 	}

// 	for _, v := range chessesBoardWithBelong {
// 		for _, c := range v {
// 			fmt.Printf("|%s ", c)
// 		}
// 		fmt.Println("|")
// 	}

// 	for {
// 		fmt.Printf("%s> ", chooseBelong)
// 		fmt.Scanf("%s", &nextCommand)
// 	}
// }
