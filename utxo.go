package buxmodels

import (
	"time"

	"github.com/BuxOrg/bux-models/common"
)

type UtxoPointer struct {
	TransactionID string `json:"transaction_id"`
	OutputIndex   int    `json:"output_index"`
}

type Utxo struct {
	Model       common.Model
	UtxoPointer `json:",inline"`

	ID           string       `json:"id"`
	XpubID       string       `json:"xpub_id"`
	Satoshis     uint64       `json:"satoshis"`
	ScriptPubKey string       `json:"script_pub_key"`
	Type         string       `json:"type"`
	DraftID      string       `json:"draft_id"`
	ReservedAt   time.Time    `json:"reserved_at"`
	SpendingTxID string       `json:"spending_tx_id"`
	Transaction  *Transaction `json:"transaction"`
}
