package spvwalletmodels

// Metadata is a bux ecosystem metadata model.
type Metadata map[string]interface{}

// XpubMetadata is a bux ecosystem xpub metadata model.
type XpubMetadata map[string]Metadata
