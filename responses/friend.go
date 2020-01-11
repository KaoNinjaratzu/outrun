package responses

import (
    "github.com/KaoNinjaratzu/outrun/netobj"
    "github.com/KaoNinjaratzu/outrun/obj"
    "github.com/KaoNinjaratzu/outrun/responses/responseobjs"
)

type FacebookIncentiveResponse struct {
    BaseResponse
    PlayerState netobj.PlayerState `json:"playerState"`
    Presents    []obj.Present      `json:"incentive"`
}

func FacebookIncentive(base responseobjs.BaseInfo, playerState netobj.PlayerState, presents []obj.Present) FacebookIncentiveResponse {
    baseResponse := NewBaseResponse(base)
    return FacebookIncentiveResponse{
        baseResponse,
        playerState,
        presents,
    }
}

func DefaultFacebookIncentive(base responseobjs.BaseInfo, player netobj.Player) FacebookIncentiveResponse {
    playerState := player.PlayerState
    presents := []obj.Present{} // Naughty this year
    return FacebookIncentive(
        base,
        playerState,
        presents,
    )
}
