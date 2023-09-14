package db

import (
	"database/sql"
	"service-api/src/models/ent"
)

type Connect interface {
	connect(client *ent.Client) *sql.DB
}
