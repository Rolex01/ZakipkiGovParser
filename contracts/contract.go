package contracts

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"time"

	"strings"

	std "bitbucket.org/crmsib/parser_gov/common"
	"bitbucket.org/crmsib/parser_gov/regions"
	shards "bitbucket.org/crmsib/parser_gov/shards"
	log "github.com/Sirupsen/logrus"
)

// ALTER TABLE the_table ALTER COLUMN col_name TYPE integer USING (trim(col_name)::integer);

// Contract scheme according to Protocol v. 5.0 for 2015 y.
type Contract struct {
	XMLName xml.Name `xml:"export"`
	//SchemeVersion []xml.Attr `xml:"contract>schemeVersion,attr"`
	Id            uint64 `xml:"contract>id"` // Warning: http://go-database-sql.org/surprises.html
	ExternalId    string `xml:"contract>externalId"`
	PublishDate   string `xml:"contract>publishDate"`
	VersionNumber int    `xml:"contract>versionNumber"`
	Foundation    struct {
		Old struct {
			NotificationNumber string `xml:"notificationNumber"`
			LotNumber          int    `xml:"lotNumber"`
			Placing            string `xml:"placing"`
			SingleCustomer     string `xml:"singleCustomer"`
			// Missed: singleCustomer
		} `xml:"oosOrder"`
		// Missed: singleCustomer
		NotOOSPlacing         string `xml:"notOosOrder>placing"`
		Placing44             string `xml:"notOosOrder>placing44FZ"`
		Placing94             string `xml:"notOosOrder>placing94FZ"`
		OtherPlacing          string `xml:"other>placing"`
		OtherNoficationNumber string `xml:"other>notificationNumber"`

		Fcs struct {
			Order struct {
				NotificationNumber string `xml:"notificationNumber"`
				LotNumber          int    `xml:"lotNumber"`
				Placing            string `xml:"placing"`
				SingleCustomer     struct {
					Reason string `xml:"reason"`
				} `xml:"singleCustomer"`
				// Missed: singleCustomer
			} `xml:"order"`
			NotificationNumber string `xml:"notificationNumber"`
			LotNumber          int    `xml:"lotNumber"`
			Placing            string `xml:"placing"`
			// Missed singleCustomer
			NotOOSPlacing string `xml:"notOosOrder>placing"`
		} `xml:"fcsOrder"`
	} `xml:"contract>foundation"`

	Finances struct {
		Budget         string `xml:"budgetFunds>budget>name"` // Name only
		BudgetOld      string `xml:"budget>name"`
		ExtraBudget    string `xml:"extrabudgetFunds>extrabudget>code"`
		ExtraBudgetOld string `xml:"extrabudget>name"`
		OKTMO          string `xml:"OKTMO>code"` // Code only! TODO: Makes all OKTMO - code only!
		Source         string `xml:"financeSource"`
		Level          string `xml:"budgetFunds>budgetLevel"`
		LevelOld       string `xml:"budgetLevel"`
	} `xml:"contract>finances"`
	ProtocolDate string `xml:"contract>protocolDate"`
	// Missed: documentBase
	SignDate string `xml:"contract>signDate"`
	ExecDate struct {
		Month string `xml:"month"`
		Year  string `xml:"year"`
	} `xml:"contract>executionDate"`
	ExecDateEnd string `xml:"contract>executionPeriod>endDate"`
	RegNum      string `xml:"contract>regNum"`
	IKU         string `xml:"contract>regNumIKU"`
	Number      string `xml:"contract>number"`
	PriceInfo   struct {
		Price    float64 `xml:"price"`
		Currency string  `xml:"currency>code"`
		Rate     float64 `xml:"currencyRate>rate"`
	} `xml:"contract>priceInfo"`
	Price  float64 `xml:"contract>price"`
	Placer string  `xml:"contract>placer>responsibleOrg>regNum"`
	// Missed: subContractors
	//ExecutionPeriod struct{} `xml:"contract>executionPeriod"`
	// Missed: enforcement
	// Missed: energyServiceContractInfo
	Products        []Product  `xml:"contract>products>product"`
	ProductsChanges struct{}   `xml:"contract>products>productChange"`
	Suppliers       []Supplier `xml:"contract>suppliers>supplier"`
	Href            string     `xml:"contract>href"`
	// Missed printForm
	// Missed scanDocuments
	// Missed medicalDocuments
	Attachments []std.Attachment `xml:"contract>scanDocuments>attachment"`
	// Missed: modification
	Status   string `xml:"contract>currentContractStage"`
	Customer struct {
		FullName  string `xml:"fullName"`
		ShortName string `xml:"shortName"`
		Inn       string `xml:"inn"`
		Kpp       string `xml:"kpp"`
		RegNum    string `xml:"regNum"`
	} `xml:"contract>customer"`

	Currency struct {
		Code string `xml:"code"`
		Name string `xml:"name"`
	} `xml:"contract>currency"`
}

type Product struct {
	Sid      uint64  `xml:"sid"`
	OKPD     string  `xml:"OKPD>code"`
	OKPD2    string  `xml:"OKPD2>code"`
	KTRU	 string  `xml:"KTRU>code"`
	Name     string  `xml:"name"`
	OKEI     string  `xml:"OKEI>nationalCode"`
	Price    float32 `xml:"price"`
	PriceRUR float32 `xml:"priceRUR"`
	Quantity float32 `xml:"quantity"`
	Sum      float32 `xml:"sum"`
	SumRUR   float32 `xml:"sumRUR"`
}

