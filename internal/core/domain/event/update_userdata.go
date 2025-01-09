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

type updateUserData struct {
	self entity.Leg
	data callevt.UpdateUserData
}

func NewUpdateUserData(self entity.Leg) Event {
	return &updateUserData{self: self}
}

func (e *updateUserData) Topic() string {
	return ievent.CloudTopicCall
}

func (e *updateUserData) Type() string {
	return ievent.CloudTypeCallUpdateUserData
}

func (e *updateUserData) Subject() string {
	return subjectKey + e.self.Call().ID
}

func (e *updateUserData) EventSubject() string {
	return e.self.PartType + "/" + e.self.PartID
}

func (e *updateUserData) TenantID() string {
	return e.self.Call().TenantID
}

func (e *updateUserData) MakeData() []byte {
	var (
		call  = e.self.Call()
		party vo.Party
	)
	e.data.Event = ievent.CloudEventUpdateUserData
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
	e.data.LastHoldTime = call.LastHoldTime.UTC()
	e.data.FirstAcdID = call.FirstAcdID
	e.data.FirstSkillID = call.FirstSkillID
	e.data.CurrSkillID = call.CurrentSkillID
	e.data.CreateTime = call.CreateTime.UTC()
	e.data.NowTime = time.Now().UTC()

	byteData, err := json.Marshal(e.data)
	if err != nil {
		return nil
	}

	ilog.Debug("%v", string(byteData))
	return byteData
}
