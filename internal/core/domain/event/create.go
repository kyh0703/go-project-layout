package event

import (
	"encoding/json"
	"time"

	"gitlab.com/ipron-core/call/internal/core/domain/entity"
	"gitlab.com/ipron-ne/iCore/ilog"
	"gitlab.com/ipron-ne/ievent"

	callevt "gitlab.com/ipron-ne/ievent/call"
)

type create struct {
	call entity.Call
	data callevt.Create
}

func NewCreate(call entity.Call) Event {
	return &create{call: call}
}

func (e *create) Topic() string {
	return ievent.CloudTopicCall
}

func (e *create) Type() string {
	return ievent.CloudTypeCallCreate
}

func (e *create) Subject() string {
	return subjectKey + e.call.ID
}

func (e *create) EventSubject() string {
	return eventSubjectKey + e.call.ID
}

func (e *create) TenantID() string {
	return e.call.TenantID
}

func (e *create) MakeData() []byte {
	e.data.Event = ievent.CloudEventCreate
	e.data.UCID = e.call.UCID
	e.data.TenantID = e.call.TenantID
	e.data.CallID = e.call.ID
	e.data.ANI = e.call.ANI
	e.data.DNIS = e.call.DNIS
	e.data.OriginCallNumber = e.call.OriginCallNumber
	e.data.RedirectNumber = e.call.RedirectNumber
	e.data.MediaType = e.call.MediaType
	e.data.CallType = e.call.CallType
	e.data.CallSubType = e.call.CallSubType
	e.data.Category = e.call.Category
	e.data.RefCallID = e.call.Event.ReferenceCallID
	e.data.UUI = e.call.UUI
	e.data.UEI = e.call.UEI
	e.data.CreateTime = e.call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()

	byteData, err := json.Marshal(e.data)
	if err != nil {
		return nil
	}

	ilog.Debug("%v", string(byteData))
	return byteData
}
