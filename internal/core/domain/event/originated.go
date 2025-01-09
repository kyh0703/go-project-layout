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

type originated struct {
	self entity.Leg
	data callevt.Originated
}

func NewOriginated(self entity.Leg) Event {
	return &originated{self: self}
}

func (e *originated) Topic() string {
	return ievent.CloudTopicCall
}

func (e *originated) Type() string {
	return ievent.CloudTypeCallOriginated
}

func (e *originated) Subject() string {
	return subjectKey + e.self.Call().ID
}

func (e *originated) EventSubject() string {
	return e.self.PartType + "/" + e.self.PartID
}

func (e *originated) TenantID() string {
	return e.self.Call().TenantID
}

func (e *originated) MakeData() []byte {
	call := e.self.Call()
	e.data.Event = ievent.CloudEventOriginated
	e.data.EventEndPointID = e.self.PartID
	e.data.EventEndPointType = e.self.PartType
	e.data.EventUserID = e.self.UserID()
	e.data.EndPointID = e.self.PartID
	e.data.EndPointName = e.self.PartName
	e.data.EndPointType = e.self.PartType
	e.data.UCID = call.UCID
	e.data.CallSeq = json.Number(strconv.Itoa(call.Seq))
	e.data.TenantID = call.TenantID
	e.data.UserID = e.self.UserID()
	e.data.CallID = call.ID
	e.data.ConnectionID = e.self.ID
	e.data.ConnectionOldState = e.self.BeforeState()
	e.data.ConnectionNewState = e.self.State
	e.data.MediaType = call.MediaType
	e.data.Category = call.Category
	e.data.CallType = call.CallType
	e.data.CallSubType = call.CallSubType
	e.data.CallerID = call.CallerID
	e.data.CallerType = call.CallerType
	e.data.CalleeID = call.CalleeID
	e.data.CalleeType = call.CalleeType
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()

	byteData, err := json.Marshal(e.data)
	if err != nil {
		return nil
	}

	ilog.Debug("%v", string(byteData))
	return byteData
}
