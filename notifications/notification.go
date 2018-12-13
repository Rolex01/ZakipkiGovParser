package notifications

import (
	"database/sql"
	"encoding/xml"
	"errors"
	//"log"
	log "github.com/Sirupsen/logrus"

	//std "bitbucket.org/crmsib/parser_gov/common"
	_ "github.com/lib/pq"
	"strings"
	"fmt"
)


// Информация о порядке проведения - неинтересно. 24.03.2015 - единогласно.
type NotificationExport struct {
	XMLName				xml.Name		`xml:"export"`
	NotificationEF		Notification	`xml:"fcsNotificationEF"`		// NotificationEF - Извещение о проведении ЭА (электронный аукцион); внесение изменений
	NotificationEP		Notification	`xml:"fcsNotificationEP"`		// NotificationEP - Извещение о проведении закупки у ЕП (единственного поставщика), внесение изменений
	NotificationOKD		Notification	`xml:"fcsNotificationOKD"`		// NotificationOKD - Извещение о проведении OK-Д (двухэтапный конкурс), внесение изменений
	NotificationOKOU	Notification	`xml:"fcsNotificationOKOU"`		// NotificationOKOU - Извещение о проведении OK-ОУ (конкурс с ограниченным участием), внесение изменений
	NotificationOK		Notification	`xml:"fcsNotificationOK"`		// NotificationOK - Извещение о проведении OK (открытый конкурс), внесение изменений
	NotificationPOT		Notification	`xml:"fcsNotificationPOT"`		// NotificationPOT -
	NotificationZK		Notification	`xml:"fcsNotificationZK"`		// NotificationZK - Извещение о проведении ЗK (запрос котировок), внесение изменений
	NotificationZP		Notification	`xml:"fcsNotificationZP"`		// NotificationZP - Извещение о проведении ЗП (запрос предложений), внесение изменений
	NotificationPO		Notification	`xml:"fcsNotificationPO"`		// NotificationPO - Извещение о проведении ПО (предварительный отбор), внесение изменений
	NotificationZKB		Notification	`xml:"fcsNotificationZKB44"`	// NotificationZKB -
	NotificationISM		Notification	`xml:"fcsNotificationISM"`		// NotificationISM -
	NotificationZakA	Notification	`xml:"fcsNotificationZakA"`		// NotificationZakA - Извещение о проведении ЗакА (закрытый аукцион), внесение изменений
	NotificationZakKOU	Notification	`xml:"fcsNotificationZakKOU"`	// NotificationZakKOU - Извещение о проведении ЗакK-ОУ (закрытый конкурс с ограниченным участием), внесение изменений
	NotificationZakK	Notification	`xml:"fcsNotificationZakK"`		// NotificationZakK - Извещение о проведении ЗакK (закрытый конкурс), внесение изменений
	NotificationZakKD	Notification	`xml:"fcsNotificationZakKD"`	// NotificationZakKD - Извещение о проведении ЗакK-Д (закрытый двухэтапный конкурс), внесение изменений
	// Notification111 - Извещение о проведении закупки способом определения поставщика (подрядчика, исполнителя),
	// установленным Правительством Российской Федерации в соответствии со ст. 111 Федерального закона № 44-ФЗ;
	// внесение изменений
	Notification111		Notification	`xml:"fcsNotification111"`
	NotificationZKBI	Notification	`xml:"fcsProtocolZKBI"`			// NotificationZKBI -
}

type Notification struct {
	ID					int64				`xml:"id"`
	ExternalID			string				`xml:"externalId"`
	PurchaseNumber		string				`xml:"purchaseNumber"`
	PurchaseObjectInfo	string				`xml:"purchaseObjectInfo"`
	Href				string				`xml:"href"`
	Org_regnum			string				`xml:"purchaseResponsible>responsibleOrg>regNum"`
	Org_fullName		string				`xml:"purchaseResponsible>responsibleOrg>fullName"`
	Org_inn				string				`xml:"purchaseResponsible>responsibleOrg>INN"`
	Org_kpp				string				`xml:"purchaseResponsible>responsibleOrg>KPP"`
	Org_role			string				`xml:"purchaseResponsible>responsibleRole"`
	Placingway_code		string				`xml:"placingWay>code"`
	DocPublishDate		string				`xml:"docPublishDate"`
	ETP_code			string				`xml:"ETP>code"`
	Lot					NotificationLot		`xml:"lot"`
	Lots				[]NotificationLot	`xml:"lots>lot"`
	Attachments			[]Attachment        `xml:"attachments>attachment"`
	Other				string				`xml:"void"`
}

