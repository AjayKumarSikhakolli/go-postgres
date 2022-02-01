package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
  _ "github.com/lib/pq"
)

const (
  host     = "172.19.104.111"
  port     = 5432
  user     = "postgres"
  password = "postgres"
  dbname   = "postgres"
)

func createDBConnection() *sql.DB {

        // Open the connection
        psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
                "password=%s dbname=%s sslmode=disable",
                host, port, user, password, dbname)

        log.Println("Psql info:", psqlInfo)

        db, err := sql.Open("postgres", psqlInfo)

        if err != nil {
                panic(err)
        }

        // check the connection
        err = db.Ping()

        if err != nil {
                panic(err)
        }

        fmt.Println("Successfully connected!")

        // return the connection
        return db
}

func CreateTable() bool{

	db := createDBConnection()

        // close the db connection
        defer db.Close()

        //Create Sample table
	sqlStatmentCreateTable := `CREATE TABLE IF NOT EXISTS sample(TimeStamp TIMESTAMP,Name TEXT,Address TEXT,MobileNo TEXT,ID INT, PRIMARY KEY (ID))`

        res, err := db.Exec(sqlStatmentCreateTable)

        if err != nil {
                log.Printf("Error %s when creating sample table", err)
                panic(err)
                return false
        }

        rows, err := res.RowsAffected()

        if err != nil {
                log.Printf("Error %s when getting rows affected", err)
                panic(err)
                return false
        }

        log.Printf("Rows affected when creating table: %d", rows)

        return true

}

func AddSampleData(Name string,Address string,MobileNo string,ID int) int{

	 db := createDBConnection()

        // close the db connection
        defer db.Close()

	sqlStatement := `INSERT INTO sample (TimeStamp,Name,Address,MobileNo,ID) VALUES ($1, $2, $3, $4,$5) RETURNING ID`

	var id int

	TimeStamp := time.Now().UTC().Format(time.RFC3339)

	err := db.QueryRow(sqlStatement,TimeStamp,Name,Address,MobileNo,ID).Scan(&id)

	if err != nil {
		panic(err)
	}

	log.Println("Record Inserted:", id)

	return id
}

func main() {

	var(
		Name,Address,MobileNo string
		ID int
	)

	Name = "John"
	Address = "NewYork"
	MobileNo = "9876543210"
	ID = 1

	Created := CreateTable()

	if Created == true {

		log.Println("Record Inserted:",AddSampleData(Name,Address,MobileNo,ID))
	}

}
