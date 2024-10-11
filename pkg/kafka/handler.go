package kafka

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"gitlab.com/ipron-ne/ievent"

	cloud "github.com/cloudevents/sdk-go/v2"
	callevt "gitlab.com/ipron-ne/ievent/call"
	mediaevt "gitlab.com/ipron-ne/ievent/media"
	sipevt "gitlab.com/ipron-ne/ievent/sip"
)

const (
	timeout  = time.Second * 2
	MaxEvent = 30
)

var ErrTimeout = errors.New("event handler timeout")

type DoneFn = func()

type EventFunc = func(ctx context.Context, e cloud.Event, done DoneFn) error

type Handler struct {
	id      uuid.UUID
	timer   *time.Timer
	bus     chan cloud.Event
	expects []string
	doneCh  chan struct{}
	isDone  bool

	onCallTerminate  EventFunc
	onCallDisconncet EventFunc
	onSipAlert       EventFunc
	onSipConnect     EventFunc
	onSipRelease     EventFunc
	onSipTransfer    EventFunc
	onSipJoin        EventFunc
	onSipSwitch      EventFunc
	onSipHeld        EventFunc
	onSipRetrieve    EventFunc
	onMediaPlaydone  EventFunc
}

func NewHandler() *Handler {
	uuid, _ := uuid.NewV4()
	handler := &Handler{
		id:     uuid,
		bus:    make(chan cloud.Event, MaxEvent),
		doneCh: make(chan struct{}, 1), // Buffered to avoid locking up the event feed
	}
	return handler
}

func (h *Handler) ID() uuid.UUID {
	return h.id
}

func (h *Handler) SetOnCallTerminate(f EventFunc)  { h.onCallTerminate = f }
func (h *Handler) SetOnCallDisconnect(f EventFunc) { h.onCallDisconncet = f }
func (h *Handler) SetOnSipAlert(f EventFunc)       { h.onSipAlert = f }
func (h *Handler) SetOnSipConnect(f EventFunc)     { h.onSipConnect = f }
func (h *Handler) SetOnSipRelease(f EventFunc)     { h.onSipRelease = f }
func (h *Handler) SetOnSipTransfer(f EventFunc)    { h.onSipTransfer = f }
func (h *Handler) SetOnSipJoin(f EventFunc)        { h.onSipJoin = f }
func (h *Handler) SetOnSipSwitch(f EventFunc)      { h.onSipSwitch = f }
func (h *Handler) SetOnSipHeld(f EventFunc)        { h.onSipHeld = f }
func (h *Handler) SetOnSipRetrieve(f EventFunc)    { h.onSipRetrieve = f }
func (h *Handler) SetOnMediaPlayDone(f EventFunc)  { h.onMediaPlaydone = f }

func (h *Handler) SetTimer(dur time.Duration) {
	if h.timer == nil {
		h.timer = time.NewTimer(dur)
	} else {
		h.timer.Reset(dur)
	}
}

func (h *Handler) OnEvent(event cloud.Event) {
	h.bus <- event
}

func (h *Handler) IfInExpects(id string) bool {
	for _, expect := range h.expects {
		if id == expect {
			return true
		}
	}
	return false
}

func (h *Handler) Expects() []string {
	return h.expects
}

func (h *Handler) Subscribe(expects ...string) {
	h.expects = expects
}

func (h *Handler) Close() {
	if h.bus != nil {
		close(h.bus)
	}
	if h.timer != nil {
		h.timer.Stop()
	}
	if h.doneCh != nil {
		close(h.doneCh)
		h.doneCh = nil
	}
}

func (h *Handler) done() {
	h.isDone = true
	if h.doneCh != nil {
		h.doneCh <- struct{}{}
	}
}

func (h *Handler) ProcEvent(ctx context.Context, event cloud.Event) error {
	if h.isDone {
		return nil
	}
	switch event.Type() {
	case ievent.CloudTypeCallTerminated:
		if h.onCallTerminate == nil {
			return nil
		}
		return h.onCallTerminate(ctx, event, h.done)
	case ievent.CloudTypeCallDisconnected:
		if h.onCallDisconncet == nil {
			return nil
		}
		var data callevt.Disconnected
		event.DataAs(&data)
		if data.EventEndPointID != data.EndPointID {
			return nil
		}
		if !h.IfInExpects(data.EventEndPointID) {
			return nil
		}
		return h.onCallDisconncet(ctx, event, h.done)
	case ievent.CloudTypeSipAlerted:
		if h.onSipAlert == nil {
			return nil
		}
		var data sipevt.Alerted
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onSipAlert(ctx, event, h.done)
	case ievent.CloudTypeSipConnected:
		if h.onSipConnect == nil {
			return nil
		}
		var data sipevt.Connected
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onSipConnect(ctx, event, h.done)
	case ievent.CloudTypeSipHeld:
		if h.onSipHeld == nil {
			return nil
		}
		var data sipevt.Held
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onSipHeld(ctx, event, h.done)
	case ievent.CloudTypeSipRetrieved:
		if h.onSipRetrieve == nil {
			return nil
		}
		var data sipevt.Retrieved
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onSipRetrieve(ctx, event, h.done)
	case ievent.CloudTypeSipReleased:
		if h.onSipRelease == nil {
			return nil
		}
		var data sipevt.Released
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onSipRelease(ctx, event, h.done)
	case ievent.CloudTypeSipTransferred:
		if h.onSipTransfer == nil {
			return nil
		}
		var data sipevt.Transferred
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onSipTransfer(ctx, event, h.done)
	case ievent.CloudTypeSipJoined:
		if h.onSipJoin == nil {
			return nil
		}
		var data sipevt.Joined
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onSipJoin(ctx, event, h.done)
	case ievent.CloudTypeSipConfSwitched:
		if h.onSipSwitch == nil {
			return nil
		}
		var data sipevt.ConfSwitched
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onSipSwitch(ctx, event, h.done)
	case ievent.CloudTypeMediaPlayDone:
		if h.onMediaPlaydone == nil {
			return nil
		}
		var data mediaevt.PlayDone
		event.DataAs(&data)
		if !h.IfInExpects(data.ConnectionID) {
			return nil
		}
		return h.onMediaPlaydone(ctx, event, h.done)
	default:
		return nil
	}
}

func (h *Handler) Poll(ctx context.Context) error {
	if h.timer == nil {
		h.timer = time.NewTimer(timeout)
	}
	for {
		select {
		case <-h.doneCh:
			return nil
		case <-ctx.Done():
			return ctx.Err()
		case <-h.timer.C:
			return ErrTimeout
		case event := <-h.bus:
			if err := h.ProcEvent(ctx, event); err != nil {
				return err
			}
		}
	}
}