type Attachment struct {
	PublishedContentId	string	`xml:"publishedContentID"`
	FileName			string	`xml:"fileName"`
	FileSize			string	`xml:"fileSize"`
	DocDescription		string	`xml:"docDescription"`
	DocDate				string	`xml:"docDate"`
	Url					string	`xml:"url"`
	Content				string	`xml:"content"`
}

type NotificationLot struct {
	LotNumber				int64			`xml:"lotNumber"`
	LotObjectInfo			string			`xml:"lotObjectInfo"`
	MaxPrice				float64			`xml:"maxPrice"`
	Currency				string			`xml:"currency>code"`
	QuantityUndefined		bool			`xml:"quantityUndefined"`
	FinanceSource			string			`xml:"financeSource"`
	Preferenses				[]Condition		`xml:"preferenses>preferense"`
	Requirements			[]Condition		`xml:"requirements>requirement"`
	Restrictions			[]Condition		`xml:"restrictions>restriction"`
	RestrictInfo			string			`xml:"restrictInfo"`
	AddInfo					string			`xml:"addInfo"`
	TotalSumPO				float64			`xml:"purchaseObjects>totalSum"`
	TotalSumDPO				float64			`xml:"drugPurchaseObjectsInfo>total"`

	PurchaseObjects			[]PurchaseObject		`xml:"purchaseObjects>purchaseObject"`
	DrugPurchaseObject		[]DrugPurchaseObject	`xml:"drugPurchaseObjectsInfo>drugPurchaseObjectInfo"`
	CustomerRequirements	[]СustomerRequirement	`xml:"customerRequirements>customerRequirement"`
	// Preferenses, requirements and restrictions were missed here
}

type PurchaseObject struct {
	OKPD_code			string					`xml:"OKPD>code"`
	OKPD2_code			string					`xml:"OKPD2>code"`
	KTRU_code			string					`xml:"KTRU>code"`
	Name				string					`xml:"name"`
	OKEI				string					`xml:"OKEI>code"`
	CustomerQuantitys	[]CustomerQuantity		`xml:"customerQuantities>customerQuantity"`
	Quantity			float64					`xml:"quantity>value"`
	Countable			bool					`xml:"quantity>undefined"`
	Price				float64					`xml:"price"` // Стоимость за единицу
	Sum					float64					`xml:"sum"`   // Стоимость по позиции
}

type DrugPurchaseObject struct {
	ExternalCode		string				`xml:"objectInfoUsingReferenceInfo>drugsInfo>drugInfo>MNNInfo>MNNExternalCode"`
	NameR				string				`xml:"objectInfoUsingReferenceInfo>drugsInfo>drugInfo>MNNInfo>MNNName"`
	BasicUnitR			bool				`xml:"objectInfoUsingReferenceInfo>drugsInfo>drugInfo>basicUnit"`
	GRLSValueR			string				`xml:"objectInfoUsingReferenceInfo>drugsInfo>drugInfo>dosageInfo>dosageGRLSValue"`
	UserOKEI_codeR		string				`xml:"objectInfoUsingReferenceInfo>drugsInfo>drugInfo>dosageInfo>dosageUserOKEI>code"`
	// DrugQuantity		float64				`xml:"objectInfoUsingReferenceInfo>drugsInfo>drugInfo>drugQuantity"`
	NameT				string				`xml:"objectInfoUsingTextForm>drugsInfo>drugInfo>MNNInfo>MNNName"`
	BasicUnitT			bool				`xml:"objectInfoUsingTextForm>drugsInfo>drugInfo>basicUnit"`
	GRLSValueT			string				`xml:"objectInfoUsingTextForm>drugsInfo>drugInfo>dosageInfo>dosageGRLSValue"`
	UserOKEI_codeT		string				`xml:"objectInfoUsingTextForm>drugsInfo>drugInfo>manualUserOKEI>code"`


	IsZNVLP				bool				`xml:"isZNVLP"`
	CustomerQuantitys	[]CustomerQuantity	`xml:"drugQuantityCustomersInfo>drugQuantityCustomerInfo"`
	Quantity			float64				`xml:"drugQuantityCustomersInfo>total"`
	OKPD2_code			string				`xml:"OKPD2>code"`
	Price				float64				`xml:"pricePerUnit"`	// Стоимость за единицу
	Sum					float64				`xml:"positionPrice"`   // Стоимость по позиции
}

