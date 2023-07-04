package buxmodels

import (
	"time"

	"github.com/BuxOrg/bux-models/common"
)

type DraftTransaction struct {
	Model common.Model

	ID            string            `json:"id"`
	Hex           string            `json:"hex"`
	XpubID        string            `json:"xpub_id"`
	ExpiresAt     time.Time         `json:"expires_at"`
	Configuration TransactionConfig `json:"configuration"`
	Status        string            `json:"status"`
	FinalTxID     string            `json:"final_tx_id"`
}
