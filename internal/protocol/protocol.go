package protocol

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

type Receive struct {
	Action    string // belong, session, position, quit
	Belong    string
	SessionID string
	Position  string
}

func InitReceive(data []byte) *Receive {
	receive := &Receive{}
	buf := bytes.NewBuffer([]byte{})

	buf.Write(data)

	res := gob.NewDecoder(buf)

	if err := res.Decode(receive); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return receive
}

func (r *Receive) ToByte() []byte {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	if err := enc.Encode(r); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return buf.Bytes()
}

func (r *Receive) HandleClientQuitMessage(line string) bool {
	if line == "quit" || line == "exit" {
		r.Action = "quit"

		return true
	}

	return false
}

func (r *Receive) HandleClientSessionIdMessage(line string, lineDoStatus bool) bool {
	if len(r.SessionID) == 0 && strings.Contains(line, "Room") && !lineDoStatus {
		r.SessionID = strings.Replace(line, "Room", "", 1)
		r.Action = "session"

		return true
	}

	return lineDoStatus
}

func (r *Receive) HandleClientBelongMessage(line string, lineDoStatus bool) bool {
	if (line == "W" || line == "B") && len(r.SessionID) > 0 && len(r.Belong) == 0 && !lineDoStatus {
		r.Action = "belong"
		r.Belong = line

		return true
	}

	return lineDoStatus
}

func (r *Receive) HandleClientPostionMessage(line string, lineDoStatus bool) bool {
	// 需要判断走法
	if len(r.SessionID) > 0 && len(r.Belong) > 0 && !lineDoStatus {
		r.Action = "position"
		r.Position = line

		return true
	}

	return lineDoStatus
}

func (r *Receive) HandleClientOtherMessage(line string, lineDoStatus bool) {
	if !lineDoStatus {
		r.Action = "other"
	}
}

type Send struct {
	OK       string
	Msg      string
	Board    [8][8]string
	EndPoint string
}

func InitSend(data []byte) *Send {
	send := &Send{}

	buf := bytes.NewBuffer(data)
	res := gob.NewDecoder(buf)

	if err := res.Decode(send); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return send
}

func (s *Send) ToByte() []byte {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	if err := enc.Encode(s); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return buf.Bytes()
}
