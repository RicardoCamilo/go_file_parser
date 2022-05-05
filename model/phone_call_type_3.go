package model

import (
	"database/sql"
	"time"
)

type PhoneCallType3 struct {
	Id           int       `csv:"id"`
	CallDateTime time.Time `csv:"created_at"`
	FirstName    string    `csv:"firstname"`
	LastName     string    `csv:"lastname"`
	ZipCode      string    `csv:"zip"`
}

func (this PhoneCallType3) ConvertToEntity() PhoneCallEntity {
	return PhoneCallEntity{
		Id:           this.Id,
		CallDateTime: this.CallDateTime,
		FirstName:    sql.NullString{String: this.FirstName, Valid: true},
		LastName:     sql.NullString{String: this.LastName, Valid: true},
		ZipCode:      sql.NullString{String: this.ZipCode, Valid: true},
	}
}
