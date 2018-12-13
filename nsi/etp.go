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

type nsiETPExport struct {
	XMLName		xml.Name		`xml:"export"`
	ETPs		[]nsiETP		`xml:"nsiETPs>nsiETP"`
}

type nsiETP struct {
	Code		string	`xml:"code"`
	Name		string	`xml:"name"`
	Description	string	`xml:"description"`
	Phone		string	`xml:"phone"`
	Address		string	`xml:"address"`
	Email		string	`xml:"email"`
	FullName	string	`xml:"fullName"`
	INN			string	`xml:"INN"`
	KPP			string	`xml:"KPP"`
	Actual		bool	`xml:"actual"`
}

func (p *nsiETP) Validate() (bool, error) {
	check := false
	var err error

	check = p.Code != ""
	if !check {
		err = errors.New("Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiETP) Identify() string {
	log.Println("nsiETP: ", p.Code, p.Name)
	return p.Name
}

func (p *nsiETP) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.Name)
	ESQL(&p.Description)
	ESQL(&p.Phone)
	ESQL(&p.Address)
	ESQL(&p.Email)
	ESQL(&p.FullName)
	ESQL(&p.INN)
	ESQL(&p.KPP)

	if db.QueryRow("SELECT COUNT(row) FROM nsietp WHERE code = $1;", p.Code).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsietp (
				code, name, description, phone, address,
				email, fullName, INN, KPP, actual
			) VALUES (
				'`, p.Code, `','`, p.Name, `','`, p.Description, `','`, p.Phone, `','`, p.Address, `',
				'`, p.Email, `','`, p.FullName, `','`, p.INN, `','`, p.KPP, `',`, p.Actual, `
			)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsietp CODE: \"", p.Code, "\"")
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsietp
				SET name = '`, p.Name, `',
					description = '`, p.Description, `',
					phone = '`, p.Phone, `',
					address = '`, p.Address, `',
					email = '`, p.Email, `',
					fullName = '`, p.FullName, `',
					INN = '`, p.INN, `',
					KPP = '`, p.KPP, `',
					actual = `, p.Actual, `
				WHERE code = '`, p.Code, `' 
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsietp CODE: \"", p.Code, "\"")
			}
			return err
		}
	}

	return nil
}

func ParseNsiETP(data []byte, filename string, db *sql.DB) error {
	var export nsiETPExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.ETPs)
	for _, val := range export.ETPs {
		func(o nsiETP) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code, o.Name)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("NsiETP Err. on", o.Code, o.Name, err)
					g_err = err
				}
			}
		}(val)
	}

	log.Println("Parsed ", amount, " nsiETP;")
	return g_err
}
