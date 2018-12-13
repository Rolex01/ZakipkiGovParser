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

type nsiOrgExport struct {
	XMLName xml.Name   `xml:"export"`
	Orgs    []nsiOrg `xml:"nsiOrganizationList>nsiOrganization"`
}

type nsiOrg struct {
	RegNum				string				`xml:"regNumber"`
	ShortName			string				`xml:"shortName"`
	FullName			string				`xml:"fullName"`
	OKTMO				string				`xml:"OKTMO>code"`
	PostalAddress		string				`xml:"postalAddress"`
	Email				string				`xml:"email"`
	INN					string				`xml:"INN"`
	KPP					string				`xml:"KPP"`
	OGRN				string				`xml:"OGRN"`
	OKVED				string				`xml:"OKVED"`
	OrganizationRole	string				`xml:"organizationRoles>organizationRoleItem>organizationRole"`
	OrganizationType	string				`xml:"organizationType>code"`
	Url					string				`xml:"url"`
	ContactPerson		ContactPersonRef	`xml:"contactPerson"`
	SubordinationType	string				`xml:"subordinationType"`
	RegistrationDate	string				`xml:"registrationDate"`
	TimeZone			int					`xml:"timeZone"`
	TimeZoneUtcOffset	string				`xml:"timeZoneUtcOffset"`
	Actual				bool				`xml:"actual"`
	Register			bool				`xml:"register"`
}

type ContactPersonRef struct {
	First  string `xml:"firstName"`
	Last   string `xml:"lastName"`
	Middle string `xml:"middleName"`
}


func (p *nsiOrg) Identify() string {
	log.Println("NsiOrg: ", p.RegNum)
	return p.RegNum
}

func (p *nsiOrg) Validate() (bool, error) {
	check := false
	var err error

	check = p.RegNum != ""
	if !check {
		err = errors.New("RegNum is empty or not valid!")
	}
	return check, err
}

func ESQL(str *string) {
	*str = strings.Replace(*str, "'", "''", -1)
	return
}

func (p *nsiOrg) Save(db *sql.DB, upd bool) error {
	var c int
	FullContactPerson := p.ContactPerson.First + " " + p.ContactPerson.Middle + " " + p.ContactPerson.Last

	ESQL(&p.RegNum)
	ESQL(&p.ShortName)
	ESQL(&p.FullName)
	ESQL(&p.OKTMO)
	ESQL(&p.PostalAddress)
	ESQL(&p.Email)
	ESQL(&p.INN)
	ESQL(&p.KPP)
	ESQL(&p.OGRN)
	ESQL(&p.OKVED)
	ESQL(&p.OrganizationRole)
	ESQL(&p.OrganizationType)
	ESQL(&p.Url)
	ESQL(&FullContactPerson)
	ESQL(&p.SubordinationType)
	ESQL(&p.RegistrationDate)
	ESQL(&p.TimeZoneUtcOffset)

	if p.RegistrationDate == "" {
		p.RegistrationDate = `null`
	} else {
		p.RegistrationDate = fmt.Sprintf(`'%s'`, p.RegistrationDate)
	}

	if db.QueryRow("SELECT COUNT(row) FROM nsiorg WHERE regnum = $1;", p.RegNum).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsiorg(
				regnum, shortname, fullname , oktmo, postaladdress,
				email, inn, kpp, ogrn, okved,
				organizationRole, organizationType, url, contactperson, subordinationtype,
				registrationDate, timeZone, timeZoneUtcOffset, actual
			) VALUES (
				'`, p.RegNum, `','`, p.ShortName, `','`, p.FullName, `','`, p.OKTMO, `','`, p.PostalAddress, `',
				'`, p.Email, `','`, p.INN, `','`, p.KPP, `','`, p.OGRN, `','`, p.OKVED, `',
				'`, p.OrganizationRole, `','`, p.OrganizationType, `','`, p.Url, `','`, FullContactPerson, `','`, p.SubordinationType, `',
				`, p.RegistrationDate, `::date,`, p.TimeZone, `,'`, p.TimeZoneUtcOffset, `',`, p.Actual, `)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println("Insert nsiorg REGNUM: \"", p.RegNum, "\"")
		}
		//log.Println("NsiOrg Inserted: ", p.RegNum, err)
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiorg
				SET shortname = '`, p.ShortName, `',
					fullname = '`, p.FullName, `',
					oktmo = '`, p.OKTMO, `',
					postaladdress = '`, p.PostalAddress, `',
					email = '`, p.Email, `',
					inn = '`, p.INN, `',
					kpp = '`, p.KPP, `',
					ogrn = '`, p.OGRN, `',
					okved = '`, p.OKVED, `',
					organizationRole = '`, p.OrganizationRole, `',
					organizationType = '`, p.OrganizationType, `',
					url = '`, p.Url, `',
					contactperson = '`, FullContactPerson, `',
					subordinationtype = '`, p.SubordinationType, `',
					registrationDate = `, p.RegistrationDate, `::date,
					timeZone = `, p.TimeZone, `,
					timeZoneUtcOffset = '`, p.TimeZoneUtcOffset, `',
					actual = `, p.Actual, `
				WHERE RegNum = '`, p.RegNum, `' 
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println("Update nsiorg REGNUM: \"", p.RegNum, "\"")
			}
			return err
		}
	}
	return nil
}

func ParseNsiOrg(data []byte, filename string, db *sql.DB) error {
	var export nsiOrgExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"
	amount := len(export.Orgs)
	for _, val := range export.Orgs {
		func(o nsiOrg) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.RegNum)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("NsiOrg Err. on", o.RegNum, err)
					g_err = err
				}
			}
		}(val)
	}

	log.Println("Parsed ", amount, " nsiOrgs;")
	return g_err
}
