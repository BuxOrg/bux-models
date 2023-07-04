package buxmodels

import "github.com/BuxOrg/bux-models/common"

type Transaction struct {
	Model                common.Model `json:",inline"`
	ID                   string       `json:"id"`
	Hex                  string       `json:"hex"`
	XpubInIDs            []string     `json:"xpub_in_ids"`
	XpubOutIDs           []string     `json:"xpub_out_ids"`
	BlockHash            string       `json:"block_hash"`
	BlockHeight          uint64       `json:"block_height"`
	Fee                  uint64       `json:"fee"`
	NumberOfInputs       uint32       `json:"number_of_inputs"`
	NumberOfOutputs      uint32       `json:"number_of_outputs"`
	DraftID              string       `json:"draft_id"`
	TotalValue           uint64       `json:"total_value"`
	OutputValue          int64        `json:"output_value"`
	Status               string       `json:"status"`
	TransactionDirection string       `json:"direction"`
}
