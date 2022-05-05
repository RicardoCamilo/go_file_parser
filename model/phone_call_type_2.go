package model

import (
	"database/sql"
	"time"
)

type PhoneCallType2 struct {
	Id           int       `csv:"id"`
	CallDateTime time.Time `csv:"created_at"`
	Disposition  string    `csv:"status"`
	PhoneNumber  string    `csv:"phone1"`
	Email        string    `csv:"email"`
	FirstName    string    `csv:"first"`
	LastName     string    `csv:"last"`
	Address1     string    `csv:"address1"`
	Address2     string    `csv:"address2"`
	City         string    `csv:"city"`
	State        string    `csv:"state"`
	ZipCode      string    `csv:"zip"`
	CalledCount  int32     `csv:"called_count"`
}

func (this PhoneCallType2) ConvertToEntity() PhoneCallEntity {
	return PhoneCallEntity{
		Id:           this.Id,
		CallDateTime: this.CallDateTime,
		Disposition:  sql.NullString{String: this.Disposition, Valid: this.Disposition != ""},
		PhoneNumber:  sql.NullString{String: this.PhoneNumber, Valid: this.PhoneNumber != ""},
		Email:        sql.NullString{String: this.Email, Valid: true},
		FirstName:    sql.NullString{String: this.FirstName, Valid: true},
		LastName:     sql.NullString{String: this.LastName, Valid: true},
		Address1:     sql.NullString{String: this.Address1, Valid: true},
		Address2:     sql.NullString{String: this.Address2, Valid: true},
		City:         sql.NullString{String: this.City, Valid: true},
		State:        sql.NullString{String: this.State, Valid: true},
		ZipCode:      sql.NullString{String: this.ZipCode, Valid: true},
		CalledCount:  sql.NullInt32{Int32: this.CalledCount, Valid: true},
	}
}
