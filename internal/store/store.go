package store

import "context"

// Store defines all relevant get/set queries against the implemented storage backend.
type Store interface {
	CreateAtReceipt(context.Context, uint, string) error
	CreateTgReceipt(context.Context, int) error
	SetAtDelivered(context.Context, string) error
}