type Supplier struct {
	// Fucking hell here ...
	INN        string `xml:"inn"`
	KPP        string `xml:"kpp"`
	Name       string `xml:"organizationName"`
	Address    string `xml:"postAddress"`
	Address2   string `xml:"factualAddress"`
	Phone      string `xml:"contactPhone"`
	FirstName  string `xml:"contactInfo>firstName"`
	LastName   string `xml:"contactInfo>lastName"`
	MiddleName string `xml:"contactInfo>middleName"`
	Company    struct {
		LegalForm     string  `xml:"legalForm>code"` // OKOPF
		FullName      string  `xml:"fullName"`
		ShortName     string  `xml:"shortName"`
		FirmName      string  `xml:"firmName"`
		Status        string  `xml:"status"`
		ContractPrice float32 `xml:"contractPrice"`
		OKPO          string  `xml:"OKPO"`
		INN           string  `xml:"INN"`
		KPP           string  `xml:"KPP"`
		RegDate       string  `xml:"registrationDate"`
		OKTMO         string  `xml:"OKTMO"`
		Address       string  `xml:"address"`
		Info          struct {
			First  string `xml:"firstName"`
			Last   string `xml:"lastName"`
			Middle string `xml:"middleName"`
		} `xml:"contractInfo"`
		Email string `xml:"contactEmail"`
		Phone string `xml:"contactPhone"`
	} `xml:"legalEntityRF"`
	ForeignCompany struct {
		FullName     string `xml:"fullName"`
		ShortName    string `xml:"shortName"`
		FirmName     string `xml:"firmName"`
		FirmNameLat  string `xml:"firmNameLat"`
		TaxPayerCode string `xml:"taxPayerCode"`
		INN          string `xml:"registerInRFTaxBodies>INN"`
		KPP          string `xml:"registerInRFTaxBodies>KPP"`
		RegDate      string `xml:"registerInRFTaxBodies>registrationDate"`
		Country      string `xml:"placeOfStayInRegCountry>country"`
		Address      string `xml:"placeOfStayInRegCountry>address"`
		Email        string `xml:"placeOfStayInRegCountry>contractEmail"`
		Phone        string `xml:"placeOfStayInRegCountry>contractPhone"`
		OKTMO        string `xml:"placeOfStayInRF>OKTMO"` // МАТЬ-ПЕРЕМАТЬ!
		AddressRF    string `xml:"placeOfStayInRF>address"`
		EmailRF      string `xml:"placeOfStayInRF>contactEmail"`
		PhoneRF      string `xml:"placeOfStayInRF>contactPhone"`
	} `xml:"legalEntityForeignState"`
	Individual struct {
		First   string `xml:"firstName"`
		Last    string `xml:"lastName"`
		Middle  string `xml:"middleName"`
		RegDate string `xml:"registrationDate"`
		INN     string `xml:"INN"`
		OKTMO   string `xml:"OKTMO"`
		Email   string `xml:"contactEMail"`
		Phone   string `xml:"contactPhone"`
		Address string `xml:"address"`
	} `xml:"individualPersonRF"`
	ForeignIndividual struct {
		First        string `xml:"firstName"`
		Last         string `xml:"lastName"`
		Middle       string `xml:"middleName"`
		FirstLat     string `xml:"firstNameLat"`
		LastLat      string `xml:"lastNameLat"`
		MiddleLat    string `xml:"middleNameLat"`
		TaxPayerCode string `xml:"taxPayerCode"`
		INN          string `xml:"registerInRFTaxBodies>INN"` // МАТЬ-ПЕРЕМАТЬ!
		RegDate      string `xml:"registerInRFTaxBodies>registrationDate"`
		Country      string `xml:"placeOfStayInRegCountry>country"`
		Address      string `xml:"placeOfStayInRegCountry>address"`
		Email        string `xml:"placeOfStayInRegCountry>contactEmail"`
		Phone        string `xml:"placeOfStayInRegCountry>contactPhone"`
		OKTMO        string `xml:"placeOfStayInRF>OKTMO"`
		AddressRF    string `xml:"placeOfStayInRF>address"`
		EmailRF      string `xml:"placeOfStayInRF>contactEmail"`
		PhoneRF      string `xml:"placeOfStayInRF>contactPhone"`
		// Missed IsIP
	} `xml:"individualPersonForeignState"`
}

func (c *Contract) String() string {
	return fmt.Sprintf("%s", c.RegNum)
}

func (c *Contract) ListSuppliersINN() string {
	suppliers := c.Suppliers
	res := ""
	for _, s := range suppliers {
		if s.Company.INN != "" {
			res += s.Company.INN + ","
		}
		if s.ForeignCompany.INN != "" {
			res += s.ForeignCompany.INN + ","
		}
		if s.Individual.INN != "" {
			res += s.Individual.INN + ","
		}
		if s.ForeignIndividual.INN != "" {
			res += s.ForeignIndividual.INN + ","
		}
		if s.INN != "" {
			res += s.INN + ","
		}
	}
	if res != "" {
		return res[:len(res)-1]
	}
	return res
}
func (c *Contract) GetRegion(code string) int {
	if c.Customer.Kpp == "" {
		return regions.GetRegion(c.Customer.Inn)
	} else {
		return regions.GetRegion(code)
	}

}
func (c *Contract) ListSuppliersKPP() string {
	suppliers := c.Suppliers
	res := ""
	log.Println(suppliers)
	for _, s := range suppliers {
		thisRes := ""
		if s.Company.KPP != "" {
			thisRes = s.Company.KPP + ","
			res += thisRes
		}
		if s.ForeignCompany.KPP != "" {
			thisRes = s.ForeignCompany.KPP + ","
			res += thisRes
		}

		if s.KPP != "" {
			thisRes += s.KPP + ","
			res += thisRes
		}
		if thisRes == "" {
			if s.Individual.INN != "" || s.ForeignIndividual.First != "" {

			} else {
				log.Println("KPP!", s.Company.KPP, s.ForeignCompany.KPP, s.KPP, s.Individual.INN, s.ForeignIndividual.INN, ";;", s, ";;", c.RegNum, ";;", c.SignDate)
			}
			return ""
		}
	}
	if res != "" {
		return res[:len(res)-1]
	}
	return res
}

func (c *Contract) ListSuppliersName() string {
	suppliers := c.Suppliers
	res := ""
	for _, s := range suppliers {
		if s.Company.INN != "" {
			res += s.Company.FullName + ","
		}
		if s.ForeignCompany.INN != "" {
			res += s.ForeignCompany.FullName + ","
		}
		if s.Individual.INN != "" {
			res += s.Individual.First + " " + s.Individual.Middle + " " + s.Individual.Last + ","
		}
		if s.ForeignIndividual.INN != "" {
			res += s.ForeignIndividual.First + " " + s.ForeignIndividual.Middle + " " + s.ForeignIndividual.Last + ","
		}
		if s.Name != "" {
			res += s.Name + ","
		} else {
			if s.FirstName != "" {
				res += s.FirstName + " " + s.MiddleName + " " + s.LastName + ","
			}
		}
	}
	if res != "" {
		return res[:len(res)-1]
	}
	return res
}

func (c *Contract) ListSuppliersAddress() string {
	suppliers := c.Suppliers
	res := ""
	sep := "~;"
	for _, s := range suppliers {
		if s.Company.Address != "" {
			res += s.Company.Address + sep
		}
		if s.ForeignCompany.Address != "" {
			res += s.ForeignCompany.Address + sep
		}
		if s.Individual.Address != "" {
			res += s.Individual.Address + sep
		}
		if s.ForeignIndividual.Address != "" {
			res += s.ForeignIndividual.Address + sep
		}
		if s.Address != "" {
			res += s.Address + sep
		} else {
			if s.Address2 != "" {
				res += s.Address2 + sep
			}
		}
	}
	if res != "" {
		return res[:len(res)-1]
	}
	return res
}

