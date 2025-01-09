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

type partyChanged struct {
	self   entity.Leg
	happen entity.Leg
	data   callevt.PartyChanged
}

func NewPartyChanged(self, happen entity.Leg) Event {
	return &partyChanged{self: self, happen: happen}
}

func (e *partyChanged) Topic() string {
	return ievent.CloudTopicCall
}

func (e *partyChanged) Type() string {
	return ievent.CloudTypeCallPartyChanged
}

func (e *partyChanged) Subject() string {
	return subjectKey + e.self.Call().ID
}

func (e *partyChanged) EventSubject() string {
	return e.self.PartType + "/" + e.self.PartID
}

func (e *partyChanged) TenantID() string {
	return e.self.Call().TenantID
}

func (e *partyChanged) singleTransfer() {
	var (
		call  = e.self.Call()
		event = &call.Event
		party vo.Party
	)
	e.data.Event = ievent.CloudEventPartyChanged
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
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.CallerID, e.data.CallerType = call.CallerID, call.CallerType
	e.data.CalleeID, e.data.CalleeType = call.CalleeID, call.CalleeType
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	e.data.ChangeFromType = event.ChangeFromParty.EndPointType
	e.data.ChangeFromID = event.ChangeFromParty.EndPointID
	e.data.ChangeToType = event.ChangeToParty.EndPointType
	e.data.ChangeToID = event.ChangeToParty.EndPointID
	for _, leg := range call.Legs {
		party.AddPartyInfos(utils.ToLegToPartyInfo(*leg))
	}
	party.RemovePartyInfos(event.TransToParty.EndPointID)
	e.data.OldPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.OldPartyInfoSet = party.GetSlice()
	party.RemovePartyInfos(event.ChangeFromParty.EndPointID)
	party.AddPartyInfos(event.ChangeToParty)
	e.data.NewPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.NewPartyInfoSet = party.GetSlice()
	e.data.LastConnTime = call.LastConnTime.UTC()
	e.data.NewCallSeq = json.Number(call.Event.NewCallSeq)
	e.data.NewCallID = call.Event.NewCallID
	e.data.NewConnectionID = call.Event.NewConnID
	e.data.NewCategory = call.Event.NewCategory
	e.data.NewCallType = call.Event.NewCallType
	e.data.NewCallSubType = call.Event.NewCallSubType
	e.data.NewANI = call.Event.NewANI
	e.data.NewDNIS = call.Event.NewDNIS
	e.data.NewCreateTime = call.Event.NewCreateTime.UTC()
	e.data.EndState = false
	e.data.EndPart = ""
	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *partyChanged) muteTransfer() {
	var (
		call  = e.self.Call()
		event = &call.Event
		party vo.Party
	)
	e.data.Event = ievent.CloudEventPartyChanged
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
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.CallerID = call.CallerID
	e.data.CallerType = call.CallerType
	e.data.CalleeID = call.CalleeID
	e.data.CalleeType = call.CalleeType
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	e.data.ChangeFromID = event.ChangeFromParty.EndPointID
	e.data.ChangeFromType = event.ChangeFromParty.EndPointType
	e.data.ChangeToID = event.ChangeToParty.EndPointID
	e.data.ChangeToType = event.ChangeToParty.EndPointType
	for _, leg := range call.Legs {
		party.AddPartyInfos(utils.ToLegToPartyInfo(*leg))
	}
	e.data.OldPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.OldPartyInfoSet = party.GetSlice()
	party.RemovePartyInfos(event.ChangeFromParty.EndPointID)
	party.AddPartyInfos(event.ChangeToParty)
	e.data.NewPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.NewPartyInfoSet = party.GetSlice()
	e.data.NewCallSeq = json.Number(call.Event.NewCallSeq)
	e.data.NewCallID = call.Event.NewCallID
	e.data.NewConnectionID = call.Event.NewConnID
	e.data.NewCategory = call.Event.NewCategory
	e.data.NewCallType = call.Event.NewCallType
	e.data.NewCallSubType = call.Event.NewCallSubType
	e.data.NewANI = call.Event.NewANI
	e.data.NewDNIS = call.Event.NewDNIS
	e.data.LastConnTime = call.LastConnTime.UTC()
	e.data.NewCreateTime = call.Event.NewCreateTime.UTC()
	e.data.EndState = false
	e.data.EndPart = ""
	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *partyChanged) muteConference() {
	var (
		call  = e.self.Call()
		party vo.Party
		event = &call.Event
	)
	e.data.Event = ievent.CloudEventPartyChanged
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
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.CallerID = call.CallerID
	e.data.CallerType = call.CallerType
	e.data.CalleeID = call.CalleeID
	e.data.CalleeType = call.CalleeType
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	for _, leg := range call.Legs {
		party.AddPartyInfos(utils.ToLegToPartyInfo(*leg))
	}
	e.data.OldPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.OldPartyInfoSet = party.GetSlice()
	party.AddPartyInfos(event.ChangeToParty)
	e.data.NewPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.NewPartyInfoSet = party.GetSlice()
	e.data.LastConnTime = call.LastConnTime.UTC()
	e.data.NewCallSeq = json.Number(call.Event.NewCallSeq)
	e.data.NewCallID = call.Event.NewCallID
	e.data.NewConnectionID = call.Event.NewConnID
	e.data.NewCategory = call.Event.NewCategory
	e.data.NewCallType = call.Event.NewCallType
	e.data.NewCallSubType = call.Event.NewCallSubType
	e.data.NewANI = call.Event.NewANI
	e.data.NewDNIS = call.Event.NewDNIS
	e.data.NewCreateTime = call.Event.NewCreateTime.UTC()
	e.data.EndState = false
	e.data.EndPart = string(ievent.CallEndPartConsult)
	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *partyChanged) MakeData() []byte {
	switch ievent.ReasonCode(e.self.Call().Reason) {
	case ievent.ReasonCodeSingleTransfer:
		e.singleTransfer()
	case ievent.ReasonCodeMuteTransfer:
		e.muteTransfer()
	case ievent.ReasonCodeMuteConference:
		e.muteConference()
	}

	byteData, err := json.Marshal(e.data)
	if err != nil {
		return nil
	}

	ilog.Debug("%v", string(byteData))
	return byteData
}
