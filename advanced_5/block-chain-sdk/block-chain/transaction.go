package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

type Transaction struct {
	Hash    []byte
	Payload []byte
}

type Hashable interface {
	Hash() ([]byte, error)
}

func NewTransaction(payload []byte) *Transaction {
	tx := Transaction{
		Payload: payload,
	}
	tx.Hash, _ = GetTransactionHash(tx)

	return &tx
}

func GetTransactionHash(tx Transaction) ([]byte, error) {
	txBytes, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}
	h := sha256.New()
	h.Write(txBytes)

	return h.Sum(nil), nil
}

func (t *Transaction) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, t.Payload)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
