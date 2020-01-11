package muxhandlers

import (
	"github.com/KaoNinjaratzu/outrun/config/eventconf"
	"github.com/KaoNinjaratzu/outrun/emess"
	"github.com/KaoNinjaratzu/outrun/helper"
	"github.com/KaoNinjaratzu/outrun/logic/conversion"
	"github.com/KaoNinjaratzu/outrun/obj"
	"github.com/KaoNinjaratzu/outrun/responses"
	"github.com/KaoNinjaratzu/outrun/status"
)

func GetEventList(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	// construct event list
	eventList := []obj.Event{}
	if eventconf.CFile.AllowEvents {
		if eventconf.CFile.EnforceGlobal || len(player.PersonalEvents) == 0 {
			for _, confEvent := range eventconf.CFile.CurrentEvents {
				newEvent := conversion.ConfiguredEventToEvent(confEvent)
				eventList = append(eventList, newEvent)
			}
		} else {
			for _, ce := range player.PersonalEvents {
				e := conversion.ConfiguredEventToEvent(ce)
				eventList = append(eventList, e)
			}
		}
	}
	helper.DebugOut("Personal event list: %v", player.PersonalEvents)
	helper.DebugOut("Global event list: %v", eventconf.CFile.CurrentEvents)
	helper.DebugOut("Event list: %v", eventList)
	response := responses.EventList(baseInfo, eventList)
	err = helper.SendResponse(response)
	if err != nil {
		helper.InternalErr("Error sending response", err)
	}
}
