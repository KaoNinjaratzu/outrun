package logic

import (
	"log"

	"github.com/KaoNinjaratzu/outrun/consts"
	"github.com/KaoNinjaratzu/outrun/db"
	"github.com/KaoNinjaratzu/outrun/db/dbaccess"
	"github.com/KaoNinjaratzu/outrun/netobj"
)

func FindPlayersByPassword(password string, silent bool) ([]netobj.Player, error) {
	playerIDs := []string{}
	players := []netobj.Player{}
	dbaccess.ForEachKey(consts.DBBucketPlayers, func(k, v []byte) error {
		playerIDs = append(playerIDs, string(k))
		return nil
	})
	for _, pid := range playerIDs {
		player, err := db.GetPlayer(pid)
		if err != nil {
			if silent {
				log.Printf("[WARN] (logic.FindPlayersByPassword) Unable to get player '%s': %s", pid, err.Error())
			} else {
				return []netobj.Player{}, err
			}
		}
		if player.Password == password {
			players = append(players, player)
		}
	}
	return players, nil
}

func FindPlayersByMigrationPassword(password string, silent bool) ([]netobj.Player, error) {
	playerIDs := []string{}
	players := []netobj.Player{}
	dbaccess.ForEachKey(consts.DBBucketPlayers, func(k, v []byte) error {
		playerIDs = append(playerIDs, string(k))
		return nil
	})
	for _, pid := range playerIDs {
		player, err := db.GetPlayer(pid)
		if err != nil {
			if silent {
				log.Printf("[WARN] (logic.FindPlayersByMigrationPassword) Unable to get player '%s': %s", pid, err.Error())
			} else {
				return []netobj.Player{}, err
			}
		}
		if player.MigrationPassword == password {
			players = append(players, player)
		}
	}
	return players, nil
}