func (c *Contract) ListSuppliersPhone() string {
	suppliers := c.Suppliers
	res := ""
	sep := ","
	for _, s := range suppliers {
		if s.Company.Phone != "" {
			res += s.Company.Phone + sep
		}
		if s.ForeignCompany.Phone != "" {
			res += s.ForeignCompany.Phone + sep
		}
		if s.Individual.Phone != "" {
			res += s.Individual.Phone + sep
		}
		if s.ForeignIndividual.Phone != "" {
			res += s.ForeignIndividual.Phone + sep
		}
		if s.Phone != "" {
			res += s.Phone + sep
		}
	}
	if res != "" {
		return res[:len(res)-1]
	}
	log.Println(suppliers)
	return res
}

func (c *Contract) GetLotNumber() int {
	if c.Foundation.Old.SingleCustomer != "" || c.Foundation.Fcs.Order.SingleCustomer.Reason != "" {
		return 1
	}
	if c.Foundation.Old.NotificationNumber != "" {
		return c.Foundation.Old.LotNumber
	}
	if c.Foundation.Fcs.NotificationNumber != "" {
		return c.Foundation.Fcs.LotNumber
	}
	if c.Foundation.Fcs.Order.NotificationNumber != "" {
		return c.Foundation.Fcs.Order.LotNumber
	}
	if c.Foundation.OtherNoficationNumber != "" {
		return -1
	}
	return -1
}

func (c *Contract) GetPurchaseNumber() string {
	if c.Foundation.Old.NotificationNumber != "" {
		return c.Foundation.Old.NotificationNumber
	}
	if c.Foundation.Fcs.NotificationNumber != "" {
		return c.Foundation.Fcs.NotificationNumber
	}
	if c.Foundation.Fcs.Order.NotificationNumber != "" {
		return c.Foundation.Fcs.Order.NotificationNumber
	}
	if c.Foundation.OtherNoficationNumber != "" {
		return c.Foundation.OtherNoficationNumber
	}
	//log.Println(c.Foundation)
	return ""
}

func (c *Contract) GetPlacing() string {
	var pacing string
	if c.Foundation.Old.Placing != "" {
		pacing = c.Foundation.Old.Placing
	}
	if c.Foundation.NotOOSPlacing != "" {
		pacing = c.Foundation.NotOOSPlacing
	}
	if c.Foundation.Placing44 != "" {
		pacing = c.Foundation.Placing44
	}
	if c.Foundation.Placing94 != "" {
		pacing = c.Foundation.Placing94
	}
	if c.Foundation.OtherPlacing != "" {
		pacing = c.Foundation.OtherPlacing
	}
	if c.Foundation.Fcs.Order.Placing != "" {
		pacing = c.Foundation.Fcs.Order.Placing
	}

	switch pacing {
	case "1":
		return "открытый конкурс"
	case "2":
		return "конкурс с ограниченным участием"
	case "3":
		return "двухэтапный конкурс"
	case "4":
		return "закрытый конкурс"
	case "5":
		return "закрытый конкурс с ограниченным участием"
	case "6":
		return "закрытый двухэтапный конкурс"
	case "7":
		return "аукцион в электронной форме"
	case "8":
		return "закрытый аукцион"
	case "9":
		return "запрос котировок"
	case "10":
		return "запрос предложений"
	case "11":
		return "закупка у единственного поставщика (подрядчика, исполнителя)"
	case "11011":
		return "открытый конкурс"
	case "11021":
		return "конкурс с ограниченным участием"
	case "11031":
		return "двухэтапный конкурс"
	case "12011":
		return "электронный аукцион"
	case "13011":
		return "запрос котировок"
	case "14011":
		return "запрос предложений "
	case "20000":
		return "закупка у единственного поставщика (подрядчика, исполнителя)"
	case "30000":
		return "способ определения поставщика (подрядчика, исполнителя),  установленный Правительством Российской Федерации в соответствии со статьей 111 Федерального закона.</xs:documentation>"
	case "11042":
		return "закрытый конкурс"
	case "11052":
		return "закрытый конкурс с ограниченным участием"
	case "11062":
		return "закрытый двухэтапный конкурс"
	case "12022":
		return "закрытый аукцион"
	default:
		return ""
	}
	//log.Println(c.Foundation)
	return ""
}

func (c *Contract) GetBudgetSource() string {
	var lvl string
	if c.Finances.Level != "" {
		lvl = c.Finances.Level
	} else {
		if c.Finances.LevelOld != "" {
			lvl = c.Finances.LevelOld
		} else {
			if c.Finances.Source != "" {
				lvl = c.Finances.Source
			}
		}
	}
	switch lvl {
	case "11":
		return "федеральный бюджет"
	case "12":
		return "бюджет субъекта Российской Федерации"
	case "13":
		return "местный бюджет"
	case "14":
		return "бюджет Пенсионного фонда Российской Федерации"
	case "15":
		return "бюджет Федерального фонда обязательного медицинского страхования"
	case "16":
		return "бюджет Фонда социального страхования Российской Федерации"
	case "17":
		return "бюджет территориального государственного внебюджетного фонда."
	case "":
		return "внебюджетные средства"
	case "01":
		return "федеральный бюджет"
	case "02":
		return "бюджет субъекта Российской Федерации"
	case "03":
		return "местный бюджет"
	case "04":
		return "бюджет Пенсионного фонда Российской Федерации"
	case "05":
		return "бюджет Федерального фонда обязательного медицинского страхования"
	case "06":
		return "бюджет Фонда социального страхования Российской Федерации"
	case "07":
		return "бюджет территориального государственного внебюджетного фонда"
	default:
		return lvl
	}
}

func (c *Contract) GetExecDate() string {
	// Convert ExecDate to database format!
	if c.ExecDate.Month != "" && c.ExecDate.Year != "" {
		//log.Println("ExecDate(old):", c.ExecDate.Year+"-"+c.ExecDate.Month+"-01", ShortDate(c.ExecDate.Month, c.ExecDate.Year))
		return std.ShortDate(c.ExecDate.Month, c.ExecDate.Year)
	}
	if c.ExecDateEnd != "" {
		//		log.Println("ExecDate(new):", c.ExecDateEnd)
		return c.ExecDateEnd
	}
	//log.Println("? Fucking hell...")
	return ""
}

