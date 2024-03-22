package protocol

import (
	"bytes"
	"encoding/gob"
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
		panic(err)
	}

	return receive
}

func (r *Receive) ToByte() []byte {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	if err := enc.Encode(r); err != nil {
		panic(err)
	}

	return buf.Bytes()
}

type Send struct {
	OK    string
	Msg   string
	Board [8][8]string
}

func InitSend(data []byte) *Send {
	send := &Send{}

	buf := bytes.NewBuffer([]byte{})

	buf.Write(data)

	res := gob.NewDecoder(buf)

	if err := res.Decode(send); err != nil {
		panic(err)
	}

	return send
}

func (s *Send) ToByte() []byte {
	buf := new(bytes.Buffer)

	enc := gob.NewEncoder(buf)

	if err := enc.Encode(s); err != nil {
		panic(err)
	}

	return buf.Bytes()
}
