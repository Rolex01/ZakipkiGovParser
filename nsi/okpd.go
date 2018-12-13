package nsi

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"strings"
	//"fmt"
	"log"

	_ "github.com/lib/pq"
	//"regexp"
	"fmt"
)

type nsiOKPDExport struct {
	XMLName xml.Name  `xml:"export"`
	OKPDs   []nsiOKPD `xml:"nsiOKPDList>nsiOKPD"`
	OKPD2s  []nsiOKPD `xml:"nsiOKPD2List>nsiOKPD2"`
}

type nsiOKPD struct {
	Id			int64  `xml:"id"`
	ParentId	int64  `xml:"parentId"`
	Code		string `xml:"code"`
	Parent		string `xml:"parentCode"`
	Name		string `xml:"name"`
	Comment		string `xml:"comment"`
	Actual		bool   `xml:"actual"`
	OKPD_v		int64  `xml:"version"`
}

func (p *nsiOKPD) Identify() string {
	return p.Code
}

func (p *nsiOKPD) Validate() (bool, error) {
	check := false
	var err error
	check = check || p.Code != ""
	if !check {
		err = errors.New("OKPD Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiOKPD) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.Parent)
	ESQL(&p.Name)
	ESQL(&p.Comment)

	if db.QueryRow("SELECT COUNT(row) FROM nsiokpd WHERE code = $1 AND okpd_version = $2;", p.Code, p.OKPD_v).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsiOKPD(id, parentId, code, parent, name, comment, actual, okpd_version)
			VALUES (`, p.Id, `,`, p.ParentId, `,'`, p.Code, `','`, p.Parent, `','`, p.Name, `','`, p.Comment, `',`, p.Actual, `,`, p.OKPD_v, `)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsiOKPD", p.OKPD_v, " CODE: \"", p.Code, "\"")
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiOKPD
				SET id = `, p.Id, `,
					parentId = `, p.ParentId, `,
					parent = '`, p.Parent, `',
					name = '`, p.Name, `',
					comment = '`, p.Comment, `',
					actual = `, p.Actual, `
				WHERE code = '`, p.Code, `'
					AND okpd_version = `, p.OKPD_v, `
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsiOKPD", p.OKPD_v, " CODE: \"", p.Code, "\"")
			}
			return err
		}
	}
	return nil
}

func ParseNsiOKPD(data []byte, filename string, db *sql.DB) error {
	var export nsiOKPDExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"

	for _, val := range export.OKPDs {
		func(o nsiOKPD) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code)
			} else {
				o.OKPD_v = 1
				if err = o.Save(db, upd); err != nil {
					log.Println("nsiOKPD_1 Err. on", o.Code, err)
					g_err = err
				}
			}
		}(val)
	}

	for _, val := range export.OKPD2s {
		func(o nsiOKPD) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code)
			} else {
				o.OKPD_v = 2
				if err = o.Save(db, upd); err != nil {
					log.Println("nsiOKPD_2 Err. on", o.Code, err)
					g_err = err
				}
			}
		}(val)
	}

	log.Printf("Parsed %d nsiOKPD_1; %d nsiOKPD_2", len(export.OKPDs), len(export.OKPD2s))
	return g_err
}
