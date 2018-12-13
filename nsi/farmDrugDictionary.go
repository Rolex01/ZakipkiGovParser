package nsi

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"strings"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type nsiFarmDrugsDictionaryExport struct {
	XMLName			xml.Name			`xml:"export"`
	FarmDrugsDics	[]nsiFarmDrugsDic	`xml:"nsiFarmDrugsDictionary>nsiFarmDrugDictionary>MNNInfo"`
}

type nsiFarmDrugsDic struct {
	Group_Code			string			`xml:"groupInfo>groupCode"`
	Group_parentCode	string			`xml:"groupInfo>parentGroupCode"`
	Group_Name			string			`xml:"groupInfo>groupName"`
	Group_actual		bool			`xml:"groupInfo>actual"`
	MNNCode				string			`xml:"MNNCode"`
	MNNDrugCode			string			`xml:"MNNDrugCode"`
	MNNExternalCode		string			`xml:"MNNExternalCode"`
	MNNHash				string			`xml:"MNNHash"`
	MNNName				string			`xml:"MNNName"`
	Ftg_Code			string			`xml:"ftgInfo>ftgCode"`
	Ftg_Name			string			`xml:"ftgInfo>ftgName"`
	MedForm_Code		string			`xml:"medicamentalFormInfo>medicamentalFormCode"`
	MedForm_Name		string			`xml:"medicamentalFormInfo>medicamentalFormName"`
	// DosagesInfo			[]dosageInfo	`xml:"dosagesInfo>dosageInfo"`
	DosagesInfo struct {
		Code		string	`xml:"dosageCode"`
		Name		string	`xml:"dosageName"`
		GRLSValue	string	`xml:"dosageGRLSValue"`
		OKEI		string	`xml:"dosageOKEI>code"`
		Value		float64	`xml:"dosageValue"`
		UserOKEI	string	`xml:"dosageUser>dosageUserOKEI>code"`
		// Factors		[]dosageFactor 		`xml:"dosageFactors>dosageFactor"`
		// FactorRange	[]dosageFactorRange `xml:"dosageFactorRanges>dosageFactorRange"`
	}	`xml:"dosagesInfo>dosageInfo"`
	// AthsInfo			[]athInfo		`xml:"athsInfo>athInfo"`
	IsZNVLP				bool			`xml:"isZNVLP"`
	IsNarcotic			bool			`xml:"isNarcotic"`
	OKPD2				string			`xml:"OKPD2>code"`
	// CreateDate		string			`xml:"createDate"`
	// StartDate		string			`xml:"startDate"`
	// EndDate			string			`xml:"endDate"`
	// PricesInfo		[]priceInfo		`xml:"pricesInfo>priceInfo"`
	// ChangeDate		string			`xml:"changeDate"`
	// LastChangeDate	string			`xml:"lastChangeDate"`
	Actual				bool			`xml:"actual"`
}

type dosageInfo struct {
	Code		string	`xml:"dosageCode"`
	Name		string	`xml:"dosageName"`
	GRLSValue	string	`xml:"dosageGRLSValue"`
	OKEI		string	`xml:"dosageOKEI>code"`
	Value		float64	`xml:"dosageValue"`
	UserOKEI	string	`xml:"dosageUser>dosageUserOKEI>code"`
	// Factors		[]dosageFactor 		`xml:"dosageFactors>dosageFactor"`
	// FactorRange	[]dosageFactorRange `xml:"dosageFactorRanges>dosageFactorRange"`
}

type dosageFactor struct {
	Code	string	`xml:"dosageFactorCode"`
	Value	float64	`xml:"dosageFactorValue"`
}

type dosageFactorRange struct {
	Code	string	`xml:"dosageFactorRangeCode"`
	Min		float64	`xml:"dosageFactorRangeMin"`
	Max		float64	`xml:"dosageFactorRangeMax"`
	Step	float64	`xml:"dosageFactorRangeStep"`
}

type athInfo struct {
	Code			string	`xml:"athCode"`
	ExternalCode	string	`xml:"athExternalCode"`
	Name			string	`xml:"athName"`
}

type priceInfo struct {
	Code		string	`xml:"priceCode"`
	Value		float64	`xml:"priceValue"`
	Type		string	`xml:"priceType"`
	SigmaValue	float64	`xml:"priceSigmaValue"`
	UsageMin	int64	`xml:"usageMin"`
	UsageMax	int64	`xml:"usageMax"`
	Author		string	`xml:"author"`
	CreateDate	string	`xml:"createDate"`
	StartDate	string	`xml:"startDate"`
	EndDate		string	`xml:"endDate"`
}

func (p *nsiFarmDrugsDic) Identify() string {
	return p.MNNCode
}

func (p *nsiFarmDrugsDic) Validate() (bool, error) {
	check := false
	var err error
	check = check || p.MNNCode != ""
	if !check {
		err = errors.New("FarmDrugsDic Code is empty or not valid!")
	}
	return check, err
}

