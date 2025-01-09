package event

import (
	"encoding/json"
	"strconv"
	"time"

	"gitlab.com/ipron-core/call/internal/core/domain/entity"
	"gitlab.com/ipron-core/call/internal/core/domain/vo"
	"gitlab.com/ipron-core/call/internal/pkg/types"
	"gitlab.com/ipron-core/call/internal/pkg/utils"
	"gitlab.com/ipron-ne/ArchDB/v2/schema/code"
	"gitlab.com/ipron-ne/iCore/ilog"
	"gitlab.com/ipron-ne/ievent"

	callevt "gitlab.com/ipron-ne/ievent/call"
)

type disconnected struct {
	self   entity.Leg
	happen entity.Leg
	data   callevt.Disconnected
}

func NewDisconnected(self, happen entity.Leg) Event {
	return &disconnected{self: self, happen: happen}
}

func (e *disconnected) Topic() string {
	return ievent.CloudTopicCall
}

func (e *disconnected) Type() string {
	return ievent.CloudTypeCallDisconnected
}

func (e *disconnected) Subject() string {
	return subjectKey + e.self.Call().ID
}

func (e *disconnected) EventSubject() string {
	return e.self.PartType + "/" + e.self.PartID
}

func (e *disconnected) TenantID() string {
	return e.self.Call().TenantID
}