func (c *Contract) GetPrice() float64 {
	if c.PriceInfo.Currency != "" {
		return c.PriceInfo.Price
	} else {
		return c.Price
	}
}

func (c *Contract) GetCurrency() string {
	if c.PriceInfo.Currency != "" {
		return c.PriceInfo.Currency
	} else {
		return c.Currency.Code
	}
}

func (c *Contract) GetRate() float64 {
	if c.PriceInfo.Rate != 0 {
		return c.PriceInfo.Rate
	} else {
		return 1
	}
}

func (c *Contract) GetPublishDate() string {
	published := c.PublishDate
	if len(published) < 5 && published != "" {
		t := strings.Split(published, " ")
		return std.ShortDate(t[0], t[1])
	}
	return published
}

func (c *Contract) GetBudget() string {
	if c.Finances.Budget != "" {
		//		log.Println("Budget:", c.Finances.Budget)
		return c.Finances.Budget
	} else {
		if c.Finances.BudgetOld != "" {
			return c.Finances.BudgetOld
		}
		if c.Finances.ExtraBudgetOld != "" {
			return c.Finances.ExtraBudgetOld
		}
		if c.Finances.ExtraBudget != "" {
			//			log.Println("Budget:", c.Finances.ExtraBudget)
			return c.Finances.ExtraBudget
		}
		//		log.Println("FinanceSource:", c.Finances.Source)
		//return c.Finances.Source
		return ""
	}
}

func (c *Contract) GetOKTMO() string {
	return ""
}

func GetNtParams(pnum string, lot int, db *sql.DB) (ntname string, nmck float64) {

	if pnum == "" {
		return "", -1
	}

	if err := db.QueryRow("select objectInfo from notifications_"+shards.GetByPnum(pnum)+" where pnum=$1::bigint", pnum).Scan(&ntname); err != nil {
		log.Println(fmt.Errorf("Can't get Notifications (%s) params: %v\n", pnum, err))
	}

	if err := db.QueryRow("select maxPrice from notifications_lots_"+shards.GetByPnum(pnum)+" where pnum=$1::bigint and lotnum=$2", pnum, lot).Scan(&nmck); err != nil {
		log.Println(fmt.Errorf("Can't get Notifications (%s) params: %v\n", pnum, err))
	}
	if nmck == 0 {
		log.Println(fmt.Errorf("Null nmck for %s with lotnum: %d", pnum, lot))
	}
	return
}

func SignNt(pnum, regnum string, price float64, customerinn string, db *sql.DB) error {

	if pnum == "" {
		log.Errorf("Purchasenumber is not defined for R#%s", regnum)
		return fmt.Errorf("Purchasenumber is not defined for R#%s", regnum)
	}

	/*res, err := db.Exec("Update notifications_lots_"+shards.GetByPnum(pnum)+" set regnum=$1, ctprice=$2 where pnum=$3::bigint and lotnum=$4;", regnum, price, pnum, lot)
	log.Error(fmt.Sprintf("Update notifications_lots_"+shards.GetByPnum(pnum)+" set regnum='%s', ctprice=%f where pnum=%s and lotnum=%d", regnum, price, pnum, lot))
	n0, err := res.RowsAffected()
	if n0 == 0 || err != nil {
		return fmt.Errorf("No rows affected (lots) on signing of %s. Additional error: %v", regnum, err)
	}
	time.Sleep(1 * time.Millisecond)*/
	res, err := db.Exec(fmt.Sprintf("Update notifications_customers_"+shards.GetByPnum(pnum)+" set regnum='%s', ctprice=%f where pnum=%s and customer_inn='%s'", regnum, price, pnum, customerinn))
	log.Error(fmt.Sprintf("Update notifications_customers_"+shards.GetByPnum(pnum)+" set regnum='%s', ctprice=%f where pnum=%s and customer_inn='%s'", regnum, price, pnum, customerinn))
	if err != nil || res == nil {
		log.Fatalf("No rows affected (customers) on signing of %s. Additional error: %v", regnum, err)
		return fmt.Errorf("No rows affected (customers) on signing of %s. Additional error: %v", regnum, err)
	}
	n0, err := res.RowsAffected()
	if n0 == 0 || err != nil {
		log.Errorf("No rows affected (customers) on signing of %s. Additional error: %v", regnum, err)
		return fmt.Errorf("No rows affected (customers) on signing of %s. Additional error: %v", regnum, err)
	}
	//log.Fatal("Set: R#%s;Price:%f for N#%s", regnum, price, pnum)
	return nil
}

func (p Product) GetOKPD() string {
	//log.Infof("OKPD CHECK: OKPD:%s OKPD2:%s", p.OKPD, p.OKPD2)
	if p.OKPD == "" {
		if p.OKPD2 == "" {
			return strings.Split(p.KTRU, "-")[0]
		}
		return p.OKPD2
	} else {
		return p.OKPD
	}

}

