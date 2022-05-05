package helper

import (
	"github.com/RicardoCamilo/fileparser/model"
	"github.com/jszwec/csvutil"
	"io"
)

func getRecords[T model.PhoneCallRecord](dec *csvutil.Decoder, phoneCalls []T) ([]T, error) {
	for {
		var phoneCall T
		if err := dec.Decode(&phoneCall); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		phoneCalls = append(phoneCalls, phoneCall)
	}
	return phoneCalls, nil
}

func mapRecordsToEntity[T model.PhoneCallRecord](phoneCalls []T) []model.PhoneCallEntity {
	var entities []model.PhoneCallEntity
	for _, phoneCall := range phoneCalls {
		entities = append(entities, phoneCall.ConvertToEntity())
	}
	return entities
}

func GetEntitiesFromRecords[T model.PhoneCallRecord](dec *csvutil.Decoder, phoneCalls []T) ([]model.PhoneCallEntity, error) {
	records, err := getRecords(dec, phoneCalls)
	if err != nil {
		return nil, err
	}
	entities := mapRecordsToEntity(records)
	return entities, nil
}
