package model

import (
	"database/sql"
	"time"
)

type PhoneCallType1 struct {
	CallDateTime time.Time `csv:"Call Date EST"`
	Disposition  string    `csv:"Disposition"`
	PhoneNumber  string    `csv:"Phone Number"`
	FirstName    string    `csv:"First Name"`
	LastName     string    `csv:"Last Name"`
	ZipCode      string    `csv:"Zipcode"`
}

func (this PhoneCallType1) ConvertToEntity() PhoneCallEntity {
	return PhoneCallEntity{
		CallDateTime: this.CallDateTime,
		Disposition:  sql.NullString{String: this.Disposition, Valid: this.Disposition != ""},
		PhoneNumber:  sql.NullString{String: this.PhoneNumber, Valid: this.PhoneNumber != ""},
		FirstName:    sql.NullString{String: this.FirstName, Valid: true},
		LastName:     sql.NullString{String: this.LastName, Valid: true},
	}
}
