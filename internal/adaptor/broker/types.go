package broker

import (
	"context"
	"time"

	cloud "github.com/cloudevents/sdk-go/v2"
	"github.com/gofrs/uuid"
	"github.com/kyh0703/go-project-layout/internal/domain/vo/event"
	"github.com/kyh0703/go-project-layout/pkg/kafka"
)

type Broker interface {
	Producer
	Consumer
}

type Producer interface {
	Produce(events ...event.Event)
}

type Consumer interface {
	Consume(id string) EventHandler
	Close(id string, handler EventHandler)
}

type EventHandler interface {
	ID() uuid.UUID
	Close()
	Subscribe(expects ...string)
	SetOnCallTerminate(kafka.EventFunc)
	SetOnCallDisconnect(kafka.EventFunc)
	SetOnSipAlert(kafka.EventFunc)
	SetOnSipConnect(kafka.EventFunc)
	SetOnSipRelease(kafka.EventFunc)
	SetOnSipTransfer(kafka.EventFunc)
	SetOnSipJoin(kafka.EventFunc)
	SetOnSipSwitch(kafka.EventFunc)
	SetOnSipHeld(kafka.EventFunc)
	SetOnSipRetrieve(kafka.EventFunc)
	SetOnMediaPlayDone(kafka.EventFunc)
	OnEvent(event cloud.Event)
	SetTimer(time.Duration)
	Poll(ctx context.Context) error
}
