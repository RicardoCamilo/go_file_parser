package model

import (
	"database/sql"
	"time"
)

type PhoneCallEntity struct {
	Id           int
	CallDateTime time.Time
	Disposition  sql.NullString
	PhoneNumber  sql.NullString
	FirstName    sql.NullString
	LastName     sql.NullString
	Address1     sql.NullString
	Address2     sql.NullString
	City         sql.NullString
	State        sql.NullString
	ZipCode      sql.NullString
	Email        sql.NullString
	CalledCount  sql.NullInt32
}

type PhoneCallRecord interface {
	ConvertToEntity() PhoneCallEntity
}
