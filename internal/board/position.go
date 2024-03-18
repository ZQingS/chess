package board

const (
	up    int = 0b0001
	down  int = 0b0010
	left  int = 0b0100
	right int = 0b1000
)

type Position struct {
	step      int
	direction int
}

// P {step 1, direction: []}
// startPostion [2, 0b0b0000100]
// array[0]: 1 ~ 8
func (p *Position) chessManWay(startPostion [2]int) [8]int {
	return [8]int{
		0b00001000,
		0b00001000,
		0b00001000,
		0b00001000,
		0b00001000,
		0b00001000,
		0b00001000,
		0b00001000,
	}
}

// func (p *Position) Move(startPostion [2]int) [2]int {

// }
