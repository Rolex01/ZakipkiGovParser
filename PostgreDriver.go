package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	User        string
	Password    string
	Host        string
	Port        string
	DBName      string
	MaxConn     int
	MaxIdleConn int
}

func (p *Database) PSQLConnString() string {
	//return "postgres://" + p.User + ":+" + p.Password + "@" + p.Host + ":" + p.Port + "/" + p.DBName + "?sslmode=disable"
	return "user=" + p.User + " password=" + p.Password + " dbname=" + p.DBName + " host=" + p.Host + " sslmode=disable" + " connect_timeout=1"
}

func InitPSQL(p *Database) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", p.PSQLConnString())
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(p.MaxConn)
	db.SetMaxIdleConns(p.MaxIdleConn)
	return db, err
}

func SqlDateTime(d string) string {
	// '0000-00-00' or '0000-00-00 00:00:00'
	Moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal("Time error:", err)
	}
	t, _ := time.Parse(time.RFC3339, d)
	return t.In(Moscow).String()[:19]
}

func ShortDate(m string, y string) string {
	var Month, Year string
	if len(string(m)) == 1 {
		Month = "0" + string(m)
	} else {
		Month = string(m)
	}
	Year = string(y)
	return Year + "-" + Month + "-01"
}
