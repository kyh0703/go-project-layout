package event

import (
	"encoding/json"
	"time"

	"gitlab.com/ipron-core/call/internal/core/domain/entity"
	"gitlab.com/ipron-ne/iCore/ilog"
	"gitlab.com/ipron-ne/ievent"

	callevt "gitlab.com/ipron-ne/ievent/call"
)

type terminated struct {
	call entity.Call
	data callevt.Terminated
}

func NewTerminated(call entity.Call) Event {
	return &terminated{call: call}
}

func (e *terminated) Topic() string {
	return ievent.CloudTopicCall
}

func (e *terminated) Type() string {
	return ievent.CloudTypeCallTerminated
}

func (e *terminated) Subject() string {
	return subjectKey + e.call.ID
}

func (e *terminated) EventSubject() string {
	return eventSubjectKey + e.call.ID
}

func (e *terminated) TenantID() string {
	return e.call.TenantID
}

func (e *terminated) MakeData() []byte {
	e.data.Event = ievent.CloudEventTerminated
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
	e.data.EndType = e.call.Event.EndType
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
