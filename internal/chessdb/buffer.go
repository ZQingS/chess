package chessdb

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ZQingS/chess/internal/board"
)

type BufferData struct {
	Board   [8][8]string               `json:"board"`
	Chesses map[string]*board.ChessMan `json:"chesses"`
}

type Buffer struct {
	Storage Storage
	Data    BufferData
}

func (b *Buffer) DataExist() bool {
	filepath := fmt.Sprintf("./db/%s", b.Storage.SessionId)

	_, err := os.Stat(filepath)

	return !os.IsNotExist(err)
}

func (b *Buffer) GetData() error {
	data, err := b.Storage.GetData()

	if err != nil {
		return err
	}

	var bufdata BufferData

	json.Unmarshal(data, &bufdata)

	b.Data = bufdata

	return nil
}

func (b *Buffer) WriteData() error {
	byteData, err := json.Marshal(b.Data)

	if err != nil {
		return err
	}

	b.Storage.WriteTable(byteData)

	return nil
}
