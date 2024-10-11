package txpool

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kyh0703/go-project-layout/configs"
	"github.com/kyh0703/go-project-layout/internal/app/transaction/dto"
)

// Transaction is an call transaction.
type Tx struct {
	id     string
	time   time.Time
	ctx    context.Context
	cancel context.CancelFunc
	expire time.Duration
}

func NewTx(inner *dto.TxConfigDto) *Tx {
	ctx, cancel := context.WithTimeout(context.Background(), configs.Env.TransactionTimeout)
	id := uuid.NewString()
	return &Tx{
		id:     id,
		time:   time.Now(),
		ctx:    ctx,
		cancel: cancel,
		expire: configs.Env.TransactionTimeout,
	}
}

func (tx *Tx) Close() {
	tx.cancel()
}
