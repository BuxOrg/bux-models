package buxmodels

import "time"

type TransactionConfig struct {
	ChangeDestinations         []*Destination       `json:"change_destinations"`
	ChangeStrategy             string               `json:"change_destinations_strategy"`
	ChangeMinimumSatoshis      uint64               `json:"change_minimum_satoshis"`
	ChangeNumberOfDestinations int                  `json:"change_number_of_destinations"`
	ChangeSatoshis             uint64               `json:"change_satoshis"`
	ExpiresIn                  time.Duration        `json:"expires_in"`
	Fee                        uint64               `json:"fee"`
	FeeUnit                    *FeeUnit             `json:"fee_unit"`
	FromUtxos                  []*UtxoPointer       `json:"from_utxos"`
	IncludeUtxos               []*UtxoPointer       `json:"include_utxos"`
	Inputs                     []*TransactionInput  `json:"inputs"`
	Outputs                    []*TransactionOutput `json:"outputs"`
	SendAllTo                  *TransactionOutput   `json:"send_all_to"`
	Sync                       *SyncConfig          `json:"sync"`
}

type TransactionInput struct {
	Utxo        `json:",inline"`
	Destination Destination `json:"destination"`
}

type TransactionOutput struct {
	OpReturn     *OpReturn       `json:"op_return,omitempty"`
	PaymailP4    *PaymailP4      `json:"paymail_p4,omitempty"`
	Satoshis     uint64          `json:"satoshis"`
	Script       string          `json:"script"`
	Scripts      []*ScriptOutput `json:"scripts,omitempty"`
	To           string          `json:"to"`
	UseForChange bool            `json:"use_for_change"`
}

type MapProtocol struct {
	App  string                 `json:"app,omitempty"`
	Keys map[string]interface{} `json:"keys,omitempty"`
	Type string                 `json:"type,omitempty"`
}

type OpReturn struct {
	Hex         string       `json:"hex,omitempty"`          // Full hex
	HexParts    []string     `json:"hex_parts,omitempty"`    // Hex into parts
	Map         *MapProtocol `json:"map,omitempty"`          // MAP protocol
	StringParts []string     `json:"string_parts,omitempty"` // String parts
}

type PaymailP4 struct {
	Alias           string `json:"alias,omitempty"`
	Domain          string `json:"domain,omitempty"`
	FromPaymail     string `json:"from_paymail,omitempty"`
	Note            string `json:"note,omitempty"`
	PubKey          string `json:"pub_key,omitempty"`
	ReceiveEndpoint string `json:"receive_endpoint,omitempty"`
	ReferenceID     string `json:"reference_id,omitempty"`
	ResolutionType  string `json:"resolution_type,omitempty"`
}

type ScriptOutput struct {
	Address    string `json:"address,omitempty"`
	Satoshis   uint64 `json:"satoshis,omitempty"`
	Script     string `json:"script"`
	ScriptType string `json:"script_type"`
}
