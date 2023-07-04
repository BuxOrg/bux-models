package buxmodels

type SyncConfig struct {
	Broadcast        bool `json:"broadcast"`
	BroadcastInstant bool `json:"broadcast_instant"`
	PaymailP2P       bool `json:"paymail_p2p"`
	SyncOnChain      bool `json:"sync_on_chain"`
}
