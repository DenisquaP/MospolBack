package rabbit

import (
	"bytes"
	"encoding/json"
)

type Message struct {
	Reciever string `json:"reciever"`
	Message  string `json:"message"`
}

func (msg Message) Serialize() ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)

	return b.Bytes(), err
}

func Deserialize(b []byte) (Message, error) {
	var msg Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)

	return msg, err
}