type CustomerQuantity struct {
	Customer_regnum	string		`xml:"customer>regNum" json:"customer>regNum"`
	Quantity		float64		`xml:"quantity" json:"quantity"`
}

type СustomerRequirement struct {
	Customer_regnum		string	`xml:"customer>regNum" json:"Customer>regNum"`
	Customer_fullname	string	`xml:"customer>fullName" json:"Customer>fullName"`
	MaxPrice			float64	`xml:"maxPrice" json:"MaxPrice"`
	DeliveryPlace		string	`xml:"deliveryPlace"`
}

type Condition struct {
	Code		string	`xml:"code"`
	Name		string	`xml:"name"`
	ShortName	string	`xml:"shortName"`
	PrefValue	string	`xml:"prefValue"`
	Content		string	`xml:"content"`
}

func ESQL(str *string) {
	*str = strings.Replace(*str, "'", "''", -1)
	return
}

func (p *Notification) InsertNotification(db *sql.DB, maxPrice float64, canceled, finished bool) error {
	var c int
	var err error

	ESQL(&p.PurchaseNumber)
	ESQL(&p.PurchaseObjectInfo)
	ESQL(&p.Href)
	ESQL(&p.Placingway_code)
	ESQL(&p.DocPublishDate)
	ESQL(&p.Org_regnum)
	ESQL(&p.Org_fullName)
	ESQL(&p.Org_inn)
	ESQL(&p.Org_kpp)
	ESQL(&p.Org_role)
	ESQL(&p.Other)

	file_year := strings.Split(p.DocPublishDate, "-")[0]

	if p.DocPublishDate == "" {
		err = errors.New("purchaseNumber has null publish date! " + p.PurchaseNumber)
		return err
	} else {
		p.DocPublishDate = fmt.Sprintf(`'%s'`, p.DocPublishDate)
	}

	// Аукционы
	if db.QueryRow("SELECT COUNT(row) FROM notifications WHERE pnum = $1;", p.PurchaseNumber).Scan(&c); c == 0 {
		my_query := fmt.Sprint(`
			WITH cte AS (
				SELECT fullname, inn, kpp, SUBSTRING(kpp, 1, 2)::int AS region
				FROM nsiorg
				WHERE regnum = '`, p.Org_regnum, `'
			)
			INSERT INTO notifications AS n (
				pnum, purchaseNumber, purchaseObjectInfo, href,
				org_regnum, org_fullname, org_inn, org_kpp, org_region, org_role,
				placingway_code, published, updated, ETP_code,
				maxprice, id
			)
			SELECT `, p.PurchaseNumber, `, '`, p.PurchaseNumber, `', '`, p.PurchaseObjectInfo, `', '`, p.Href, `',
				'`, p.Org_regnum, `', fullname, inn, kpp, region, '`, p.Org_role, `',
				'`, p.Placingway_code, `', `, p.DocPublishDate, `::timestamp, `, p.DocPublishDate, `::timestamp, '`, p.ETP_code, `',
				`, maxPrice, `, `, p.ID, `
			FROM (SELECT null AS void) t
				LEFT JOIN cte  ON 1 = 1
		`)

		_, err = db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Insert Notification:", my_query)
		} else {
			log.Println(`Insert Notification PurchaseNumber: "` + p.PurchaseNumber + `"`)
		}
	} else {
		my_query := fmt.Sprint(`
			WITH cte AS (
				SELECT fullname, inn, kpp, SUBSTRING(kpp, 1, 2)::int AS region
				FROM nsiorg
				WHERE regnum = '`, p.Org_regnum, `'
			)
			UPDATE notifications
			SET purchaseObjectInfo = '`, p.PurchaseObjectInfo, `',
				href = '`, p.Href, `',
				org_regnum = '`, p.Org_regnum, `',
				org_fullname = fullname,
				org_inn = inn,
				org_kpp = kpp,
				org_region = region,
				org_role = '`, p.Org_role, `',
				placingway_code = '`, p.Placingway_code, `',
				published = LEAST(published, `, p.DocPublishDate, `::timestamp),
				updated = GREATEST(published, `, p.DocPublishDate, `::timestamp),
				ETP_code = '`, p.ETP_code, `',
				maxprice = `, maxPrice, `
			FROM (SELECT null AS void) t
				LEFT JOIN cte  ON 1 = 1
			WHERE pnum = `, p.PurchaseNumber, `
		`)

		_, err = db.Exec(my_query)
		if err != nil {
			log.Println("ERROR Update Notification:", my_query)
		} else {
			log.Println(`Update Notification PurchaseNumber: "` + p.PurchaseNumber + `"`)
		}
	}
	if err != nil { return err}

	// Лоты аукционов, Позиции, Заказчики
	var lot_query []string
	var object_query []string
	var drug_query []string
	var customer_query []string

	var lot_q string
	var object_q string
	var drug_q string
	var customer_q string

	for _, lot := range p.Lots {
		var preferenses []string
		var requirements []string
		var restrictions []string
		var lotsnamed []string

		ESQL(&lot.LotObjectInfo)
		ESQL(&lot.Currency)
		ESQL(&lot.FinanceSource)
		ESQL(&lot.RestrictInfo)
		ESQL(&lot.AddInfo)

		for _, val := range lot.Preferenses {
			if val.Code != "" { preferenses = append(preferenses, val.Code) }
			if val.ShortName != "" { preferenses = append(preferenses, val.ShortName) }
		}
		for _, val := range lot.Requirements {
			if val.Code != "" { requirements = append(requirements, val.Code) }
			if val.ShortName != "" { requirements = append(requirements, val.ShortName) }
		}
		for _, val := range lot.Restrictions {
			if val.Code != "" { restrictions = append(restrictions, val.Code) }
			if val.ShortName != "" { restrictions = append(restrictions, val.ShortName) }
		}
		for _, val := range lot.PurchaseObjects {
			ESQL(&val.Name)
			lotsnamed = append(lotsnamed, val.Name)
		}
		for _, val := range lot.DrugPurchaseObject {
			ESQL(&val.NameR)
			ESQL(&val.NameT)
			lotsnamed = append(lotsnamed, val.NameR + val.NameT)
		}

		lot_q = fmt.Sprint(`
			SELECT `, p.PurchaseNumber, `, `, lot.LotNumber, `, '`, lot.LotObjectInfo, `', `, lot.MaxPrice, `, '`, lot.Currency, `',
				`, lot.QuantityUndefined, `, '`, lot.FinanceSource, `',
				'{`, strings.Join(preferenses, ","), `}'::varchar(20)[],
				'{`, strings.Join(requirements, ","), `}'::varchar(20)[],
				'{`, strings.Join(restrictions, ","), `}'::varchar(20)[], 
				'`, lot.RestrictInfo, `', '`, lot.AddInfo, `', `, lot.TotalSumPO + lot.TotalSumDPO, `, '`, strings.Join(lotsnamed, "\n"), `', `, p.DocPublishDate, `::timestamp`,
		)
		lot_query = append(lot_query, lot_q)

		// Позиции лота
		for _, object := range lot.PurchaseObjects {
			ESQL(&object.OKPD_code)
			ESQL(&object.OKPD2_code)
			ESQL(&object.KTRU_code)
			ESQL(&object.Name)
			ESQL(&object.OKEI)

			var okpd string
			var okpd_v int
			if object.OKPD2_code != "" {
				okpd = object.OKPD2_code
				okpd_v = 2
			} else {
				okpd = object.OKPD_code
				okpd_v = 1
			}

			if len(object.CustomerQuantitys) > 0 {
				for _, quantityC := range object.CustomerQuantitys {
					ESQL(&quantityC.Customer_regnum)

					object_q = fmt.Sprint(`
					SELECT `, p.PurchaseNumber, `, `, lot.LotNumber, `, '`, okpd, `', '`, object.KTRU_code, `', `, okpd_v, `,
						'`, object.Name, `', '`, object.OKEI, `', '`, quantityC.Customer_regnum, `', `, quantityC.Quantity, `, `, object.Quantity, `,
						`, object.Countable, `, `, object.Price, `, `, object.Sum, `, `, p.DocPublishDate, `::timestamp`,
					)
					object_query = append(object_query, object_q)
				}
			} else {
				object_q = fmt.Sprint(`
					SELECT `, p.PurchaseNumber, `, `, lot.LotNumber, `, '`, okpd, `', '`, object.KTRU_code, `', `, okpd_v, `,
						'`, object.Name, `', '`, object.OKEI, `', null, 0, `, object.Quantity, `,
						`, object.Countable, `, `, object.Price, `, `, object.Sum, `, `, p.DocPublishDate, `::timestamp`,
				)
				object_query = append(object_query, object_q)
			}
		}

		// Лекарства лота
		for _, drug := range lot.DrugPurchaseObject {
			ESQL(&drug.ExternalCode)
			ESQL(&drug.NameR)
			ESQL(&drug.NameT)
			ESQL(&drug.GRLSValueR)
			ESQL(&drug.GRLSValueT)
			ESQL(&drug.UserOKEI_codeR)
			ESQL(&drug.UserOKEI_codeT)
			ESQL(&drug.OKPD2_code)

			var okpd string = drug.OKPD2_code
			var okpd_v int = 2

			if len(drug.CustomerQuantitys) > 0 {
				for _, quantityC := range drug.CustomerQuantitys {
					ESQL(&quantityC.Customer_regnum)

					drug_q = fmt.Sprint(`
					SELECT `, p.PurchaseNumber, `, `, lot.LotNumber, `,'`, drug.ExternalCode, `',`, drug.BasicUnitR || drug.BasicUnitT, `,`, drug.IsZNVLP, `,
						'`, okpd, `', `, okpd_v, `,'`, drug.NameR + drug.NameT, `','`, quantityC.Customer_regnum, `',`, quantityC.Quantity, `,`, drug.Quantity, `,
						`, drug.Price, `,`, drug.Sum, `,`, p.DocPublishDate, `::timestamp`,`,
						'`, drug.GRLSValueR + drug.GRLSValueT, `','`, drug.UserOKEI_codeR + drug.UserOKEI_codeT, `'`,
					)
					drug_query = append(drug_query, drug_q)
				}
			} else {
				drug_q = fmt.Sprint(`
					SELECT `, p.PurchaseNumber, `, `, lot.LotNumber, `,'`, drug.ExternalCode, `',`, drug.BasicUnitR || drug.BasicUnitT, `,`, drug.IsZNVLP, `,
						'`, okpd, `', `, okpd_v, `,'`, drug.NameR + drug.NameT, `', null, 0, `, drug.Quantity, `,
						`, drug.Price, `,`, drug.Sum, `,`, p.DocPublishDate, `::timestamp`,`,
						'`, drug.GRLSValueR + drug.GRLSValueT, `','`, drug.UserOKEI_codeR + drug.UserOKEI_codeT, `'`,
				)
				drug_query = append(drug_query, drug_q)
			}
		}

		// Заказчики
		for _, customer := range lot.CustomerRequirements {
			ESQL(&customer.Customer_regnum)
			ESQL(&customer.Customer_fullname)
			ESQL(&customer.DeliveryPlace)

			customer_q = fmt.Sprint(`
				SELECT `, p.PurchaseNumber, ` AS pnum, `, lot.LotNumber, ` AS lotnum, `, customer.MaxPrice, ` AS maxprice, '`,
					customer.Customer_regnum, `' AS customer_regnum, '`, customer.Customer_fullname, `' AS customer_name, `, p.DocPublishDate, `::timestamp AS published`,
			)
			customer_query = append(customer_query, customer_q)
		}
	}

	lot_exec := fmt.Sprint(`
			INSERT INTO notifications_lots AS n (
				pnum, lotnum, lotObjectInfo, maxprice, currency,
				quantityUndefined, finance_source, preferenses, requirements, restrictions,
				restrictInfo, addInfo, totalSum, lotsnamed, n_published
			)`, strings.Join(lot_query, " UNION ALL "),
	)
	object_exec := fmt.Sprint(`
			INSERT INTO notifications_objects AS n (
				pnum, lotnum, okpd, ktru, okpd_v,
				name, okei, customer_regnum, customer_quantity, quantity,
				countable, price, sum, n_published
			)`, strings.Join(object_query, " UNION ALL "),
	)
	drug_exec := fmt.Sprint(`
			INSERT INTO notifications_drugs AS n (
				pnum, lotnum, ExternalCode, BasicUnit, IsZNVLP,
				okpd, okpd_v, name, customer_regnum, customer_quantity, quantity,
				price, sum, n_published, dosageGRLSValue, UserOKEICode
			)`, strings.Join(drug_query, " UNION ALL "),
	)
	customer_exec := fmt.Sprint(`
			INSERT INTO notifications_customers AS n (
				pnum, lotnum, maxprice, customer_regnum, customer_name, n_published
			)`, strings.Join(customer_query, " UNION ALL "),
	)

	// Добавление лотов
	if err == nil {
		tx, err := db.Begin()
		if err != nil {
			log.Println("E1 lot", err)
			return err
		}
		defer tx.Rollback()

		prep_del, err := tx.Prepare("DELETE FROM notifications_lots WHERE pnum = $1;")
		if err != nil {
			log.Println("E2 lot", err)
			return err
		}

		_, err = prep_del.Exec(p.PurchaseNumber)
		prep_del.Close()
		if err != nil {
			log.Println("E3 lot", err)
			return err
		}

		_, err = tx.Exec(lot_exec)
		if err != nil {
			log.Println("ERROR Insert NotificationLot:", lot_exec)
			return err
		}

		err = tx.Commit()
		if err != nil {
			log.Println("ERROR Commit NotificationLot:", lot_exec)
			return err
		} else {
			log.Println(`Insert NotificationLot PurchaseNumber: "` + p.PurchaseNumber + `"`)
		}
	}

	// Добавление позиций
	if err == nil && len(object_query) > 0 {
		tx, err := db.Begin()
		if err != nil {
			log.Println("E1 object", err)
			return err
		}
		defer tx.Rollback()

		prep_del, err := tx.Prepare("DELETE FROM notifications_objects WHERE pnum = $1;")
		if err != nil {
			log.Println("E2 object", err)
			return err
		}

		_, err = prep_del.Exec(p.PurchaseNumber)
		prep_del.Close()
		if err != nil {
			log.Println("E3 object", err)
			return err
		}

		_, err = tx.Exec(object_exec)
		if err != nil {
			log.Println("ERROR Insert NotificationObject:", object_exec)
			return err
		}

		err = tx.Commit()
		if err != nil {
			log.Println("ERROR Commit NotificationObject:", object_exec)
			return err
		} else {
			log.Println(`Insert NotificationObject PurchaseNumber: "` + p.PurchaseNumber + `"`)
		}
	} else if err == nil && len(drug_query) > 0 {	// Добавление лекарств
		tx, err := db.Begin()
		if err != nil {
			log.Println("E1 drug", err)
			return err
		}
		defer tx.Rollback()

		prep_del, err := tx.Prepare("DELETE FROM notifications_drugs WHERE pnum = $1;")
		if err != nil {
			log.Println("E2 drug", err)
			return err
		}

		_, err = prep_del.Exec(p.PurchaseNumber)
		prep_del.Close()
		if err != nil {
			log.Println("E3 drug", err)
			return err
		}

		_, err = tx.Exec(drug_exec)
		if err != nil {
			log.Println("ERROR Insert NotificationDrug:", drug_exec)
			return err
		}

		err = tx.Commit()
		if err != nil {
			log.Println("ERROR Commit NotificationDrug:", drug_exec)
			return err
		} else {
			log.Println(`Insert NotificationDrug PurchaseNumber: "` + p.PurchaseNumber + `"`)
		}
	} else if p.Other != "111" {
		err = errors.New("Empty notifications_drugs and notifications_objects")
	}

	// Добавление заказчиков
	if err == nil && p.Other != "111" {
		tx, err := db.Begin()
		if err != nil {
			log.Println("E1 customer", err)
			return err
		}
		defer tx.Rollback()

		prep_del, err := tx.Prepare("DELETE FROM notifications_customers WHERE pnum = $1;")
		if err != nil {
			log.Println("E2 customer", err)
			return err
		}

		_, err = prep_del.Exec(p.PurchaseNumber)
		prep_del.Close()
		if err != nil {
			log.Println("E3 customer", err)
			return err
		}

		_, err = tx.Exec(customer_exec)
		if err != nil {
			log.Println("ERROR Insert NotificationCustomer:", customer_exec)
			return err
		}

		err = tx.Commit()
		if err != nil {
			log.Println("ERROR Commit NotificationCustomer:", customer_exec)
			return err
		} else {
			log.Println(`Insert NotificationCustomer PurchaseNumber: "` + p.PurchaseNumber + `"`)
		}
		/*
		if regnum != "" {
			db.Exec("update contracts_"+regnum[:1]+" set nmck=$1 where regnum=$2", maxprice, regnum)
			db.Exec("update products_"+regnum[:1]+" set nmck=$1 where regnum=$2 and customerinn=$3", maxprice, regnum, customerinn)
		} //*/
	}

	// Добавление документации
	if err == nil {
		for _, file := range p.Attachments {
			ESQL(&file.PublishedContentId)
			ESQL(&file.FileName)
			ESQL(&file.FileSize)
			ESQL(&file.DocDescription)
			ESQL(&file.DocDate)
			ESQL(&file.Url)
			ESQL(&file.Content)

			if file.DocDate == "" {
				file.DocDate = `null`
			} else {
				file.DocDate = fmt.Sprintf(`'%s'`, file.DocDate)
			}

			file_q := fmt.Sprint(`
				SELECT `, p.PurchaseNumber, ` AS pnum, '`, file.PublishedContentId, `' AS publishedContentId, '`, file.FileName, `' AS name, '`, file.FileSize, `' AS size, '`, file.DocDescription, `' AS description,
					`, file.DocDate, `::timestamp AS docdata, '`, file.Url, `' AS url, '`, file.Content, `' AS content, `, p.DocPublishDate, `::timestamp AS n_published`,
			)

			file_exec := fmt.Sprint(`
				INSERT INTO notifications_files_`, file_year, ` AS n (
					pnum, publishedContentId, name, size, description,
					docdata, url, content, n_published
				)`, file_q, `
				ON CONFLICT (url) DO
				UPDATE
				SET publishedContentId = EXCLUDED.publishedContentId,
					name = EXCLUDED.name,
					size = EXCLUDED.size,
					description = EXCLUDED.description,
					docdata = EXCLUDED.docdata,
					content = EXCLUDED.content,
					n_published = LEAST(n.n_published, EXCLUDED.n_published)
				WHERE n.pnum = EXCLUDED.pnum
					AND n.url = EXCLUDED.url;
			`)

			_, err := db.Exec(file_exec)
			if err != nil {
				log.Println("ERROR Insert NotificationFile:", file_exec)
				return err
			} else {
				log.Println(`Insert/Update NotificationFile  PurchaseNumber: "` + p.PurchaseNumber + `"`)
			}
		}
	}

	return err
}

