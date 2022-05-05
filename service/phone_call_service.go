package service

import (
	"database/sql"
	"github.com/RicardoCamilo/fileparser/db"
	"github.com/RicardoCamilo/fileparser/helper"
	"github.com/RicardoCamilo/fileparser/model"
	"github.com/jszwec/csvutil"
	"log"
)

func SavePhoneCall[T model.PhoneCallRecord](dec *csvutil.Decoder, database db.Database) (*csvutil.Decoder, bool) {
	entities, err := helper.GetEntitiesFromRecords(dec, make([]T, 0))
	if err != nil {
		log.Fatal(err)
		return nil, true
	}
	saveOrUpdate(entities, database)
	return dec, false
}

func saveOrUpdate(phoneCalls []model.PhoneCallEntity, database db.Database) {
	var err error
	for _, phoneCall := range phoneCalls {

		if phoneCall.Id != 0 {
			_, err = database.GetPhoneCallById(phoneCall.Id)
		} else {
			err = sql.ErrNoRows
		}

		switch err {
		case nil:
			_, err := database.UpdatePhoneCall(phoneCall)
			if err != nil {
				log.Println(err, phoneCall)
			}
		case sql.ErrNoRows:
			if phoneCall.Id != 0 {
				err = database.AddPhoneCallWithId(&phoneCall)
			} else {
				err = database.AddPhoneCall(&phoneCall)
			}
			if err != nil {
				log.Println(err)
			}
		default:
			log.Println(err)
		}
	}
}
