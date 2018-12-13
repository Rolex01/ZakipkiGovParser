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

type nsiPurchasePreferenceListExport struct {
	XMLName					xml.Name				`xml:"export"`
	PurchasePreferences		[]nsiPurchasePreference	`xml:"nsiPurchasePreferenceList>nsiPurchasePreference"`
}
type nsiPurchasePreference struct {
	Id				int64	`xml:"id"`
	Name			string	`xml:"name"`
	ShortName		string	`xml:"shortName"`
	Type			string	`xml:"type"`
	PrefEstimateApp	bool	`xml:"prefEstimateApp"`
	prefValue		int64	`xml:"prefValue"`
	Code			string	`xml:"tenderPlanPurchaseGroups>code"`
	Actual			bool	`xml:"actual"`
}

func (p *nsiPurchasePreference) Identify() string {
	log.Println("nsiPurchasePreference: ", p.ShortName, p.Code)
	return p.Code
}

func (p *nsiPurchasePreference) Validate() (bool, error) {
	check := false
	var err error
	check = check || p.Name != ""
	if !check {
		err = errors.New("PurchasePreferences NAME is empty or not valid!")
	}
	return check, err
}

func (p *nsiPurchasePreference) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Code)
	ESQL(&p.Name)
	ESQL(&p.ShortName)
	ESQL(&p.Type)

	if db.QueryRow("SELECT COUNT(row) FROM nsipurchasepreference WHERE id = $1;", p.Id).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsiPurchasePreference(id, name, shortName, type, prefEstimateApp, prefValue, code, actual)
			VALUES (`, p.Id, `,'`, p.Name, `','`, p.ShortName, `','`, p.Type, `',`, p.PrefEstimateApp, `,
				`, p.prefValue, `,'`, p.Code, `',`, p.Actual, `
			)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsiPurchasePreference ID:", p.Id, ", CODE: \"", p.Code, "\"")
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiPurchasePreference
				SET name = '`, p.Name, `',
					shortName = '`, p.ShortName, `',
					type = '`, p.Type, `',
					prefEstimateApp = `, p.PrefEstimateApp, `,
					prefValue = `, p.prefValue, `,
					code = '"`, p.Code, `',
					actual = `, p.Actual, `
				WHERE id = `, p.Id, ` 
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsiPurchasePreference ID:", p.Id, ", CODE: \"", p.Code, "\"")
			}
			return err
		}
	}
	return nil
}

func ParseNsiPurchasePreference(data []byte, filename string, db *sql.DB) error {
	var export nsiPurchasePreferenceListExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.PurchasePreferences)

	var Cancel_cnt int = 0
	var ttt int = 0
	var wg sync.WaitGroup
	var f_mod = 100
	var f_sleep = 5
	for _, val := range export.PurchasePreferences {
		ttt++
		for true {
			if ttt - Cancel_cnt < f_mod {
				break
			}
			log.Println("wait:", ttt, Cancel_cnt)
			time.Sleep(time.Second * time.Duration(f_sleep))
		}
		wg.Add(1)
		func(o nsiPurchasePreference) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.Name)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("nsiPurchasePreference Err. on", o.Code, o.Name, err)
					g_err = err
				}
			}
			Cancel_cnt++
			wg.Done()
		}(val)
	}
	wg.Wait()

	log.Println("Parsed ", amount, " nsiPurchasePreference;")
	return g_err
}