func (p *Notification) Save(db *sql.DB) error {
	var err error

	log.Println("Saving: ", p.PurchaseNumber, " ", p.DocPublishDate)

	if len(p.Lots) == 0 {
		p.Lot.LotObjectInfo = p.PurchaseObjectInfo
		p.Lots = append(p.Lots, p.Lot)
	}

	var maxPrice float64
	for _, i := range p.Lots {
		maxPrice += i.MaxPrice
		//if i.PurchaseObjects[0].Quantity == 0 && len(i.DrugPurchaseObject[0].ExternalCode) == 0 {
		//	err = errors.New("Error: Q=" + strconv.FormatFloat(i.PurchaseObjects[0].Quantity, 'g', 1, 64) + "; ExtCode=\"" + i.DrugPurchaseObject[0].ExternalCode + "\"")
		//}
	}

	//if err == nil {
		err = p.InsertNotification(db, maxPrice, false, false)
	//}

	return err
}

func ParseNotification(data []byte, db *sql.DB) error {
	var export NotificationExport
	var NotificationAll []Notification
	var g_err error

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(export.NotificationEF.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationEF) }
	if len(export.NotificationEP.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationEP) }
	if len(export.NotificationOKD.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationOKD) }
	if len(export.NotificationOKOU.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationOKOU) }
	if len(export.NotificationOK.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationOK) }
	if len(export.NotificationPOT.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationPOT) }
	if len(export.NotificationZK.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationZK) }
	if len(export.NotificationZP.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationZP) }
	if len(export.NotificationPO.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationPO) }
	if len(export.NotificationZKB.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationZKB) }
	if len(export.NotificationISM.PurchaseNumber) == 19 { export.NotificationISM.Other = "111"; NotificationAll = append(NotificationAll, export.NotificationISM) }
	if len(export.NotificationZakA.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationZakA) }
	if len(export.NotificationZakKOU.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationZakKOU) }
	if len(export.NotificationZakK.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationZakK) }
	if len(export.NotificationZakKD.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationZakKD) }
	if len(export.Notification111.PurchaseNumber) == 19 { export.Notification111.Other = "111"; NotificationAll = append(NotificationAll, export.Notification111) }
	if len(export.NotificationZKBI.PurchaseNumber) == 19 { NotificationAll = append(NotificationAll, export.NotificationZKBI) }

	for i, val := range NotificationAll {
		func(i int, o Notification) {
			ch, err := o.Validate()
			if !ch {
				log.Println("Not valid!", err, o.PurchaseNumber)
			} else {
				if err = o.Save(db); err != nil {
					log.Println(i, "fcsNotification_" + o.Placingway_code, " Err. on", o.PurchaseNumber, ";   DATE:", o.DocPublishDate, err)
					g_err = err
				}
			}
		}(i, val)
	}

	return g_err
}

