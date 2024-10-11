package txpool

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gitlab.com/ipron-core/call/configs"
	"gitlab.com/ipron-core/call/internal/app/flow"
	"gitlab.com/ipron-core/call/internal/app/flow/strategy"
	"gitlab.com/ipron-core/call/internal/app/transaction"
	"gitlab.com/ipron-core/call/internal/app/transaction/dto"
	"gitlab.com/ipron-core/call/internal/app/types"
	"gitlab.com/ipron-ne/iCore/ilog"
	"gitlab.com/ipron-ne/iCore/otrace"
)

// Transaction is an call transaction.
type Tx struct {
	id      string
	inner   *dto.TxConfigDto
	time    time.Time
	ctx     context.Context
	cancel  context.CancelFunc
	expire  time.Duration
	factory *strategy.CallStrategyFactory
	log     ilog.Log
}

func NewTx(inner *dto.TxConfigDto, factory *strategy.CallStrategyFactory) *Tx {
	ctx := otrace.ContextWithSpan(context.Background(), inner.Span)
	ctx, cancel := context.WithTimeout(ctx, configs.Env.TransactionTimeout)

	id := uuid.NewString()
	fields := ilog.Fields{
		"tx_id":   id,
		"call_id": inner.FirstCall.ID,
		"ucid":    inner.FirstCall.UCID,
	}

	return &Tx{
		id:      id,
		inner:   inner,
		time:    time.Now(),
		ctx:     ctx,
		cancel:  cancel,
		expire:  configs.Env.TransactionTimeout,
		factory: factory,
		log:     ilog.NewLog(inner.Type).WithFields(fields),
	}
}

func (tx *Tx) Close() {
	tx.cancel()
}

func (tx *Tx) Type() string {
	return tx.inner.Type
}

func (tx *Tx) holdTx() {
	var (
		strategy flow.CallflowStrategy
		result   *types.TxResult
	)
	tx.log.Level3("(strategy-hold) start")
	strategy = tx.factory.GetCallStrategy(flow.StrategyHold)
	result = strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-hold) end, %v", result.Results())
}

func (tx *Tx) unHoldTx() {
	var (
		strategy flow.CallflowStrategy
		result   *types.TxResult
	)
	tx.log.Level3("(strategy-unhold) start")
	strategy = tx.factory.GetCallStrategy(flow.StrategyUnHold)
	result = strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-unhold) end, %v", result.Results())
}

func (tx *Tx) singleTransferTx() {
	tx.log.Level3("(strategy-singletransfer) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategySingleTransfer)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-singletransfer) end, %v", result.Results())

	if result.IsEnd() {
		tx.log.Level3("(strategy-release) start")
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())
	}
}

func (tx *Tx) muteTransferTx() {
	tx.log.Level3("(strategy-mutetransfer) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategyMuteTransfer)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-mutetransfer) end: %v", result.Results())

	if result.IsEnd() {
		tx.log.Level3("(strategy-release) start: %v", len(tx.inner.FirstCall.Legs))
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())

		tx.log.Level3("(strategy-release) start: %v", len(tx.inner.SecondCall.Legs))
		tx.inner.FirstCall = tx.inner.SecondCall
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())
	}
}

func (tx *Tx) releaseTx() {
	tx.log.Level3("(strategy-release) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategyRelease)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-release) end: %v", result.Results())
}

func (tx *Tx) terminateTx() {
	tx.log.Level3("(strategy-release) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategyRelease)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-release) end: %v", result.Results())
}

func (tx *Tx) makeCallTx() {
	tx.log.Level3("(strategy-originate) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategyOriginate)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-originate) end: %v", result.Results())

	if result.IsEnd() || result.Error() != nil {
		tx.log.Level3("(strategy-release) start")
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())
		return
	}

	tx.log.Level3("(strategy-makecall) start")
	strategy = tx.factory.GetCallStrategy(flow.StrategyMakeCall)
	result = strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-makecall) end, %v", result.Results())

	if result.IsEnd() {
		tx.log.Level3("(strategy-release) start")
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())
	}
}

func (tx *Tx) routeTx() {
	tx.log.Level3("(strategy-route) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategyRoute)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-route) end, %v", result.Results())

	if result.IsEnd() {
		tx.log.Level3("(strategy-release) start")
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())
	}
}

func (tx *Tx) singleConferenceTx() {
	tx.log.Level3("(strategy-singleconference) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategySingleConference)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-singleconference) end, %v", result.Results())

	if result.IsEnd() {
		tx.log.Level3("(strategy-release) start")
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())
	}
}

func (tx *Tx) muteConferenceTx() {
	tx.log.Level3("(strategy-mute-conference) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategyMuteConference)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-mute-conference) end, %v", result.Results())

	if result.IsEnd() {
		tx.log.Level3("(strategy-release) start: %v", len(tx.inner.FirstCall.Legs))
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())

		tx.log.Level3("(strategy-release) start: %v", len(tx.inner.SecondCall.Legs))
		tx.inner.FirstCall = tx.inner.SecondCall
		tx.inner.FirstCall.End.Ment = result.Ment()
		strategy = tx.factory.GetCallStrategy(flow.StrategyRelease)
		strategy.Execute(tx.ctx, tx.inner)
		tx.log.Level3("(strategy-release) end: %v", result.Results())
	}
}

func (tx *Tx) joinTx() {
	tx.log.Level3("(strategy-join) start")
	strategy := tx.factory.GetCallStrategy(flow.StrategyJoin)
	result := strategy.Execute(tx.ctx, tx.inner)
	tx.log.Level3("(strategy-join) end, %v", result.Results())
}

func (tx *Tx) doTx() {
	switch tx.Type() {
	case transaction.TxTypeDidCall, transaction.TxTypeFlowCall, transaction.TxTypeDeviceCall, transaction.TxTypeMakeCallEx, transaction.TxTypeMakeCall:
		tx.makeCallTx()
	case transaction.TxTypeDeviceHold:
		tx.holdTx()
	case transaction.TxTypeDeviceUnHold:
		tx.unHoldTx()
	case transaction.TxTypeDeviceSingleTransfer, transaction.TxTypeSingleTransfer:
		tx.singleTransferTx()
	case transaction.TxTypeDeviceMuteTransfer, transaction.TxTypeMuteTransfer:
		tx.muteTransferTx()
	case transaction.TxTypeDeviceRelease, transaction.TxTypeRelease:
		tx.releaseTx()
	case transaction.TxTypeTerminate:
		tx.terminateTx()
	case transaction.TxTypeRoute:
		tx.routeTx()
	case transaction.TxTypeSingleConference:
		tx.singleConferenceTx()
	case transaction.TxTypeMuteConference:
		tx.muteConferenceTx()
	case transaction.TxTypeJoinCall:
		tx.joinTx()
	}
}
