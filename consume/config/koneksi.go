package main

import (

	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	buatDB()
	buatTabel()
}

func koneksiMysql() (db *sql.DB){
	dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/")
    if err != nil {
        panic(err.Error())
    }
	return db
}

func buatDB(){

	db := koneksiMysql()

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS spenews")
	
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("berhasil terhubung dengan database")
	}
}

func koneksiDB() (db *sql.DB) {

    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "spenews"
    db, err := sql.Open(dbDriver, dbUser+":"+  dbPass + "@/" + dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}

func buatTabel(){

	db := koneksiDB()

	_, err := db.Exec("USE spenews")
	
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("berhasil memilih database")
	}

	stmt, err := db.Prepare(`CREATE TABLE bisnis (
		id int NOT NULL AUTO_INCREMENT, 
		sourceid varchar(50), 
		sourcename varchar(50), 
		author varchar(50),
		title varchar(50),
		description varchar(255),
		url varchar(150),
		urlToImage varchar(255),
		publishedAt varchar(100),
		content	varchar(255),
		PRIMARY KEY (id))`)

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("berhasil membuat tabel lur..")
	}
	defer db.Close()
}