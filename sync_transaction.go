package buxmodels

import (
	"time"

	"github.com/BuxOrg/bux-models/common"
)

type SyncTransaction struct {
	Model common.Model

	ID              string      `json:"id"`
	Configuration   SyncConfig  `json:"configuration"`
	LastAttempt     time.Time   `json:"last_attempt"`
	Results         SyncResults `json:"results"`
	BroadcastStatus string      `json:"broadcast_status"`
	P2PStatus       string      `json:"p2p_status"`
	SyncStatus      string      `json:"sync_status"`
}
