package buxmodels

import "github.com/BuxOrg/bux-models/common"

type PaymailAddress struct {
	Model common.Model

	ID              string `json:"id"`
	XpubID          string `json:"xpub_id"`
	Alias           string `json:"alias"`
	Domain          string `json:"domain"`
	PublicName      string `json:"public_name"`
	Avatar          string `json:"avatar"`
	ExternalXpubKey string `json:"external_xpub_key"`
}
