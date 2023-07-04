package buxmodels

import "github.com/BuxOrg/bux-models/common"

type Xpub struct {
	// Base model
	Model common.Model

	// Model specific fields
	ID              string `json:"id"`
	CurrentBalance  uint64 `json:"current_balance"`
	NextInternalNum uint32 `json:"next_internal_num"`
	NextExternalNum uint32 `json:"next_external_num"`
}