func (e *disconnected) clear() {
	var (
		call  = e.self.Call()
		party vo.Party
	)
	e.data.Event = ievent.CloudEventDisconnected
	e.data.EventEndPointID = e.happen.PartID
	e.data.EventEndPointType = e.happen.PartType
	e.data.EventUserID = e.happen.UserID()
	e.data.EndPointID = e.self.PartID
	e.data.EndPointType = e.self.PartType
	e.data.EndPointName = e.self.PartName
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
	for _, leg := range call.Legs {
		party.AddPartyInfos(utils.ToLegToPartyInfo(*leg))
	}
	e.data.OldPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.OldPartyInfoSet = party.GetSlice()
	party.RemovePartyInfos(e.happen.PartID)
	e.data.NewPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.NewPartyInfoSet = party.GetSlice()
	e.data.LastRingTime = call.LastRingTime.UTC()
	e.data.LastConnTime = call.LastConnTime.UTC()
	e.data.LastHoldTime = call.LastHoldTime.UTC()
	e.data.LastConfTime = call.LastConfTime.UTC()
	e.data.FirstFlowInTime = call.FirstFlowInTime.UTC()
	e.data.FirstFlowID = call.FirstFlowID
	e.data.CurrFlowID = call.CurrentFlowID
	e.data.FirstAcdID = call.FirstAcdID
	e.data.CurrAcdID = call.CurrentAcdID
	e.data.FirstAcdInTime = call.FirstAcdInTime.UTC()
	e.data.CurrAcdInTime = call.CurrentAcdInTime.UTC()
	e.data.FirstSkillID = call.FirstSkillID
	e.data.CurrSkillID = call.CurrentSkillID

	switch {
	case e.happen.PartID == e.self.PartID:
		e.data.EndState = true
	default:
		e.data.EndState = false
	}

	switch {
	case call.IsTerminate:
		e.data.EndPart = string(ievent.CallEndPartSystem)
	case e.happen.Call().CallType == string(code.CallTypeFeature):
		e.data.EndPart = string(ievent.CallEndPartService)
	case call.End.ReqPartType == types.PartTrunk:
		e.data.EndPart = string(ievent.CallEndPartCustom)
	case call.End.ReqPartType == types.PartUser,
		call.End.ReqPartType == types.PartPhone:
		e.data.EndPart = string(ievent.CallEndPartUser)
	default:
		e.data.EndPart = string(ievent.CallEndPartService)
	}

	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *disconnected) singleTransfer() {
	var (
		call  = e.self.Call()
		event = &call.Event
		party vo.Party
	)
	e.data.Event = ievent.CloudEventDisconnected
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
	party.RemovePartyInfos(event.TransToParty.EndPointID)
	e.data.OldPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.OldPartyInfoSet = party.GetSlice()
	party.RemovePartyInfos(e.happen.PartID)
	party.AddPartyInfos(event.TransToParty)
	e.data.NewPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.NewPartyInfoSet = party.GetSlice()
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	e.data.TransToType = event.TransToParty.EndPointType
	e.data.TransToID = event.TransToParty.EndPointID
	e.data.LastRingTime = call.LastRingTime.UTC()
	e.data.LastConnTime = call.LastConnTime.UTC()
	e.data.LastHoldTime = call.LastHoldTime.UTC()
	e.data.LastConfTime = call.LastConfTime.UTC()
	e.data.FirstFlowInTime = call.FirstFlowInTime.UTC()
	e.data.FirstFlowID = call.FirstFlowID
	e.data.CurrFlowID = call.CurrentFlowID
	e.data.FirstAcdID = call.FirstAcdID
	e.data.CurrAcdID = call.CurrentAcdID
	e.data.FirstAcdInTime = call.FirstAcdInTime.UTC()
	e.data.CurrAcdInTime = call.CurrentAcdInTime.UTC()
	e.data.FirstSkillID = call.FirstSkillID
	e.data.CurrSkillID = call.CurrentSkillID
	e.data.EndState = true
	e.data.EndPart = string(ievent.CallEndPartConsult)
	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *disconnected) muteTransfer() {
	var (
		call  = e.self.Call()
		event = call.Event
		party vo.Party
	)
	e.data.Event = ievent.CloudEventDisconnected
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
	e.data.OldPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.OldPartyInfoSet = party.GetSlice()
	party.ClearPartyInfos()
	e.data.NewPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.NewPartyInfoSet = party.GetSlice()
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	e.data.TransToType = event.TransToParty.EndPointType
	e.data.TransToID = event.TransToParty.EndPointID
	e.data.LastRingTime = call.LastRingTime.UTC()
	e.data.LastConnTime = call.LastConnTime.UTC()
	e.data.LastHoldTime = call.LastHoldTime.UTC()
	e.data.LastConfTime = call.LastConfTime.UTC()
	e.data.FirstFlowInTime = call.FirstFlowInTime.UTC()
	e.data.FirstFlowID = call.FirstFlowID
	e.data.CurrFlowID = call.CurrentFlowID
	e.data.FirstAcdID = call.FirstAcdID
	e.data.CurrAcdID = call.CurrentAcdID
	e.data.FirstAcdInTime = call.FirstAcdInTime.UTC()
	e.data.CurrAcdInTime = call.CurrentAcdInTime.UTC()
	e.data.FirstSkillID = call.FirstSkillID
	e.data.CurrSkillID = call.CurrentSkillID
	e.data.EndState = true
	e.data.EndPart = string(ievent.CallEndPartConsult)
	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *disconnected) muteConference() {
	var (
		call  = e.self.Call()
		event = &call.Event
		party vo.Party
	)
	e.data.Event = ievent.CloudEventDisconnected
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
	e.data.OldPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.OldPartyInfoSet = party.GetSlice()
	party.ClearPartyInfos()
	e.data.NewPartyCount = json.Number(strconv.Itoa(party.PartyInfosCount()))
	e.data.NewPartyInfoSet = party.GetSlice()
	e.data.ANI = call.ANI
	e.data.DNIS = call.DNIS
	e.data.OriginalNum = e.self.OriginalNum
	e.data.RealNum = e.self.RealNum
	e.data.AccessCode = e.self.AccessCode
	e.data.Pattern = e.self.Pattern
	e.data.UUI = call.UUI
	e.data.UEI = call.UEI
	e.data.Reason = call.Reason
	e.data.TransToType = event.TransToParty.EndPointType
	e.data.TransToID = event.TransToParty.EndPointID
	e.data.LastRingTime = call.LastRingTime.UTC()
	e.data.LastConnTime = call.LastConnTime.UTC()
	e.data.LastHoldTime = call.LastHoldTime.UTC()
	e.data.LastConfTime = call.LastConfTime.UTC()
	e.data.FirstFlowInTime = call.FirstFlowInTime.UTC()
	e.data.FirstFlowID = call.FirstFlowID
	e.data.CurrFlowID = call.CurrentFlowID
	e.data.FirstAcdID = call.FirstAcdID
	e.data.CurrAcdID = call.CurrentAcdID
	e.data.FirstAcdInTime = call.FirstAcdInTime.UTC()
	e.data.CurrAcdInTime = call.CurrentAcdInTime.UTC()
	e.data.FirstSkillID = call.FirstSkillID
	e.data.CurrSkillID = call.CurrentSkillID
	e.data.EndState = false
	e.data.EndPart = string(ievent.CallEndPartConsult)
	e.data.UserInCount = json.Number(strconv.Itoa(call.UserInCount))
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()
}

func (e *disconnected) MakeData() []byte {
	call := e.self.Call()
	switch ievent.ReasonCode(call.Reason) {
	case ievent.ReasonCodeSingleTransfer:
		e.singleTransfer()
	case ievent.ReasonCodeMuteTransfer:
		e.muteTransfer()
	case ievent.ReasonCodeMuteConference:
		e.muteConference()
	default:
		e.clear()
	}

	byteData, err := json.Marshal(e.data)
	if err != nil {
		return nil
	}

	ilog.Debug("%v", string(byteData))
	return byteData
}
