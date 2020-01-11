package muxhandlers

import (
	"github.com/KaoNinjaratzu/outrun/emess"
	"github.com/KaoNinjaratzu/outrun/helper"
	"github.com/KaoNinjaratzu/outrun/responses"
	"github.com/KaoNinjaratzu/outrun/status"
)

func GetOptionUserResult(helper *helper.Helper) {
	player, err := helper.GetCallingPlayer()
	if err != nil {
		helper.InternalErr("Error getting calling player", err)
		return
	}
	baseInfo := helper.BaseInfo(emess.OK, status.OK)
	response := responses.OptionUserResult(baseInfo, player.OptionUserResult)
	helper.SendResponse(response)
}
