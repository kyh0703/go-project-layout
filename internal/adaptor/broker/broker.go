package broker

import (
	"github.com/kyh0703/go-project-layout/internal/domain/vo/event"
	"github.com/kyh0703/go-project-layout/pkg/kafka"
)

type KafkaBroker struct {
	manager *kafka.Manager
}

func ProvideBroker(
	manager *kafka.Manager,
) Broker {
	return &KafkaBroker{
		manager: manager,
	}
}

// Produce implements Broker.
func (b *KafkaBroker) Produce(events ...event.Event) {
	b.manager.Produce(events...)
}

// Consume implements Broker.
func (b *KafkaBroker) Consume(id string) EventHandler {
	handler := kafka.NewHandler()
	b.manager.Register(id, handler)
	return handler
}

// Close implements Broker.
func (b *KafkaBroker) Close(id string, handler EventHandler) {
	b.manager.Deregister(id, handler)
}
