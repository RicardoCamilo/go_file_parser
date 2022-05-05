package db

import (
	"database/sql"
	"fmt"
	"github.com/RicardoCamilo/fileparser/model"
	_ "github.com/lib/pq"
	"log"
)

const (
	HOST = "localhost"
	PORT = 5432
)

var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	log.Println(dsn)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil
}

func (db *Database) AddPhoneCall(phoneCall *model.PhoneCallEntity) error {
	var id int
	query := `INSERT INTO phone_calls(
				  call_date_time, disposition, phone_number, 
				  first_name, last_name, address1, address2, 
				  city, state, zip, email, called_count
				) 
				VALUES 
				  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	err := db.Conn.QueryRow(
		query,
		phoneCall.CallDateTime,
		phoneCall.Disposition,
		phoneCall.PhoneNumber,
		phoneCall.FirstName,
		phoneCall.LastName,
		phoneCall.Address1,
		phoneCall.Address2,
		phoneCall.City,
		phoneCall.State,
		phoneCall.ZipCode,
		phoneCall.Email,
		phoneCall.CalledCount,
	).Scan(&id)

	if err != nil {
		return err
	}

	fmt.Println("New record ID is:", id)
	phoneCall.Id = id

	return nil
}

func (db Database) GetPhoneCallById(phoneCallId int) (model.PhoneCallEntity, error) {
	phoneCall := model.PhoneCallEntity{}
	query := `SELECT * FROM phone_calls WHERE id = $1;`
	row := db.Conn.QueryRow(query, phoneCallId)

	err := row.Scan(
		&phoneCall.Id,
		&phoneCall.CallDateTime,
		&phoneCall.Disposition,
		&phoneCall.PhoneNumber,
		&phoneCall.FirstName,
		&phoneCall.LastName,
		&phoneCall.Address1,
		&phoneCall.Address2,
		&phoneCall.City,
		&phoneCall.State,
		&phoneCall.ZipCode,
		&phoneCall.Email,
		&phoneCall.CalledCount,
	)

	switch err {
	case sql.ErrNoRows:
		return phoneCall, sql.ErrNoRows
	default:
		return phoneCall, err
	}
}

func (db Database) UpdatePhoneCall(phoneCallData model.PhoneCallEntity) (model.PhoneCallEntity, error) {
	phoneCall := model.PhoneCallEntity{}
	query := `UPDATE phone_calls  
				SET call_date_time=$1, disposition=$2, phone_number=$3, 
				  first_name=$4, last_name=$5, address1=$6, address2=$7, 
				  city=$8, state=$9, zip=$10, email=$11, called_count=$12
				WHERE id=$13 
				RETURNING id;`
	err := db.Conn.QueryRow(query,
		phoneCallData.CallDateTime,
		phoneCallData.Disposition,
		phoneCallData.PhoneNumber,
		phoneCallData.FirstName,
		phoneCallData.LastName,
		phoneCallData.Address1,
		phoneCallData.Address2,
		phoneCallData.City,
		phoneCallData.State,
		phoneCallData.ZipCode,
		phoneCallData.Email,
		phoneCallData.CalledCount,
		phoneCallData.Id,
	).Scan(
		&phoneCall.Id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return phoneCall, ErrNoMatch
		}
		return phoneCall, err
	}
	return phoneCall, nil
}

func (db *Database) AddPhoneCallWithId(phoneCall *model.PhoneCallEntity) error {
	var id int
	query := `INSERT INTO phone_calls(
				  id, call_date_time, disposition, phone_number, 
				  first_name, last_name, address1, address2, 
				  city, state, zip, email, called_count
				) 
				VALUES 
				  ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
	err := db.Conn.QueryRow(
		query,
		phoneCall.Id,
		phoneCall.CallDateTime,
		phoneCall.Disposition,
		phoneCall.PhoneNumber,
		phoneCall.FirstName,
		phoneCall.LastName,
		phoneCall.Address1,
		phoneCall.Address2,
		phoneCall.City,
		phoneCall.State,
		phoneCall.ZipCode,
		phoneCall.Email,
		phoneCall.CalledCount,
	).Scan(&id)

	if err != nil {
		return err
	}

	fmt.Println("New record ID is:", id)
	phoneCall.Id = id

	return nil
}
