package metric

type Metric interface {
	IncreaseRequestCount()
	IncreaseSuccessCount()
	IncreaseFailCount()
	IncreaseTx(tenant, txType string)
	DecreaseTx(tenant, txType string)
}
