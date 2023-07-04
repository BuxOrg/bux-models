package buxmodels

import (
	"time"

	"github.com/BuxOrg/bux-models/common"
)

type AccessKey struct {
	Model common.Model

	ID        string    `json:"id"`
	XpubID    string    `json:"xpub_id"`
	RevokedAt time.Time `json:"revoked_at"`
	Key       string    `json:"key"`
}
