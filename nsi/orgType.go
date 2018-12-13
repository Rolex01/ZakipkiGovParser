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

type nsiOrgTypeExport struct {
	XMLName		xml.Name		`xml:"export"`
	OrgTypes    []nsiOrgType	`xml:"nsiOrganizationTypesList>nsiOrganizationType"`
}

type nsiOrgType struct {
	Code		string	`xml:"code"`
	Name		string	`xml:"name"`
	Description	string	`xml:"description"`
}

func (p *nsiOrgType) Validate() (bool, error) {
	check := false
	var err error

	check = p.Code != ""
	if !check {
		err = errors.New("Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiOrgType) Identify() string {
	log.Println("nsiOrgType: ", p.Code, p.Name)
	return p.Name
}

func (p *nsiOrgType) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.Name)
	ESQL(&p.Description)

	if db.QueryRow("SELECT COUNT(row) FROM nsiorgtype WHERE code = $1;", p.Code).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`INSERT INTO nsiorgtype(code, name, description) VALUES ('`, p.Code, `','`, p.Name, `','`, p.Description, `')`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsiorgtype CODE: \"", p.Code, "\"")
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiorgtype
				SET name = '`, p.Name, `',
					description = '`, p.Description, `'
				WHERE code = '`, p.Code, `' 
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsiorgtype CODE: \"", p.Code, "\"")
			}
			return err
		}
	}
	return nil
}

func ParseNsiOrgType(data []byte, filename string, db *sql.DB) error {
	var export nsiOrgTypeExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.OrgTypes)
	for _, val := range export.OrgTypes {
		func(o nsiOrgType) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code, o.Name)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("NsiOrgType Err. on", o.Code, o.Name, err)
					g_err = err
				}
			}
		}(val)
	}

	log.Println("Parsed ", amount, " nsiOrgTypes;")
	return g_err
}
