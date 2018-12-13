package nsi

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"strings"

	"log"
	_ "github.com/lib/pq"
	"fmt"
)

type nsiPlacingWayExport struct {
	XMLName			xml.Name		`xml:"export"`
	PlacingWays    []nsiPlacingWay	`xml:"nsiPlacingWayList>nsiPlacingWay"`
}

type nsiPlacingWay struct {
	PlacingWayId 	int		`xml:"placingWayId"`
	Code 			string	`xml:"code"`
	Name 			string	`xml:"name"`
	Type 			string	`xml:"type"`
	SubsystemType	string	`xml:"subsystemType"`
	Actual 			bool	`xml:"actual"`
	IsProcedure		bool	`xml:"isProcedure"`
}

func (p *nsiPlacingWay) Identify() string {
	log.Println("NsiPlacingWay: ", p.Code, p.Name, p.Type)
	return p.Code
}

func (p *nsiPlacingWay) Validate() (bool, error) {
	check := false
	var err error

	check = p.Code != ""
	if !check {
		err = errors.New("Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiPlacingWay) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.Name)
	ESQL(&p.Type)
	ESQL(&p.SubsystemType)

	if db.QueryRow("SELECT COUNT(row) FROM nsiPlacingWay WHERE placingWayId = $1;", p.PlacingWayId).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsiPlacingWay(placingWayId, code, name , type, subsystemType, actual, isProcedure)
			VALUES (`, p.PlacingWayId, `,'`, p.Code, `','`, p.Name, `','`, p.Type, `','`, p.SubsystemType, `',`, p.Actual, `,`, p.IsProcedure, `)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsiPlacingWay PlacingWayId: \"", p.PlacingWayId, "\"")
		}
		//log.Println("NsiOrg Inserted: ", p.RegNum, err)
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiPlacingWay
				SET code = '`, p.Code, `',
					name = '`, p.Name, `',
					type = '`, p.Type, `',
					subsystemType = '`, p.SubsystemType, `',
					actual = `, p.Actual, `,
					isProcedure = `, p.IsProcedure, `
				WHERE placingWayId = `, p.PlacingWayId, `
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsiPlacingWay PlacingWayId: \"", p.PlacingWayId, "\"")
			}
			return err
		}
	}
	return nil
}

func ParseNsiPlacingWay(data []byte, filename string, db *sql.DB) error {
	var export nsiPlacingWayExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.PlacingWays)
	for _, val := range export.PlacingWays {
		func(o nsiPlacingWay) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("NsiPlacingWay Err. on", o.Code, err)
					g_err = err
				}
			}
		}(val)
	}

	log.Println("Parsed ", amount, " nsiPlacingWays;")
	return g_err
}