func (p *nsiFarmDrugsDic) Save(db *sql.DB, upd bool) error {
	var c int

	ESQL(&p.Group_Code)
	ESQL(&p.Group_parentCode)
	ESQL(&p.Group_Name)
	ESQL(&p.MNNCode)
	ESQL(&p.MNNDrugCode)
	ESQL(&p.MNNExternalCode)
	ESQL(&p.MNNHash)
	ESQL(&p.MNNName)
	ESQL(&p.Ftg_Code)
	ESQL(&p.Ftg_Name)
	ESQL(&p.MedForm_Code)
	ESQL(&p.MedForm_Name)
	ESQL(&p.DosagesInfo.Code)
	ESQL(&p.DosagesInfo.Name)
	ESQL(&p.DosagesInfo.GRLSValue)
	ESQL(&p.DosagesInfo.OKEI)
	ESQL(&p.DosagesInfo.UserOKEI)
	ESQL(&p.OKPD2)

	if db.QueryRow("SELECT COUNT(row) FROM nsiFarmDrugsDic WHERE MNNCode = $1;", p.MNNCode).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			INSERT INTO nsiFarmDrugsDic (
				Group_Code, Group_parentCode, Group_Name, Group_actual,
				MNNCode, MNNDrugCode, MNNExternalCode, MNNHash, MNNName,
				Ftg_Code, Ftg_Name, MedForm_Code, MedForm_Name,
				Dosages_Code, Dosages_Name, Dosages_GRLSValue, Dosages_OKEI, Dosages_Value, Dosages_UserOKEI,
				IsZNVLP, IsNarcotic, OKPD2, Actual 
			) VALUES (
				'`, p.Group_Code, `','`, p.Group_parentCode, `','`, p.Group_Name, `',`, p.Group_actual, `,
				'`, p.MNNCode, `','`, p.MNNDrugCode, `','`, p.MNNExternalCode, `','`, p.MNNHash, `','`, p.MNNName, `',
				'`, p.Ftg_Code, `','`, p.Ftg_Name, `','`, p.MedForm_Code, `','`, p.MedForm_Name, `',
				'`, p.DosagesInfo.Code, `','`, p.DosagesInfo.Name, `','`, p.DosagesInfo.GRLSValue, `','`, p.DosagesInfo.OKEI, `',`, p.DosagesInfo.Value, `,'`, p.DosagesInfo.UserOKEI, `',
				`, p.IsZNVLP, `,`, p.IsNarcotic, `,'`, p.OKPD2, `',`, p.Actual, `
			)
		`)

		_, err := db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert:", my_query)
		} else {
			log.Println(`Insert nsiFarmDrugsDic CODE: "`, p.MNNCode, `"`)
		}
		return err
	} else {
		if upd {
			my_query := fmt.Sprint(`
				UPDATE nsiFarmDrugsDic
				SET Group_Code = `, p.Group_Code, `,
					Group_parentCode = `, p.Group_parentCode, `,
					Group_Name = '`, p.Group_Name, `',
					Group_actual = `, p.Group_actual, `,
					MNNCode = '`, p.MNNCode, `',
					MNNDrugCode = '`, p.MNNDrugCode, `',
					MNNExternalCode = '`, p.MNNExternalCode, `',
					MNNName = '`, p.MNNName, `',
					Ftg_Code = '`, p.Ftg_Code, `',
					Ftg_Name = '`, p.Ftg_Name, `',
					MedForm_Code = '`, p.MedForm_Code, `',
					MedForm_Name = '`, p.MedForm_Name, `',
					Dosages_Code = '`, p.DosagesInfo.Code, `',
					Dosages_Name = '`, p.DosagesInfo.Name, `',
					Dosages_GRLSValue = '`, p.DosagesInfo.GRLSValue, `',
					Dosages_OKEI = '`, p.DosagesInfo.OKEI, `',
					Dosages_Value = `, p.DosagesInfo.Value, `,
					Dosages_UserOKEI = '`, p.DosagesInfo.UserOKEI, `',
					IsZNVLP = `, p.IsZNVLP, `,
					IsNarcotic = `, p.IsNarcotic, `,
					OKPD2 = '`, p.OKPD2, `',
					actual = `, p.Actual, `
				WHERE MNNCode = '`, p.MNNCode, `'
			`)

			_, err := db.Exec(my_query)
			if err != nil {
				log.Println("ERROR Update:", my_query)
			} else {
				log.Println(`Update nsiFarmDrugsDic CODE: "`, p.MNNCode, `"`)
			}
			return err
		}
	}

	return nil
}

func ParseNsiFarmDrugsDic(data []byte, filename string, db *sql.DB) error {
	var export nsiFarmDrugsDictionaryExport
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Fatal(err)
		return err
	}
	filename = strings.Split(filename, "_")[1]
	upd := filename == "inc"

	for _, val := range export.FarmDrugsDics {
		func(o nsiFarmDrugsDic) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.MNNCode)
			} else {
				if err = o.Save(db, upd); err != nil {
					log.Println("nsiFarmDrugsDic Err. on", o.MNNCode, err)
					g_err = err
				}
			}
		}(val)
	}

	log.Printf("Parsed %d nsiFarmDrugsDic;", len(export.FarmDrugsDics))
	return g_err
}
