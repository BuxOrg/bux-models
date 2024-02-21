package buxmodels

import "github.com/BuxOrg/bux-models/common"

type ContactStatus string

type Contact struct {
	common.Model

	ID string `json:"id"`

	XpubID string `json:"xpubID"`

	FullName string `json:"fullName"`

	Paymail string `json:"paymail"`

	PubKey string `json:"pubKey"`

	Status ContactStatus `json:"status"`
}
