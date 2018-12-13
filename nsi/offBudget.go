package nsi

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"strings"
	//"fmt"
	"log"

	//_ "bitbucket.org/crmsib/kraken2/schm"
	_ "github.com/lib/pq"
	//"regexp"
	"sync"
	"time"
	"fmt"
)

type nsiOffBudgetExport struct {
	XMLName xml.Name       `xml:"export"`
	Budgets []nsiOffBudget `xml:"nsiOffBudgetList>nsiOffBudget"`
}
type nsiOffBudget struct {
	Code			string	`xml:"code"`
	Name			string	`xml:"name"`
	SubsystemType	string	`xml:"subsystemType"`
	Actual			bool	`xml:"actual"`
}

func (p *nsiOffBudget) Identify() string {
	log.Println("nsiOffBudget: ", p.Code)
	return p.Code
}

func (p *nsiOffBudget) Validate() (bool, error) {
	check := false
	var err error
	check = check || p.Code != ""
	if !check {
		err = errors.New("Budget Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiOffBudget) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.Name)
	ESQL(&p.SubsystemType)

	if db.QueryRow("SELECT count(row) from nsiOffBudget where code = $1;", p.Code).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsiOffBudget(code, name, subsystemType, actual)
			VALUES ('`, p.Code, `','`, p.Name, `','`, p.SubsystemType, `',`, p.Actual, `)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsiOffBudget CODE: \"", p.Code, "\"")
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiOffBudget
				SET name = '`, p.Name, `',
					subsystemType = '`, p.SubsystemType, `', 
					actual = `, p.Actual, `
				WHERE code = '`, p.Code, `' 
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsiOffBudget CODE: \"", p.Code, "\"")
			}
			return err
		}
	}

	return nil
}

func ParseNsiOffBudget(data []byte, filename string, db *sql.DB) error {
	var export nsiOffBudgetExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.Budgets)

	var Cancel_cnt int = 0
	var ttt int = 0
	var wg sync.WaitGroup
	var f_mod = 100
	var f_sleep = 5
	for _, val := range export.Budgets {
		ttt++
		for true {
			if ttt - Cancel_cnt < f_mod {
				break
			}
			log.Println("wait:", ttt, Cancel_cnt)
			time.Sleep(time.Second * time.Duration(f_sleep))
		}
		wg.Add(1)
		func(o nsiOffBudget) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Code)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("nsiOffBudget Err. on", o.Code, o.Name, err)
					g_err = err
				}
			}
			Cancel_cnt++
			wg.Done()
		}(val)
	}
	wg.Wait()

	log.Println("Parsed ", amount, " nsiOffBudget;")
	return g_err
}
