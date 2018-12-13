package nsi

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"strings"
	//"fmt"
	"log"

	_ "github.com/lib/pq"
	"fmt"
)

type nsiBudgetExport struct {
	XMLName xml.Name    `xml:"export"`
	Budgets []nsiBudget `xml:"nsiBudgetList>nsiBudget"`
}

type nsiBudget struct {
	Code   string `xml:"code"`
	Name   string `xml:"name"`
	Actual bool   `xml:"actual"`
}

func (p *nsiBudget) Identify() string {
	log.Println("nsiBudget: ", p.Code)
	return p.Code
}

func (p *nsiBudget) Validate() (bool, error) {
	check := false
	var err error

	check = check || p.Code != ""
	if !check {
		err = errors.New("Budget Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiBudget) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.Name)


	if db.QueryRow("SELECT COUNT(row) FROM nsibudget WHERE code = $1;", p.Code).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`INSERT INTO nsiBudget(code, name, actual) VALUES ('`, p.Code, `','`, p.Name, `','`, p.Actual, `')`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsiBudget CODE: \"", p.Code, "\"")
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiBudget
				SET name = '`, p.Name, `',
					actual = `, p.Actual, `
				WHERE code = '`, p.Code, `' 
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsiBudget CODE: \"", p.Code, "\"")
			}
			return err
		}
	}
	return nil
}

func ParseNsiBudget(data []byte, filename string, db *sql.DB) error {
	var export nsiBudgetExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.Budgets)
	for _, val := range export.Budgets {
		func(o nsiBudget) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("nsiBudget Err. on", o.Code, o.Name, err)
					g_err = err
				}
			}
		}(val)
	}

	log.Println("Parsed ", amount, " nsiBudget;")
	return g_err
}
