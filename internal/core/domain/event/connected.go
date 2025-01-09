package event

import (
	"encoding/json"
	"strconv"
	"time"

	"gitlab.com/ipron-core/call/internal/core/domain/entity"
	"gitlab.com/ipron-core/call/internal/core/domain/vo"
	"gitlab.com/ipron-core/call/internal/pkg/utils"
	"gitlab.com/ipron-ne/iCore/ilog"
	"gitlab.com/ipron-ne/ievent"

	callevt "gitlab.com/ipron-ne/ievent/call"
)

type connected struct {
	self   entity.Leg
	happen entity.Leg
	data   callevt.Connected
}

func NewConnected(self, happen entity.Leg) Event {
	return &connected{self: self, happen: happen}
}

func (e *connected) Topic() string {
	return ievent.CloudTopicCall
}

func (e *connected) Type() string {
	return ievent.CloudTypeCallConnected
}

func (e *connected) Subject() string {
	return subjectKey + e.self.Call().ID
}

func (e *connected) EventSubject() string {
	return e.self.PartType + "/" + e.self.PartID
}

func (e *connected) TenantID() string {
	return e.self.Call().TenantID
}

func (e *connected) normal() {
	var (
		call  = e.self.Call()
		party vo.Party
	)
	e.data.Event = ievent.CloudEventConnected
	e.data.EventEndPointID = e.happen.PartID
	e.data.EventEndPointType = e.happen.PartType
	e.data.EventUserID = e.happen.UserID()
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
	for _, leg := range call.Legs {
		party.AddPartyInfos(utils.ToLegToPartyInfo(*leg))
	}
	e.data.PartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.PartyInfoSet = party.GetSlice()
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	e.data.LastRingTime = call.LastRingTime.UTC()
	e.data.FirstFlowID = call.FirstFlowID
	e.data.CurrFlowId = call.CurrentFlowID
	e.data.FirstAcdID = call.FirstAcdID
	e.data.CurrAcdID = call.CurrentAcdID
	e.data.FirstAcdInTime = call.FirstAcdInTime.UTC()
	e.data.CurrAcdInTime = call.CurrentAcdInTime.UTC()
	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.FirstSkillID = call.FirstSkillID
	e.data.CurrSkillID = call.CurrentSkillID
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *connected) muteConference() {
	var (
		call  = e.self.Call()
		party vo.Party
	)
	e.data.Event = ievent.CloudEventConnected
	e.data.EventEndPointID = e.happen.PartID
	e.data.EventEndPointType = e.happen.PartType
	e.data.EventUserID = e.happen.UserID()
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
	for _, leg := range call.Legs {
		party.AddPartyInfos(utils.ToLegToPartyInfo(*leg))
	}
	party.AddPartyInfos(utils.ToLegToPartyInfo(e.happen))
	e.data.PartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.PartyInfoSet = party.GetSlice()
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	e.data.LastRingTime = call.LastRingTime.UTC()
	e.data.FirstFlowID = call.FirstFlowID
	e.data.CurrFlowId = call.CurrentFlowID
	e.data.FirstAcdID = call.FirstAcdID
	e.data.CurrAcdID = call.CurrentAcdID
	e.data.FirstAcdInTime = call.FirstAcdInTime.UTC()
	e.data.CurrAcdInTime = call.CurrentAcdInTime.UTC()
	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.FirstSkillID = call.FirstSkillID
	e.data.CurrSkillID = call.CurrentSkillID
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *connected) MakeData() []byte {
	switch ievent.ReasonCode(e.self.Call().Reason) {
	case ievent.ReasonCodeMuteConference:
		e.muteConference()
	default:
		e.normal()
	}

	byteData, err := json.Marshal(e.data)
	if err != nil {
		return nil
	}

	ilog.Debug("%v", string(byteData))
	return byteData
}
