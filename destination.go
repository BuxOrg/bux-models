package buxmodels

import (
	"time"

	"github.com/BuxOrg/bux-models/common"
)

type Destination struct {
	Model common.Model

	ID            string    `json:"id"`
	XpubID        string    `json:"xpub_id"`
	LockingScript string    `json:"locking_script"`
	Type          string    `json:"type"`
	Chain         uint32    `json:"chain"`
	Num           uint32    `json:"num"`
	Address       string    `json:"address"`
	DraftID       string    `json:"draft_id"`
	Monitor       time.Time `json:"monitor"`
}
