package main

import (
	"github.com/RicardoCamilo/fileparser/db"
	"github.com/RicardoCamilo/fileparser/decoder"
	"github.com/RicardoCamilo/fileparser/model"
	"github.com/RicardoCamilo/fileparser/service"
	"github.com/joho/godotenv"
	"github.com/jszwec/csvutil"
	"log"
	"os"
	"strings"
)

const (
	header1 = "Call Date EST,Disposition,Phone Number,First Name,Last Name,Zipcode"
	header2 = "id,created_at,status,phone1,email,first,last,address1,address2,city,state,zip,called_count"
	header3 = "id,created_at,firstname,lastname,zip"
)

func main() {
	err := godotenv.Load(".env")
	user, password, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")

	database, err := db.Initialize(user, password, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
		return
	}
	defer database.Conn.Close()

	csvFile, err := os.ReadFile(os.Args[1:][0])
	if err != nil {
		log.Fatal(err)
		return
	}
	dec, err := decoder.GetDecoder(csvFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	switch getHeaderType(dec) {
	case header1:
		dec = decoder.SetDecoderTimeFormat(dec, "02/01/06 15:04:05")
		service.SavePhoneCall[model.PhoneCallType1](dec, database)
	case header2:
		dec = decoder.SetDecoderTimeFormat(dec, "2006-01-02T15:04:05-07:00")
		service.SavePhoneCall[model.PhoneCallType2](dec, database)
	case header3:
		dec = decoder.SetDecoderTimeFormat(dec, "2006-01-02T15:04:05-05:00")
		service.SavePhoneCall[model.PhoneCallType3](dec, database)
	default:
		log.Println("This CSV file is not supported")
	}
}

func getHeaderType(dec *csvutil.Decoder) string {
	return strings.Join(dec.Header(), ",")
}
