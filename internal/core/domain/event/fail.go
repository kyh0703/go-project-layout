package event

import (
	"encoding/json"
	"strconv"
	"time"

	"gitlab.com/ipron-core/call/internal/core/domain/entity"
	"gitlab.com/ipron-ne/iCore/ilog"
	"gitlab.com/ipron-ne/ievent"

	callevt "gitlab.com/ipron-ne/ievent/call"
)

type fail struct {
	self entity.Leg
	data callevt.Terminated
}

func NewFailed(self entity.Leg) Event {
	return &fail{self: self}
}

func (e *fail) Type() string {
	return ievent.CloudTypeCallTerminated
}

func (*fail) Topic() string {
	return ""
}

func (e *fail) Subject() string {
	return subjectKey + e.self.Call().ID
}

func (e *fail) EventSubject() string {
	return e.self.PartType + "/" + e.self.PartID
}

func (e *fail) TenantID() string {
	return e.self.Call().TenantID
}

func (e *fail) MakeData() []byte {
	call := e.self.Call()
	e.data.Event = ievent.CloudEventTerminated
	e.data.UCID = call.UCID
	e.data.CallSeq = json.Number(strconv.Itoa(call.Seq))
	e.data.TenantID = call.TenantID
	e.data.CallID = call.ID
	e.data.MediaType = call.MediaType
	e.data.Category = call.Category
	e.data.CallType = call.CallType
	e.data.CallSubType = call.CallSubType
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginCallNumber = call.OriginCallNumber
	e.data.RedirectNumber = call.RedirectNumber
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()

	byteData, err := json.Marshal(e.data)
	if err != nil {
		return nil
	}

	ilog.Debug("%v", string(byteData))
	return byteData
}