// Identify prints uid of notification
func (p *Notification) Identify() string {
	log.Println(p.PurchaseNumber + p.Placingway_code)
	return p.Placingway_code
}

// Validate the model
func (p *Notification) Validate() (bool, error) {
	check := false
	var err error

	NotificationS := fmt.Sprint(`
		ID: "`, p.ID, `"
		ExternalID: "`, p.ExternalID, `"
		PurchaseNumber: "`, p.PurchaseNumber, `"
		PurchaseObjectInfo: "`, p.PurchaseObjectInfo, `"
		Href: "`, p.Href, `"
		Org_regnum: "`, p.Org_regnum, `"
		Org_fullName: "`, p.Org_fullName, `"
		Org_inn: "`, p.Org_inn, `"
		Org_kpp: "`, p.Org_kpp, `"
		Org_role: "`, p.Org_role, `"
		Placingway_code: "`, p.Placingway_code, `"
		DocPublishDate: "`, p.DocPublishDate, `"
		ETP_code: "`, p.ETP_code, `"
	`)

	check = check || (p.PurchaseNumber != "")
	if !check {
		err = errors.New("PurchaseNumber is empty!" + NotificationS)
	}

	check = check && (p.DocPublishDate != "")
	if !check {
		err = errors.New("PublishDate is empty!" + NotificationS)
	}
	/*//check = check && (p.Lot != *empty)
	check = check && (p.Attachments != nil)
	if !check {
		err = errors.New("Attachments list is empty!")
	}
	check = check && (p.ObjectInfo != "")
	if !check {
		err = errors.New("ObjectInfo is empty!")
	}*/
	/*
	check = check && (p.Href != "")
	if !check {
		err = errors.New("Not valid!" + NotificationS)
		log.Println(p)
	}
	*/
	/*
	if !check {
		p.Href = "http://zakupki.gov.ru/epz/order/notice/ea44/view/common-info.html?regNumber=" + p.PurchaseNumber
		log.Println("Link was empty: ", p.PurchaseNumber, "Generated a new one: ", p.Link)
		check = true
		//err = errors.New("Link is empty!")
	} //*/
	return check, err
}