func ApplyProcs(regnum string, db *sql.DB) error {
	log.Infoln("Applying Procs:", regnum)
	rows, err := db.Query(`
		SELECT id, published, stage, paid
		  FROM contracts_proc_` + shards.GetByPnum(regnum) + `
		 WHERE applied = false AND regnum = $1
		 ORDER BY id DESC
		 LIMIT 1;
	`, regnum)
	if err != nil {
		log.Errorf("Error in applying procs (#R %s): %v\n", regnum, err)
		return err
	}
	log.Infoln("Query !:", regnum)

	var idA []int
	var publishedA []time.Time
	var stageA []string
	var paidA []float64

	for rows.Next() {

		var id int
		var published time.Time
		var stage string
		var paid float64

		if err := rows.Scan(&id, &published, &stage, &paid); err != nil {
			log.Errorln("ERR ON APPLY PROCS: ", err)
			rows.Close()
			return err
		}
		idA = append(idA, id)
		publishedA = append(publishedA, published)
		stageA = append(stageA, stage)
		paidA = append(paidA, paid)
	}
	rows.Close()
	for i, _ := range idA {
		var n0 int64
		var id int = idA[i]
		var published time.Time = publishedA[i]
		var stage string = stageA[i]
		var paid float64 = paidA[i]

		log.Infoln("TT:", id, published, stage, paid)
		var cpaid float64
		var cstatus string
		err := db.QueryRow("select Status, Paid from contracts_"+shards.GetByRegnum(regnum)+" where regnum=$1", regnum).Scan(&cstatus, &cpaid)
		log.Infoln("dfadsf")
		if err == nil && paid == 0 && cpaid == 0 && cstatus == stage {
			log.Info("Nothing to update here. ", regnum)
		} else {
			log.Infoln("UPP", regnum,
				stage, cpaid+paid)

			res0, err := db.Exec(`Update contracts_`+shards.GetByRegnum(regnum)+`  set Status=$2, Paid=$3 where regnum=$1;`, regnum,
				stage, cpaid+paid)
			if err != nil {
				log.Fatal("Pew-pew baby, fatal error is here: ", err)
				return fmt.Errorf(`Update contracts_`+shards.GetByRegnum(regnum)+`
						set Status='%s', Paid=(select Paid from contracts_`+shards.GetByRegnum(regnum)+` where regnum='%s') + %f where regnum='%s'; With additional error: %v
						`, stage, regnum, paid+cpaid, regnum, err)
				/*res1, err := db.Exec(`Update products_`+shards.GetByRegnum(p.RegNum)+` set Status=$2, Paid=(select Paid from contracts_`+shards.GetByRegnum(p.RegNum)+` where regnum=$1) + $3 where regnum=$1;`, p.RegNum,
				p.Stage, p.GetPaid())*/
			}
			n0, _ = res0.RowsAffected()
			/*res0, err = db.Exec(fmt.Sprintf(`Update products_`+shards.GetByRegnum(regnum)+` set Status='%s', Paid=(select Paid from contracts_`+shards.GetByRegnum(regnum)+` where regnum='%s') + %f where regnum='%s';`, stage, regnum, paid+cpaid, regnum))

			n0, _ = res0.RowsAffected()
			if err != nil {
				return fmt.Errorf(fmt.Sprintf(`Update products_`+shards.GetByRegnum(regnum)+` set Status='%s', Paid=(select Paid from contracts_`+shards.GetByRegnum(regnum)+` where regnum='%s') + %f where regnum='%s'; With error: %v`, stage, regnum, paid+cpaid, regnum, err))
			}
			*/
			//res1.RowsAffected()
		}
		_, err = db.Exec(`update contracts_proc_`+shards.GetByRegnum(regnum)+` set applied=$3 where id=$1 and regnum=$2`,
			id, regnum, n0 > 0)

		return err

	}

	return nil
}

func (c *Contract) Validate() bool {
	var err error
	check := true

	if c.RegNum == "" {
		check = false
		err = errors.New("Contract(null) is not valid: null RegNum!")
		log.Println(err)
		return check
	}
	if c.PublishDate == "" {
		err = errors.New("Contract(" + c.RegNum + ") is not valid: null publishDate!")
		log.Println(err)
	} else {
		c.PublishDate = std.SqlDateTime(c.PublishDate)
	}
	if c.SignDate == "" {
		check = false
		err = errors.New("Contract(" + c.RegNum + ") is not valid: null signDate!")
		log.Println(err)
	}

	return check
}

func (c *Contract) Identify() {
	log.Println("Contract: " + c.RegNum + "; #" + c.GetPurchaseNumber())
}

func ParseContract(data []byte, db *sql.DB) error {
	var export Contract
	err := xml.Unmarshal(data, &export)
	if err != nil {
		//log.Println(err)
		return err
	}
	ch := export.Validate()
	if !ch {
		//log.Println("Not valid!")
		return errors.New("Not valid")
	} else {

		if export.RegNum == "1222513070016000246" {
			log.Error("!!! 1222513070016000246: ", export.Products)
		}
		export.Save(db)
		export.Identify()
	}
	//log.Println(export)
	return nil
}

func (p *Product) Validate() (bool, error) {
	check := true
	var err error
	if p.Name == "" {
		check = false
		err = errors.New("Product isn't valid: null Name!")
	}
	if p.Price == 0 {
		check = false
		err = errors.New("Product isn't valid: zero Price!")
		return check, err
	}
	if p.Quantity == 0 {
		err = errors.New("Product probably isn't valid: zero Quantity!")
		return check, err
	}
	if p.Sum == 0 {
		check = false
		err = errors.New("Product isn't valid: zero Sum!")
		return check, err
	}
	return check, err
}

// ContractExists - does contract record exist ? (Sharding by signed date)
func ContractExists(db *sql.DB, regnum, published string) bool {
	var c bool

	err := db.QueryRow(`
		SELECT EXISTS(
			SELECT *
			FROM contracts_` + shards.GetByRegnum(regnum) + `
			WHERE regnum = $1)
	`, regnum).Scan(&c)
	if err != nil {
		log.Println("ContractExists err: ", err)
	}

	return c
}

// ContractUp - is contract record up-to-date ? (Comparing by published date)
func ContractUp(db *sql.DB, regnum, published string, version int) (bool, int, string) {
	var v int
	var p string

	db.QueryRow(`
		SELECT version, published
		FROM contracts_` + shards.GetByRegnum(regnum) + `
		WHERE regnum = $1
			AND (published <= to_timestamp(translate($2, 'T', ' '), 'YYYY-MM-DD HH24:MI:SS') OR version >= $3)
	`, regnum, published, version).Scan(&v, &p)
	return p != "", v, p
}

func (c *Contract) productsAmountCheck(db *sql.DB) (res bool, err error) {
	err = errors.New("Null result")
	err = db.QueryRow(`
		SELECT (
			SELECT COUNT(*)
			FROM products_` + shards.GetByRegnum(c.RegNum) + `
			WHERE regnum = $1
		) = $2
	`, c.RegNum, len(c.Products)).Scan(&res)
	return
}

