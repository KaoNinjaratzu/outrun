package logic

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/KaoNinjaratzu/outrun/netobj"
)

// GenerateLoginPasskey is used by LoginDelta to verify the login passkey sent by the game.
func GenerateLoginPasskey(player netobj.Player) string {
	data := []byte(player.Key + ":dho5v5yy7n2uswa5iblb:" + player.ID + ":" + player.Password)
	sum := md5.Sum(data)
	return hex.EncodeToString(sum[:])
}
