package chessdb

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// 1. 保存数组
// 2. 保存位置关系
// 3. 保存 session id 对应关系，session id 为索引

// 1. 索引文件
// 2. 表文件

type Storage struct {
	SessionId string
}

func (s *Storage) WriteTable(data []byte) error {
	encodeString := base64.StdEncoding.EncodeToString(data)

	filepath := fmt.Sprintf("./db/%s", s.SessionId)

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))

	if err != nil {
		return err
	}

	n, err := file.Write([]byte(encodeString))

	if err == nil && n < len([]byte(encodeString)) {
		err = io.ErrShortWrite
	}

	if err1 := file.Close(); err == nil {
		err = err1
	}

	return err
}

func (s *Storage) GetData() ([]byte, error) {
	file, err := os.Open(s.SessionId)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	res, err := base64.URLEncoding.DecodeString(string(bytes))

	return res, nil
}