func (c *Contract) Save(db *sql.DB) error {
	if std.YearFromText(c.SignDate) < 2014 {
		return nil
	}
	exists := ContractExists(db, c.RegNum, c.PublishDate)

	//err := insertOrUpdateSupplier(db, c.ListSuppliersName(), c.ListSuppliersINN(), c.ListSuppliersKPP(), c.ListSuppliersAddress(), c.ListSuppliersPhone())

	//if err != nil {
	//	return err
	//}

	log.Println("Contract Save Info: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";EX:", exists)
	if exists {
		//uptodate, v0, p0 := ContractUp(db, c.RegNum, c.PublishDate, c.VersionNumber)
		uptodate, v, p := ContractUp(db, c.RegNum, c.PublishDate, c.VersionNumber)
		if v < c.VersionNumber {
			uptodate = false
		}
		log.Println("Contract Up? R# ", c.RegNum, "/", uptodate, "/", v, "/", p, " compare to ", c.GetPublishDate(), "/", c.VersionNumber)
		// uptodate=true
		// Проверим, в норме ли количество объектов? Если не проходит - апдейтим к бобрам. А как быть с ценой контракта? А все океюшки должно быть.
		checkAmount, err := c.productsAmountCheck(db)
		log.Println("Contract PROD-CHECK: R#", c.RegNum, "; CHR:", checkAmount, "; ERR:", err)
		if !checkAmount && err == nil {
			uptodate = false
		}
		if !uptodate {

			log.Println("Going to update (1) Contract R#", c.RegNum, " N#", c.GetPurchaseNumber(), " V#", c.VersionNumber, " PD#", c.PublishDate, " SD#", c.SignDate)
			err = updateProducts(db, c.RegNum, c.GetPurchaseNumber(), c.Number, c.PublishDate, c.SignDate, c.GetExecDate(), c.GetBudget(), c.GetBudgetSource(), c.Customer.Inn, c.Customer.Kpp, c.Customer.FullName, c.GetRegion(c.Customer.Kpp), c.ListSuppliersINN(), c.ListSuppliersKPP(), c.ListSuppliersName(), c.ListSuppliersAddress(), c.ListSuppliersPhone(), c.GetCurrency(), c.GetRate(), c.GetPrice(), c.Status, c.GetPlacing(), c.Products, c.VersionNumber)
			if err != nil {
				log.Println("Contract Save Error On Update: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
				return err
			}
			err = updateContract(db, c.RegNum, c.GetPurchaseNumber(), c.PublishDate, c.SignDate, c.GetExecDate(), c.GetBudget(), c.Finances.OKTMO, c.ListSuppliersINN(), c.Customer.Inn, c.Placer, c.GetCurrency(), c.Href, c.Status, c.GetPrice(), c.GetRate(), c.VersionNumber, c.GetLotNumber())
			if err != nil {
				log.Println("Contract Save Error On Update: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
				return err
			}
		}
	} else {

		// Insert
		singleCustomer := c.Foundation.Fcs.Order.SingleCustomer.Reason != "" || c.Foundation.Old.SingleCustomer != ""

		err := insertContract(db, c.RegNum, c.GetPurchaseNumber(), c.PublishDate, c.SignDate, c.GetExecDate(), c.GetBudget(), c.Finances.OKTMO, c.ListSuppliersINN(), c.Customer.Inn, c.Placer, c.GetCurrency(), c.Href, c.Status, c.GetPrice(), c.GetRate(), c.VersionNumber, c.GetLotNumber(), singleCustomer)
		if err != nil {
			log.Println("Contract Save Error On Insert: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";V:", c.VersionNumber, ";E:", err)
			uptodate, _, _ := ContractUp(db, c.RegNum, c.PublishDate, c.VersionNumber)
			if err != nil {
				log.Println("Contract Save Error: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
				return err
			}
			if !uptodate {
				log.Println("Going to update (2) Contract R#", c.RegNum, " N#", c.GetPurchaseNumber())
				err = updateProducts(db, c.RegNum, c.GetPurchaseNumber(), c.Number, c.PublishDate, c.SignDate, c.GetExecDate(), c.GetBudget(), c.GetBudgetSource(), c.Customer.Inn, c.Customer.Kpp, c.Customer.FullName, c.GetRegion(c.Customer.Kpp), c.ListSuppliersINN(), c.ListSuppliersKPP(), c.ListSuppliersName(), c.ListSuppliersAddress(), c.ListSuppliersPhone(), c.GetCurrency(), c.GetRate(), c.GetPrice(), c.Status, c.GetPlacing(), c.Products, c.VersionNumber)
				if err != nil {
					log.Println("Contract Save Error On Update: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
					return err
				}
				err = updateContract(db, c.RegNum, c.GetPurchaseNumber(), c.PublishDate, c.SignDate, c.GetExecDate(), c.GetBudget(), c.Finances.OKTMO, c.ListSuppliersINN(), c.Customer.Inn, c.Placer, c.GetCurrency(), c.Href, c.Status, c.GetPrice(), c.GetRate(), c.VersionNumber, c.GetLotNumber())
				if err != nil {
					log.Println("Contract Save Error On Update: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
					return err
				}
			}
			return err
		}
		err = InsertProducts(db, c.RegNum, c.GetPurchaseNumber(), c.Number, c.PublishDate, c.SignDate, c.GetExecDate(), c.GetBudget(), c.GetBudgetSource(), c.Customer.Inn, c.Customer.Kpp, c.Customer.FullName, c.GetRegion(c.Customer.Kpp), c.ListSuppliersINN(), c.ListSuppliersKPP(), c.ListSuppliersName(), c.ListSuppliersAddress(), c.ListSuppliersPhone(), c.GetCurrency(), c.GetRate(), c.GetPrice(), c.Status, c.GetPlacing(), c.VersionNumber, c.Products)
		log.Info("Products insert for ", c.RegNum, c.Products)
		if err != nil {
			log.Println("Contract Save Error On InsertProducts: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
			return err
		}
		err = ApplyProcs(c.RegNum, db)
		if err != nil {
			log.Println("Contract Save Error On InsertProducts: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
			return err
		}
		err = SignNt(c.GetPurchaseNumber(), c.RegNum, c.GetPrice(), c.Customer.Inn, db)
		if err != nil {
			log.Println("Contract Save Error On Update: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
			return err
		}
		err = insertProductFiles(db, c.RegNum, c.SignDate, c.Attachments)
		if err != nil {
			log.Println("Contract Save Error On InsertFiles: R#", c.RegNum, "; N#", c.GetPurchaseNumber(), ";E:", err)
			return err
		}
		return err
	}
	return nil
}

func insertContract(db *sql.DB, regnum, purchaseNumber, published, signed, exec, budget, oktmo, suppliers, customer, placer, currency, link, status string, price, rate float64, version, lot int, singleCustomer bool) error {
	var ntname string
	var nmck float64

	if singleCustomer {
		log.Println("SingleCustomer: ", regnum)
		if purchaseNumber != "" {
			ntname, nmck = GetNtParams(purchaseNumber, lot, db)
		}
	} else {
		ntname, nmck = GetNtParams(purchaseNumber, lot, db)
	}
	_, err := db.Exec(`INSERT INTO contracts_`+shards.GetByRegnum(regnum)+` (RegNum , PurchaseNumber, Published, Signed, Exec, Budget, OKTMO, Suppliers, Customer, Placer, Price, Currency, Rate, Link, Status, Paid, Updated, Version, Lot, ntname, nmck)	VALUES($1, $2, to_timestamp(translate($3, 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'), $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, to_timestamp(translate($3, 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'), $17, $18, $19 ,$20) RETURNING row;`, regnum,
		purchaseNumber,
		published,
		signed,
		exec,
		budget,
		oktmo,
		suppliers,
		customer,
		placer,
		price,
		currency,
		rate,
		link,
		status,
		0., // Paid
		version,
		lot,
		ntname,
		nmck,
	)
	return err
}

func updateContract(db *sql.DB, regnum, purchaseNumber, published, signed, exec, budget, oktmo, suppliers, customer, placer, currency, link, status string, price, rate float64, version, lot int) error {
	_, err := db.Exec(fmt.Sprintf(`UPDATE contracts_`+shards.GetByRegnum(regnum)+` SET RegNum='%s',
		PurchaseNumber='%s',
		Published=to_timestamp(translate('%s', 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'),
		Signed='%s',
		Exec='%s',
		Budget='%s',
		OKTMO='%s',
		Suppliers='%s',
		Customer='%s',
		Placer='%s',
		Price='%f',
		Currency='%s',
		Rate='%f',
		Link='%s',
		Status='%s',
		Updated=to_timestamp(translate('%s', 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'),
		version='%d', lot='%d' WHERE regnum='%s' RETURNING row;`,
		regnum,
		purchaseNumber,
		published,
		signed,
		exec,
		budget,
		oktmo,
		suppliers,
		customer,
		placer,
		price,
		currency,
		rate,
		link,
		status,
		//regnum, // Paid
		published,
		version,
		lot,
		regnum))
	if err != nil {
		log.Println(fmt.Sprintf(`UPDATE contracts_`+shards.GetByRegnum(regnum)+` SET RegNum='%s',
			PurchaseNumber='%s',
			Published=to_timestamp(translate('%s', 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'),
			Signed='%s',
			Exec='%s',
			Budget='%s',
			OKTMO='%s',
			Suppliers='%s',
			Customer='%s',
			Placer='%s',
			Price='%f',
			Currency='%s',
			Rate='%f',
			Link='%s',
			Status='%s',
			Updated=to_timestamp(translate('%s', 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'),
			version='%d', lot='%d' WHERE regnum='%s' RETURNING row;`,
			regnum,
			purchaseNumber,
			published,
			signed,
			exec,
			budget,
			oktmo,
			suppliers,
			customer,
			placer,
			price,
			currency,
			rate,
			link,
			status,
			//regnum, // Paid
			published,
			version,
			lot,
			regnum))
		return err
	}
	err = ApplyProcs(regnum, db)
	return err
}

func InsertProducts(db *sql.DB, regnum, purchaseNumber, number, published, signed, exec, budget, budgetSource, customerINN, customerKPP, customerName string, customerRegion int, supplierINN, supplierKPP, supplierName, supplierAddress, supplierPhone, currency string, rate, contractPrice float64, status, placing string, version int, products []Product) error {
	log.Println(regnum, ":", customerRegion)
	// Checking for previous version
	var preVersion sql.NullInt64
	err := db.QueryRow("select max(version) from products_"+shards.GetByRegnum(regnum)+" where regnum=$1", regnum).Scan(&preVersion)
	if err != nil {
		return err
	}
	log.Println(preVersion.Valid, int(preVersion.Int64), version)
	// if preVersion.Valid && int(preVersion.Int64) > version {
	if (int(preVersion.Int64) > version) || (int(preVersion.Int64) == 0) {

		log.Println("inserting...")
		for _, p := range products {

			_, err := db.Exec(`INSERT INTO products_`+shards.GetByRegnum(regnum)+`(RegNum,PurchaseNumber,Name,OKPD,OKPDInfo,Units,Quantity,Price,PriceRu,Sum,SumRu,Rate,Published,Signed,Exec,Budget,BudgetSource,CustomerINN,CustomerKPP,Customer,CustomerRegion,SupplierINN,SupplierKPP,Supplier,SupplierAddress,SupplierPhone,ContractPrice,Currency,Paid,Status,PlacingWay, Number, Updated, Version, Sid, Ktru) VALUES($1,$2,$3,$4,(SELECT nsiokpd.name FROM nsiokpd WHERE nsiokpd.code = $4 and actual = TRUE),$5,$6,$7,$8,$9,$10,$11,to_timestamp(translate($12,'T',' '),'YYYY-MM-DD HH24:MI:SS'),to_timestamp(translate($13,'T',' '),'YYYY-MM-DD HH24:MI:SS'),to_timestamp(translate($14,'T',' '),'YYYY-MM-DD HH24:MI:SS'),$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,to_timestamp(translate($12,'T',' '),'YYYY-MM-DD HH24:MI:SS'),$32,$33,$34) RETURNING row;`, regnum,
				purchaseNumber,
				p.Name,
				p.GetOKPD(),
				p.OKEI,
				p.Quantity,
				p.Price,
				p.PriceRUR,
				p.Sum,
				p.SumRUR,
				rate,
				published,
				signed,
				exec,
				budget,
				budgetSource,
				customerINN,
				customerKPP,
				customerName,
				customerRegion,
				supplierINN,
				supplierKPP,
				supplierName,
				supplierAddress,
				supplierPhone,
				contractPrice,
				currency,
				0.,
				status,
				placing,
				number,
				version,
				p.Sid,
				p.KTRU,
			)
			if err != nil {
				log.Println("sid:", p.Sid, ";  KTRU:", p.KTRU)
				log.Error(err, fmt.Sprintf(`INSERT INTO products_`+shards.GetByRegnum(regnum)+`(RegNum,PurchaseNumber,Name,OKPD,OKPDInfo,Units,Quantity,Price,PriceRu,Sum,SumRu,Rate,Published,Signed,Exec,Budget,CustomerINN,Customer,CustomerRegion,SupplierINN,Supplier,ContractPrice,Currency,Paid,Status) VALUES('%s','%s','%s','%s',(SELECT nsiokpd.name FROM nsiokpd WHERE nsiokpd.code = '%s' and actual = TRUE),(SELECT localname FROM nsiokei WHERE code = '%s' and actual = TRUE),%f,%f,%f,%f,%f,%f,to_timestamp(translate('%s','T',' '),'YYYY-MM-DD HH24:MI:SS'),to_timestamp(translate('%s','T',' '),'YYYY-MM-DD HH24:MI:SS'),to_timestamp(translate('%s','T',' '),'YYYY-MM-DD HH24:MI:SS'),'%s','%s',(SELECT fullname FROM nsiOrg where inn = '%s' AND actual = TRUE),(SELECT nsioktmo.fullname from nsioktmo where nsioktmo.code in (SELECT nsiOrg.oktmo FROM nsiOrg where nsiOrg.inn = '%s' AND actual = TRUE) AND actual = TRUE),'%s','%s',%f,'%s',%f,'%s') RETURNING row;`, regnum, purchaseNumber, p.Name, p.GetOKPD(), p.GetOKPD(), p.OKEI, p.Quantity, p.Price, p.PriceRUR, p.Sum, p.SumRUR, rate, published, signed, exec, budget, customerINN, customerINN, customerINN, supplierINN, supplierName, contractPrice, currency, 0., status))
				return err
			}
		}
		return err
	} else {
		log.Println("Version is older than existed!", regnum)
	}
	return err
}

func updateProducts(db *sql.DB, regnum, purchaseNumber, number, published, signed, exec, budget, budgetSource, customerINN, customerKPP, customerName string, customerRegion int, supplierINN, supplierKPP, supplierName, supplierAddress, supplierPhone, currency string, rate, contractPrice float64, status, placing string, products []Product, version int) error {
	_, err := db.Exec(`DELETE FROM products_`+shards.GetByRegnum(regnum)+` WHERE regnum=$1;`, regnum)
	if err != nil {
		return err
	}
	for _, p := range products {
		_, err := db.Exec(`INSERT INTO products_`+shards.GetByRegnum(regnum)+`(RegNum,PurchaseNumber,Name,OKPD,OKPDInfo,Units,Quantity,Price,PriceRu,Sum,SumRu,Rate,Published,Signed,Exec,Budget,BudgetSource,CustomerINN,CustomerKPP,Customer,CustomerRegion,SupplierINN,SupplierKPP,Supplier,SupplierAddress,SupplierPhone,ContractPrice,Currency,Paid,Status,PlacingWay, Number, Updated, Version, Sid, Ktru) VALUES($1,$2,$3,$4,(SELECT nsiokpd.name FROM nsiokpd WHERE nsiokpd.code = $4 and actual = TRUE),$5,$6,$7,$8,$9,$10,$11,to_timestamp(translate($12,'T',' '),'YYYY-MM-DD HH24:MI:SS'),to_timestamp(translate($13,'T',' '),'YYYY-MM-DD HH24:MI:SS'),to_timestamp(translate($14,'T',' '),'YYYY-MM-DD HH24:MI:SS'),$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,to_timestamp(translate($12,'T',' '),'YYYY-MM-DD HH24:MI:SS'),$32,$33,$34) RETURNING row;`, regnum,
			purchaseNumber,
			p.Name,
			p.GetOKPD(),
			p.OKEI,
			p.Quantity,
			p.Price,
			p.PriceRUR,
			p.Sum,
			p.SumRUR,
			rate,
			published,
			signed,
			exec,
			budget,
			budgetSource,
			customerINN,
			customerKPP,
			customerName,
			customerRegion,
			supplierINN,
			supplierKPP,
			supplierName,
			supplierAddress,
			supplierPhone,
			contractPrice,
			currency,
			0.,
			status,
			placing,
			number,
			version,
			p.Sid,
			p.KTRU,)
		if err != nil {
			log.Println("sid:", p.Sid, ";  KTRU:", p.KTRU)
			log.Error(err, fmt.Sprintf(`INSERT INTO products_`+shards.GetByRegnum(regnum)+`(RegNum,PurchaseNumber,Name,OKPD,OKPDInfo,Units,Quantity,Price,PriceRu,Sum,SumRu,Rate,Published,Signed,Exec,Budget,CustomerINN,Customer,CustomerRegion,SupplierINN,Supplier,ContractPrice,Currency,Paid,Status) VALUES('%s','%s','%s','%s',(SELECT nsiokpd.name FROM nsiokpd WHERE nsiokpd.code = '%s' and actual = TRUE),(SELECT localname FROM nsiokei WHERE code = '%s' and actual = TRUE),%f,%f,%f,%f,%f,%f,to_timestamp(translate('%s','T',' '),'YYYY-MM-DD HH24:MI:SS'),to_timestamp(translate('%s','T',' '),'YYYY-MM-DD HH24:MI:SS'),to_timestamp(translate('%s','T',' '),'YYYY-MM-DD HH24:MI:SS'),'%s','%s',(SELECT fullname FROM nsiOrg where inn = '%s' AND actual = TRUE),(SELECT nsioktmo.fullname from nsioktmo where nsioktmo.code in (SELECT nsiOrg.oktmo FROM nsiOrg where nsiOrg.inn = '%s' AND actual = TRUE) AND actual = TRUE),'%s','%s',%f,'%s',%f,'%s') RETURNING row;`, regnum, purchaseNumber, p.Name, p.GetOKPD(), p.GetOKPD(), p.OKEI, p.Quantity, p.Price, p.PriceRUR, p.Sum, p.SumRUR, rate, published, signed, exec, budget, customerINN, customerINN, customerINN, supplierINN, supplierName, contractPrice, currency, 0., status))
			return err
		}
	}

	return err
}

func insertOrUpdateSupplier(db *sql.DB, name, inn, kpp, address, phone string) error {

	var supplierExist int

	err := db.QueryRow(`
		SELECT COUNT(*)
		FROM nsisuppliers
		WHERE inn = $1
		  AND kpp = $2
		LIMIT 1
	`, inn, kpp).Scan(&supplierExist)

	if err != nil {
		return fmt.Errorf("Fails on `insertOrUpdateSupplier` with INN %v and KPP %v, error: %v", err)
	}

	if supplierExist > 0 {

		_, err = db.Exec(`
			UPDATE nsisuppliers SET name = $1, address = $2, phone = $3
			WHERE inn = $4
			  AND kpp = $5
		`, name, address, phone, inn, kpp)

	} else {

		_, err = db.Exec(`
			INSERT INTO nsisuppliers (name, inn, kpp, address, phone)
			VALUES ($1, $2, $3, $4, $5)
		`, name, inn, kpp, address, phone)
	}

	if err != nil {
		return fmt.Errorf("Fails on `insertOrUpdateSupplier` with INN %v and KPP %v, error: %v", err)
	}

	return nil
}

func insertProductFiles(db *sql.DB, regnum, signed string, attachments []std.Attachment) error {
	var count int
	var err error
	for i := range attachments {
		_, err = db.Exec("INSERT INTO contracts_files_"+shards.GetByRegnum(regnum)+"(regnum, description, name, url, size, ignore,downloaded, published)	VALUES($1, $2, $3, $4, $5, $6,$7, to_timestamp(translate($8, 'T', ' '), 'YYYY-MM-DD HH24:MI:SS')) RETURNING row;", regnum,
			attachments[i].DocDescription,
			attachments[i].FileName,
			attachments[i].Url,
			attachments[i].FileSize,
			false,
			false,
			signed,
		)
		if err != nil {
			return err
		}
		count = count + 1
	}
	return err
}
