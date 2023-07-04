package buxmodels

import (
	"time"

	"github.com/BuxOrg/bux-models/common"
)

type BlockHeader struct {
	Model common.Model

	ID                string    `json:"id"`
	Height            uint32    `json:"height"`
	Time              uint32    `json:"time"`
	Nonce             uint32    `json:"nonce"`
	Version           uint32    `json:"version"`
	HashPreviousBlock string    `json:"hash_previous_block"`
	HashMerkleRoot    string    `json:"hash_merkle_root"`
	Bits              string    `json:"bits"`
	Synced            time.Time `json:"synced"`
}
