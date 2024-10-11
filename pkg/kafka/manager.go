package kafka

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kyh0703/go-project-layoutinternal/domain/vo/event"
	"github.com/kyh0703/go-project-layoutpkg/eventhandler"
	"gitlab.com/ipron-core/call/configs"
	"gitlab.com/ipron-ne/iCore/ilog"
	"gitlab.com/ipron-ne/ievent"

	cloud "github.com/cloudevents/sdk-go/v2"
	callevt "gitlab.com/ipron-ne/ievent/call"
	mediaevt "gitlab.com/ipron-ne/ievent/media"
	sipevt "gitlab.com/ipron-ne/ievent/sip"
)

type Manager struct {
	log         ilog.Log
	lock        sync.RWMutex
	wg          sync.WaitGroup
	ctx         context.Context
	cancel      context.CancelFunc
	subjectMap  map[string]eventhandler.Subject
	producerMap map[string]*Producer
	consumerMap map[string]*Consumer
}

func ProvideManager() *Manager {
	ctx, cancel := context.WithCancel(context.Background())
	manager := &Manager{
		log:         ilog.NewLog("kafka-manager"),
		ctx:         ctx,
		cancel:      cancel,
		subjectMap:  make(map[string]eventhandler.Subject),
		producerMap: make(map[string]*Producer),
		consumerMap: make(map[string]*Consumer),
	}
	manager.InitLoop()
	go manager.poll()
	return manager
}

func (m *Manager) InitLoop() {
	m.log.Info("init kafka, broker: %v", configs.Env.KafkaBroker)
reconnect_producer:
	if err := m.initProducer(); err != nil {
		time.Sleep(time.Second)
		goto reconnect_producer
	}
reconnect_consumer:
	if err := m.initConsumer(); err != nil {
		time.Sleep(time.Second)
		goto reconnect_consumer
	}
}

func (m *Manager) initConsumer() error {
	m.log.Info("- init kafka consumer")

	call := NewConsumer(m.ctx, ievent.CloudTopicCall, m.Consume)
	if err := call.Connect(configs.Env.KafkaBroker); err != nil {
		call.Close()
		return err
	}

	media := NewConsumer(m.ctx, ievent.CloudTopicMedia, m.Consume)
	if err := media.Connect(configs.Env.KafkaBroker); err != nil {
		media.Close()
		return err
	}

	sip := NewConsumer(m.ctx, ievent.CloudTypeSip, m.Consume)
	if err := sip.Connect(configs.Env.KafkaBroker); err != nil {
		sip.Close()
		return err
	}

	m.registerConsumer(call)
	m.registerConsumer(media)
	m.registerConsumer(sip)
	return nil
}

func (m *Manager) initProducer() error {
	m.log.Info("- init kafka producer")
	var (
		call     = NewProducer(m.ctx, ievent.CloudTopicCall)
		tracking = NewProducer(m.ctx, ievent.CloudTopicTracking)
	)
	if err := call.Connect(configs.Env.KafkaBroker); err != nil {
		call.Close()
		return err
	}
	if err := tracking.Connect(configs.Env.KafkaBroker); err != nil {
		tracking.Close()
		return err
	}
	m.registerProducer(call)
	m.registerProducer(tracking)
	return nil
}

func (m *Manager) registerProducer(p *Producer) {
	m.producerMap[p.Topic()] = p
}

func (m *Manager) registerConsumer(c *Consumer) {
	m.consumerMap[c.Topic()] = c
}

// Close Event Manager.
func (m *Manager) Close() {
	m.cancel()
	m.wg.Wait()
}

// Register register event listener.
func (m *Manager) Register(id string, l eventhandler.Listener) {
	m.lock.Lock()
	defer m.lock.Unlock()

	subject, ok := m.subjectMap[id]
	if !ok {
		subject = eventhandler.NewDispatcher(id)
	}
	subject.Register(l)
	m.subjectMap[id] = subject
}

// Deregister is delete subject.
func (m *Manager) Deregister(id string, l eventhandler.Listener) {
	m.lock.Lock()
	defer m.lock.Unlock()

	defer l.Close()

	// get dispatcher
	subject, ok := m.subjectMap[id]
	if !ok {
		return
	}

	// detach listener
	subject.Deregister(l)

	// if listener count is zero, delete subject
	if subject.ListenerCount() == 0 {
		delete(m.subjectMap, id)
	}
}

func (m *Manager) Consume(_ context.Context, ce cloud.Event) {
	var callID string
	switch ce.Type() {
	case ievent.CloudTypeCallTerminated:
		var data callevt.Terminated
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeCallDisconnected:
		var data callevt.Disconnected
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeMediaPlayDone:
		var data mediaevt.PlayDone
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeSipAlerted:
		var data sipevt.Alerted
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeSipConnected:
		var data sipevt.Connected
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeSipHeld:
		var data sipevt.Held
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeSipRetrieved:
		var data sipevt.Retrieved
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeSipReleased:
		var data sipevt.Released
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeSipTransferred:
		var data sipevt.Transferred
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeSipJoined:
		var data sipevt.Joined
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	case ievent.CloudTypeSipConfSwitched:
		var data sipevt.ConfSwitched
		if err := ce.DataAs(&data); err != nil {
			return
		}
		callID = data.CallID
	default:
		return
	}

	m.lock.RLock()
	defer m.lock.RUnlock()

	subject, ok := m.subjectMap[callID]
	if !ok {
		return
	}
	subject.Notify(ce)
}

// ProduceEvent send the call event to kafka.
func (m *Manager) Produce(events ...event.Event) {
	for _, e := range events {
		ce := cloud.NewEvent()
		ce.SetID(uuid.New().String())
		ce.SetSource(ievent.CloudSourceCallID)
		ce.SetSpecVersion(ievent.CloudSpenV1)
		ce.SetType(e.Type())
		ce.SetSubject(e.Subject())
		ce.SetExtension("eventsubject", e.EventSubject())
		ce.SetExtension("tenantid", e.TenantID())
		ce.SetTime(time.Now())
		if err := ce.SetData(cloud.ApplicationJSON, e.MakeData()); err != nil {
			continue
		}
		producer, ok := m.producerMap[e.Topic()]
		if !ok {
			m.log.Error("not found producer, topic: %v", e.Topic())
			continue
		}
		producer.Send(ce)
	}
}

// print is a loop to print the event from kafka.
func (m *Manager) print() {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if len(m.subjectMap) == 0 {
		return
	}

	type garbageLog struct {
		CallId        string    `json:"call_id"`
		ListenerCount int       `json:"listener_cnt"`
		CreateAt      time.Time `json:"create_at"`
	}
	garbage := make([]garbageLog, 0)

	for _, subject := range m.subjectMap {
		garbage = append(garbage, garbageLog{
			subject.ID(),
			subject.ListenerCount(),
			subject.CreateAt(),
		})
	}

	b, err := json.Marshal(garbage)
	if err != nil {
		return
	}

	m.log.Info("%v", string(b))
}

// poll is a loop to poll the event from kafka.
func (m *Manager) poll() {
	m.wg.Add(1)
	defer m.wg.Done()

	for {
		select {
		case <-time.After(30 * time.Second):
			m.print()
		case <-m.ctx.Done():
			return
		}
	}
}
