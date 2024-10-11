package txpool

import (
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/kyh0703/go-project-layout/internal/adaptor/metric"
	"github.com/kyh0703/go-project-layout/internal/app/flow/strategy"
	"github.com/kyh0703/go-project-layout/internal/app/transaction/dto"
)

type TxPool struct {
	metric     metric.Metric
	txMap      map[uuid.UUID]*Tx
	txCountMap map[string]int
	wg         sync.WaitGroup
	mutex      sync.RWMutex
	factory    *strategy.CallStrategyFactory
}

func ProvideTxPool(
	metric metric.Metric,
	factory *strategy.CallStrategyFactory,
) *TxPool {
	return &TxPool{
		metric:     metric,
		txMap:      make(map[uuid.UUID]*Tx),
		txCountMap: make(map[string]int),
		factory:    factory,
	}
}

func (pool *TxPool) PrintTx() {
	pool.mutex.RLock()
	defer pool.mutex.RUnlock()

	for _, tx := range pool.txMap {
		fmt.Println(tx)
	}
}

func (pool *TxPool) BeginTx(inner *dto.TxConfigDto) {
	// create tx data.
	tx := NewTx(inner, pool.factory)

	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	// compute tx type and count.
	uuid, _ := uuid.NewV4()
	pool.txCountMap[tx.Type()]++
	pool.metric.IncreaseTx(tx.inner.FirstCall.TenantID, tx.Type())
	pool.txMap[uuid] = tx
	pool.wg.Add(1)
	go func() {
		defer pool.wg.Done()
		defer pool.EndTx(uuid)
		tx.doTx()
	}()
}

func (pool *TxPool) EndTx(id uuid.UUID) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	tx, ok := pool.txMap[id]
	if !ok {
		return
	}
	tx.Close()
	pool.metric.DecreaseTx(tx.inner.FirstCall.TenantID, tx.Type())
	pool.txCountMap[tx.Type()]--
	delete(pool.txMap, id)
}

func (pool *TxPool) EndAllTx() {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	for _, tx := range pool.txMap {
		tx.Close()
	}
	pool.wg.Wait()
	for k := range pool.txMap {
		delete(pool.txMap, k)
	}
	for t := range pool.txCountMap {
		delete(pool.txCountMap, t)
	}
}
