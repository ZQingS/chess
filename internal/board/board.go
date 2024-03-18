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

func (b *Board) MoveChess(startX, startY int, endPostion [2]int, chess string) *Board {
	b.RemoveChess(startX, startY)
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
