package muxhandlers

import (
    "github.com/KaoNinjaratzu/outrun/emess"
    "github.com/KaoNinjaratzu/outrun/helper"
    "github.com/KaoNinjaratzu/outrun/responses"
    "github.com/KaoNinjaratzu/outrun/status"
)

func SendApollo(helper *helper.Helper) {
    baseInfo := helper.BaseInfo(emess.OK, status.OK)
    response := responses.NewBaseResponse(baseInfo)
    err := helper.SendResponse(response)
    if err != nil {
        helper.InternalErr("Error sending response", err)
    }
}
