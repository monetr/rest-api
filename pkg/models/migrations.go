package models

import "time"

type Migration struct {
	tableName string `pg:"migrations"`

	MigrationId int       `pg:"migration_id,notnull,pk"`
	MD5Hash     string    `pg:"md5_hash,notnull"`
	AppliedAt   time.Time `pg:"applied_at,notnull"`
}
