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

type nsiOKEIExport struct {
	XMLName xml.Name  `xml:"export"`
	OKEIs   []nsiOKEI `xml:"nsiOKEIList>nsiOKEI"`
}

type nsiOKEI struct {
	Code                string `xml:"code"`
	FullName            string `xml:"fullName"`
	SectionCode         string `xml:"section>code"`
	SectionName         string `xml:"section>name"`
	GroupId             int64  `xml:"group>id"`
	GroupName           string `xml:"group>name"`
	LocalName           string `xml:"localName"`
	InternationalName   string `xml:"internationalName"`
	LocalSymbol         string `xml:"localSymbol"`
	InternationalSymbol string `xml:"internationalSymbol"`
	Actual              bool   `xml:"actual"`
	isTempForKTRU		bool   `xml:"isTemporaryForKTRU"`
}

func (p *nsiOKEI) Identify() string {
	log.Println("nsiOKEI: ", p.Code, p.FullName)
	return p.Code
}

func (p *nsiOKEI) Validate() (bool, error) {
	check := false
	var err error
	check = check || p.Code != ""
	if !check {
		err = errors.New("OKPD Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiOKEI) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.FullName)
	ESQL(&p.SectionCode)
	ESQL(&p.SectionName)
	ESQL(&p.GroupName)
	ESQL(&p.LocalName)
	ESQL(&p.InternationalName)
	ESQL(&p.LocalSymbol)
	ESQL(&p.InternationalSymbol)

	if db.QueryRow("SELECT COUNT(row) FROM nsiokei WHERE code = $1;", p.Code).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsiOKEI(
				Code, FullName, SectionCode, SectionName, GroupId,
				GroupName, LocalName, InternationalName, LocalSymbol, InternationalSymbol,
				Actual, isTemporaryForKTRU
			) VALUES (
				'`, p.Code, `','`, p.FullName, `','`, p.SectionCode, `','`, p.SectionName, `',`, p.GroupId, `,
				'`, p.GroupName, `','`, p.LocalName, `','`, p.InternationalName, `','`, p.LocalSymbol, `','`, p.InternationalSymbol, `',
				`, p.Actual, `,`, p.isTempForKTRU, `)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsiOKEI CODE: \"", p.Code, "\"")
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiOKEI
				SET FullName = '`, p.FullName, `',
					SectionCode = '`, p.SectionCode, `',
					SectionName = '`, p.SectionName, `',
					GroupId = `, p.GroupId, `,
					GroupName = '`, p.GroupName, `',
					LocalName = '`, p.LocalName, `',
					InternationalName = '`, p.InternationalName, `',
					LocalSymbol = '`, p.LocalSymbol, `',
					InternationalSymbol = '`, p.InternationalSymbol, `',
					Actual = `, p.Actual, `,
					isTemporaryForKTRU = `, p.isTempForKTRU, `
				WHERE code = '`, p.Code, `'
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsiOKEI CODE: \"", p.Code, "\"")
			}
			return err
		}
	}
	return nil
}

func ParseNsiOKEI(data []byte, filename string, db *sql.DB) error {
	var export nsiOKEIExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.OKEIs)
	for _, val := range export.OKEIs {
		func(o nsiOKEI) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("nsiOKEI Err. on", o.Code, o.FullName, err)
					g_err = err
				}
			}
		}(val)
	}

	log.Println("Parsed ", amount, " nsiOKEI;")
	return g_err
}
