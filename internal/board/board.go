package board

type Board struct {
	chesses [8][8]string
}

func InitBoard() *Board {
	var chesses [8][8]string

	var row [8]string = [8]string{"*", "*", "*", "*", "*", "*", "*", "*"}

	for i := 0; i < 8; i++ {
		chesses[i] = row
	}

	return &Board{
		chesses: chesses,
	}
}

func (b *Board) SetChesses(chesses [8][8]string) *Board {
	return &Board{
		chesses: chesses,
	}
}

func (b *Board) HandlePostionCommand(positionCommand string) [8][8]string {
	// firstPostion := string(positionCommand[0])

	// runeArray := []rune(firstPostion)

	// if unicode.IsUpper(runeArray[0]) {

	// }

	return b.GetChesses()
}

func (b *Board) MoveChess(startX, startY int, endPostion [2]int) *Board {
	chess := b.chesses[startX-1][startY-1]

	b.RemoveChess(startX-1, startY-1)
	b.SetChess(endPostion, chess)

	return b
}

func (b *Board) SetChess(position [2]int, icon string) *Board {
	b.chesses[position[0]-1][position[1]-1] = icon

	return b
}

func (b *Board) RemoveChess(x, y int) *Board {
	b.chesses[x][y] = "*"

	return b
}

func (b *Board) GetChess(x, y int) string {
	return b.chesses[x][y]
}

func (b *Board) GetChesses() [8][8]string {
	return b.chesses
}

func (b *Board) GetReverseChesses() [8][8]string {
	var chesses [8][8]string

	for x, v := range b.chesses {
		for y, c := range v {
			chesses[8-x-1][8-y-1] = c
		}
	}
	return chesses
}
