package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
)

func generateDiff() {
	referenceDb := pg.Connect(&pg.Options{
		Network:         "tcp",
		Addr:            fmt.Sprintf("%s:%d", referencePostgresHost, referencePostgresPort),
		User:            referencePostgresUser,
		Password:        referencePostgresPassword,
		Database:        referencePostgresDatabase,
		ApplicationName: "migrate",
	})
	desiredDb := pg.Connect(&pg.Options{
		Network:         "tcp",
		Addr:            fmt.Sprintf("%s:%d", desiredPostgresHost, desiredPostgresPort),
		User:            desiredPostgresUser,
		Password:        desiredPostgresPassword,
		Database:        desiredPostgresDatabase,
		ApplicationName: "migrate",
	})

	mustPingDatabase(referenceDb)
	mustPingDatabase(desiredDb)


}

func mustPingDatabase(db *pg.DB) {
	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}
}
