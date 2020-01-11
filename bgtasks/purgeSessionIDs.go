package bgtasks

import (
	"time"

	"github.com/KaoNinjaratzu/outrun/db"
)

func PurgeSessionIDs() {
	for true {
		time.Sleep(10 * time.Minute)
		db.PurgeAllExpiredSessionIDs()
	}
}
