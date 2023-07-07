package buxmodels

import "errors"

var ErrDraftNotFound = errors.New("corresponding draft transaction not found")

var ErrMissingXPriv = errors.New("missing xPriv key")
