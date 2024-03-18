package board

import (
	"errors"
	"fmt"
)

var (
	Powers  [2]string = [2]string{"B", "W"}
	Symbols [6]string = [6]string{"K", "Q", "R", "N", "B", "P"}
)

type ChessMan struct {
	SymbolName string
	Belong     string
	Class      int
	Position   [2]int
	Icon       string
}

func InitChess(symbolName, belong string, class int, position [2]int, icon string) (*ChessMan, error) {
	if err := valiate(symbolName, belong, class, position); err != nil {
		return nil, err
	}

	chess := &ChessMan{
		SymbolName: symbolName,
		Belong:     belong,
		Class:      class,
		Position:   position,
		Icon:       icon,
	}

	return chess, nil
}

func valiate(symbolName, belong string, class int, position [2]int) error {
	if err := valiatePosition(position); err != nil {
		return err
	}

	if err := valiateBelong(belong); err != nil {
		return err
	}

	if err := valiateSymbol(symbolName); err != nil {
		return err
	}

	if err := valiateClass(class); err != nil {
		return err
	}

	return nil
}

func valiatePosition(position [2]int) error {
	if position[0] <= 8 && position[0] >= 1 {
		return nil
	}

	if position[1] <= 8 && position[1] >= 1 {
		return nil
	}

	return errors.New(fmt.Sprintf("valiatePosition valiateERROR: X: %d, Y: %d", position[0], position[1]))
}

func valiateBelong(belong string) error {
	for _, element := range Powers {
		if belong == element {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Belong valiate ERROR: %s", belong))
}

func valiateSymbol(symbolName string) error {
	for _, element := range Symbols {
		if symbolName == element {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("valiateSymbol valiateERROR: %s", symbolName))
}

func valiateClass(class int) error {
	return nil
	// if class <= 16 && class >= 1 {
	// 	return nil
	// }

	// return errors.New(fmt.Sprintf("valiateClass valiateERROR: %s", class))
}
