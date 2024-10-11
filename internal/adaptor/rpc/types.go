package rpc

import (
	"context"
)


type Rpc interface {
	Connect(context.Context) error
	Close()
}
