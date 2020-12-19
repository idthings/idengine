package datastore

import (
	"context"
)

// Interface interfaces to a datastore, and should allow easier mocking
type Interface interface {

	/*
	   Connect() returns:
	      true if the datastore is ready to be used
	      false otherwise
	*/
	Connect() bool

	StoreSecret(ctx context.Context, id string, secret string) error

	FetchSecret(ctx context.Context, id string) (string, error)
}
