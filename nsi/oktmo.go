package nsi

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"strings"

	"log"

	_ "github.com/lib/pq"
	"fmt"
	"sync"
	"time"
)

type nsiOKTMOExport struct {
	XMLName xml.Name   `xml:"export"`
	OKTMOs  []nsiOKTMO `xml:"nsiOKTMOList>nsiOKTMO"`
}
type nsiOKTMO struct {
	Code			string	`xml:"code"`
	Parent			string	`xml:"parentCode"`
	Section			string	`xml:"section"`
	Name			string	`xml:"fullName"`
	LastUpdateDate	string	`xml:"lastUpdateDate"`
	Actual			bool	`xml:"actual"`
}

func (p *nsiOKTMO) Identify() string {
	log.Println("NsiOKTMO: ", p.Code)
	return p.Code
}

func (p *nsiOKTMO) Validate() (bool, error) {
	check := false
	var err error
	check = check || p.Code != ""
	if !check {
		err = errors.New("OKTMO Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiOKTMO) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.Parent)
	ESQL(&p.Section)
	ESQL(&p.Name)
	ESQL(&p.LastUpdateDate)

	if p.LastUpdateDate == "" {
		p.LastUpdateDate = `null`
	} else {
		p.LastUpdateDate = fmt.Sprintf(`'%s'`, p.LastUpdateDate)
	}

	if db.QueryRow("SELECT COUNT(row) FROM nsioktmo WHERE code = $1;", p.Code).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsiOKTMO(code, parent, section, fullName, lastUpdateDate, actual)
			VALUES ('`, p.Code, `','`, p.Parent, `','`, p.Section, `',
				'`, p.Name, `',`, p.LastUpdateDate, `::timestamp,`, p.Actual, `
			)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsioktmo CODE: \"", p.Code, "\"")
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiOKTMO
				SET parent = '`, p.Parent, `',
					section = '`, p.Section, `',
					fullName = '`, p.Name, `',
					lastUpdateDate = `, p.LastUpdateDate, `::timestamp,
					actual = `, p.Actual, `
				WHERE code = '`, p.Code, `' 
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsioktmo CODE: \"", p.Code, "\"")
			}
			return err
		}
	}
	return nil
}

func ParseNsiOKTMO(data []byte, filename string, db *sql.DB) error {
	var export nsiOKTMOExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.OKTMOs)

	var Cancel_cnt int = 0
	var ttt int = 0
	var wg sync.WaitGroup
	var f_mod = 100
	var f_sleep = 5
	for _, val := range export.OKTMOs {
		ttt++
		for true {
			if ttt - Cancel_cnt < f_mod {
				break
			}
			log.Println("wait:", ttt, Cancel_cnt)
			time.Sleep(time.Second * time.Duration(f_sleep))
		}
		wg.Add(1)
		func(o nsiOKTMO) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("nsiOKTMO Err. on", o.Code, o.Name, err)
					g_err = err
				}
			}
			Cancel_cnt++
			wg.Done()
		}(val)
	}
	wg.Wait()

	log.Println("Parsed ", amount, " nsiOKTMO;")
	return g_err
}
