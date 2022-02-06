package block

import (
	"encoding/json"
	"time"

	"block-chain-inspection/internal"
)

type Data struct {
	Message string `json:"message"`
}

type Block struct {
	Data     Data      `json:"data"`
	Date     time.Time `json:"date"`
	PrevHash []byte    `json:"prev_hash"`
}

func NewData(message string) *Data {
	return &Data{
		Message: message,
	}
}

func NewBlock(data *Data, prevHash []byte) *Block {
	return &Block{
		Data:     *data,
		Date:     time.Now(),
		PrevHash: prevHash,
	}
}

func (b *Block) JSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Block) Sha256() ([]byte, error) {
	buf, err := b.JSON()
	if err != nil {
		return nil, err
	}

	return internal.CryptoSha256(buf), nil
}
