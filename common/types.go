package common

//import "encoding/xml"
import (
	"bytes"
	"database/sql"
	"time"

	"bitbucket.org/crmsib/parser_gov/shards"
	//"encoding/json"

	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

type Org struct {
	RegNum   string `xml:"regNum" json:"RegNum"`
	FullName string `xml:"fullName" json:"FullName"`
	INN      string `xml:"INN" json:"INN"`
	KPP      string `xml:"KPP" json:"KPP"`
}

type KladrPlace struct {
	Kladr struct {
		Type string `xml:"kladr>kladrType"`
		Code string `xml:"kladr>kladrCode"`
		Name string `xml:"kladr>fullName"`
	} `xml:"kladr"`
	Country struct {
		Code string `xml:"countryCode"`
		Name string `xml:"countryFullName"`
	} `xml:"country"`
	DeliveryPlace string `xml:"deliveryPlace"`
	NoKladr       struct {
		Region     string `xml:"region"`
		Settlement string `xml:"settlement"`
	} `xml:"noKladrForRegionSettlement"`
}

type PurchaseObject struct {
	OKPD                OKPDref                  `xml:"OKPD"`
	Name                string                   `xml:"name"`
	OKEI                OKEIref                  `xml:"OKEI"`
	QuantityPerCustomer []QuantityPerCustomerRef `xml:"customerQuantities>customerQuantity"`
	Quantity            struct {
		Value     float64 `xml:"quantity"`
		Countable bool    `xml:"undefined"`
	} `xml:"quantity"`
	Price float64 `xml:"price"` // Стоимость за единицу
	Sum   float64 `xml:"sum"`   // Стоимость по позиции
}

type OKPDref struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}

type OKEIref struct {
	Code  string `xml:"code"`
	NCode string `xml:"nationalCode"`
}

type OKOPFref struct {
	Code string `xml:"code"`
	Name string `xml:"fullName"`
}

type QuantityPerCustomerRef struct {
	Customer Org     `xml:"customer" json:"customer"`
	Value    float64 `xml:"quantity" json:"quantity"`
}

type Attachment struct {
	Url                string `xml:"url"`
	DocDescription     string `xml:"docDescription"`
	FileName           string `xml:"fileName"`
	FileSize           string `xml:"fileSize"`
	PublishedContentId int    `xml:"publishedContentID"`
}

type ETPRef struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
	Url  string `xml:"url"`
}

type PurchaseResponsible struct {
	Company Org     `xml:"responsibleOrg"`
	Info    OrgInfo `xml:"responsibleInfo"`
	Role    string  `xml:"responsibleRole"`
	Special Org     `xml:"specializedOrg"`
	/*
		Роль организации, осуществляющей закупку
		CU - Заказчик;
		OCU - Заказчик в качестве организатора совместного аукциона;
		RA - Уполномоченный орган;
		ORA- Уполномоченный орган в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ;
		AI - Уполномоченное учреждение;
		OAI- Уполномоченное учреждение в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ;
		OA - Организация, осуществляющая полномочия заказчика на осуществление закупок на основании договора (соглашения);
		OOA- Организация, осуществляющая полномочия заказчика на осуществление закупок на основании договора (соглашения) в качестве организатора совместного конкурса (аукциона) согласно ст. 25 №44ФЗ</xs:documentation>
	*/

} //purchaseResponsible

type OrgInfo struct {
	First  string `xml:"contactPerson>firstName"`
	Last   string `xml:"contactPerson>lastName"`
	Middle string `xml:"contactPerson>middleName"`
	Email  string `xml:"contactEMail"`
	Phone  string `xml:"contractPhone"`
	More   string `xml:"addInfo"`
}

type AccountRef struct {
	BankAddress     string `xml:"bankAddress"`
	BankName        string `xml:"bankName"`
	BIK             string `xml:"bik"`
	CorrAccount     string `xml:"corrAccount"`
	PaymentAccount  string `xml:"paymentAccount"`
	PersonalAccount string `xml:"personalAccount"`
}

type BudgetRef struct {
	Code string `xml:"code"`
	Name string `xml:"name"`
}

type ContactPersonRef struct {
	First  string `xml:"firstName"`
	Last   string `xml:"lastName"`
	Middle string `xml:"middleName"`
}

func MonthFromText(pub string) int {
	t := strings.Split(pub, "-")
	m, _ := strconv.Atoi(t[1])
	return m
}

func YearFromText(pub string) int {
	t := strings.Split(pub, "-")
	y, _ := strconv.Atoi(t[0])
	return y
}

func NotificationExistsM(db *sql.DB, purchaseNumber, shard string) bool {
	var c int
	//log.Println("notifications_" + shard)
	err := db.QueryRow("select count(*) from notifications_"+shard+" where purchaseNumber = $1;", purchaseNumber).Scan(&c)
	if err != nil {
		log.Println("Fucked: ", err, shard, purchaseNumber)
	}
	//log.Println(c)
	if c == 0 {
		return false
	} else {
		return true
	}
}

func SignContract(db *sql.DB, purchaseNumber, regnum string) error {
	check := NotificationExistsM(db, purchaseNumber, shards.GetByPnum(purchaseNumber))
	if check {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()
		stmt, err := tx.Prepare("UPDATE notifications_" + shards.GetByPnum(purchaseNumber) + " SET regnum=$2 WHERE purchaseNumber=$1")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(purchaseNumber, regnum)
		stmt.Close()
		if err != nil {
			return err
		}
		err = tx.Commit()
		return err
	}

	return fmt.Errorf("Notifications with PurchaseNumber %s wasnt found! Can't be signed!", purchaseNumber)
}

func BytesToUtf8(text []byte) []byte {
	sr := bytes.NewReader(text)
	tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	if err != nil {
		log.Fatal("BytesToUtf8: ", err)
	}
	return buf
}
func TextToUtf8(text string) string {
	sr := strings.NewReader(text)
	tr := transform.NewReader(sr, charmap.KOI8R.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	output := string(buf)
	output = strings.Replace(output, string(0x00), "", -1)

	if err != nil {
		return ""
		log.Println("TextToUtf8: ", err)
	}

	return output
}
func SqlDateTime(d string) string {
	// '0000-00-00' or '0000-00-00 00:00:00'
	Moscow, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatal("Time error:", err)
	}
	t, _ := time.Parse(time.RFC3339, d)
	return t.In(Moscow).String()[:19]
}

func ShortDate(m string, y string) string {
	var Month, Year string
	if len(string(m)) == 1 {
		Month = "0" + string(m)
	} else {
		Month = string(m)
	}
	Year = string(y)
	return Year + "-" + Month + "-01"
}
