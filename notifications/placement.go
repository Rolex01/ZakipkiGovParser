package notifications

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"log"

	_ "github.com/lib/pq"
	"fmt"
)

type PlacementResult struct { // Извещение о проведении ЭА (электронный аукцион); внесение изменений
	XMLName						xml.Name	`xml:"export"`
	PurchaseNumber				string		`xml:"fcsPlacementResult>purchaseNumber"`
	ProtocolNumber				string		`xml:"fcsPlacementResult>protocolNumber"`
	LotNumber					int64		`xml:"fcsPlacementResult>lotNumber"`
	FoundationProtocolNumber	string		`xml:"fcsPlacementResult>foundationProtocolNumber"`
	VersionNumber				int			`xml:"fcsPlacementResult>versionNumber"`
	CreateDate					string		`xml:"fcsPlacementResult>createDate"`
	ProcedurelFailed			bool		`xml:"fcsPlacementResult>procedurelFailed"`
	AbandonedReason	struct {
		Code					string		`xml:"code"`
		ObjectName				string		`xml:"objectName"`
		Name					string		`xml:"name"`
		Type					string		`xml:"type"`
	} `xml:"fcsPlacementResult>abandonedReason"`
	Applications				[]Application	`xml:"fcsPlacementResult>applications>application"`
	Result						string			`xml:"fcsPlacementResult>result"`

	DocPublishDate				string		`xml:"fcsPlacementResult>docPublishDate"`
	Link						string		`xml:"fcsPlacementResult>href"`
}

type Application struct {
	JournalNumber	string	`xml:"journalNumber"`
	AppRating		int64	`xml:"appRating"`
	Result			string	`xml:"result"`
	Price			float64	`xml:"price"`
}

func (p *PlacementResult) Identify() string {
	log.Println("PlacementResult: ", p.PurchaseNumber, p.ProcedurelFailed)
	return p.PurchaseNumber
}

func (p *PlacementResult) Validate() (bool, error) {
	var err error
	check := true
	if p.PurchaseNumber == "" {
		err = errors.New("PurchaseNumber is empty!")
		check = false
	} else if p.CreateDate == "" {
		err = errors.New("CreateDate is empty!")
		check = false
	} else if p.Result == "" {
		err = errors.New("Result is empty!")
		check = false
	}

	return check, err
}

func (p *PlacementResult) Save(db *sql.DB) error {
	var g_err error

	ESQL(&p.PurchaseNumber)
	ESQL(&p.ProtocolNumber)
	ESQL(&p.FoundationProtocolNumber)
	ESQL(&p.CreateDate)
	ESQL(&p.AbandonedReason.Code)
	ESQL(&p.AbandonedReason.ObjectName)
	ESQL(&p.AbandonedReason.Name)
	ESQL(&p.AbandonedReason.Type)
	ESQL(&p.Result)
	ESQL(&p.Link)

	if p.CreateDate == "" {
		p.CreateDate = `null`
	} else {
		p.CreateDate = fmt.Sprintf(`'%s'`, p.CreateDate)
	}

	if p.DocPublishDate == "" {
		p.DocPublishDate = `null`
	} else {
		p.DocPublishDate = fmt.Sprintf(`'%s'`, p.DocPublishDate)
	}

	for _, app := range p.Applications {
		ESQL(&app.JournalNumber)
		ESQL(&app.Result)

		placement_q := fmt.Sprint(`
			SELECT `, p.PurchaseNumber, ` AS pnum, '`, p.ProtocolNumber, `' AS PtcNumber, `, p.LotNumber, ` AS LotNumber, '`, p.FoundationProtocolNumber, `' AS FPtcNumber,
				`, p.VersionNumber, ` AS Version, `, p.CreateDate, `::timestamp AS CreateDate, `, p.ProcedurelFailed, ` AS Failed, 
				'`, p.AbandonedReason.Code, `' AS AR_Code, '`, p.AbandonedReason.ObjectName, `' AS AR_ObjectName, '`, p.AbandonedReason.Name, `' AS AR_Name, '`, p.AbandonedReason.Type, `' AS AR_Type,
				'`, app.JournalNumber, `' AS app_Number, `, app.AppRating, ` AS app_Rating, '`, app.Result, `' AS app_Result, `, app.Price, ` AS app_Price,
				'`, p.Result, `' AS Result, `, p.DocPublishDate, `::timestamp AS DocPublishDate, '`, p.Link, `' AS Link`,
		)

		placement_exec := fmt.Sprint(`
			INSERT INTO notifications_placement AS n (pnum, PtcNumber, LotNumber, FPtcNumber,
				Version, CreateDate, Failed,
				AR_Code, AR_ObjectName, AR_Name, AR_Type,
				app_Number, app_Rating, app_Result, app_Price,
				Result, DocPublishDate, Link
			)`, placement_q, `
			ON CONFLICT (pnum, LotNumber, Version, app_Number) DO
			UPDATE
			SET PtcNumber = EXCLUDED.PtcNumber,
				FPtcNumber = EXCLUDED.FPtcNumber,
				CreateDate = EXCLUDED.CreateDate,
				Failed = EXCLUDED.Failed,
				AR_Code = EXCLUDED.AR_Code,
				AR_ObjectName = EXCLUDED.AR_ObjectName,
				AR_Name = EXCLUDED.AR_Name,
				AR_Type = EXCLUDED.AR_Type,
				app_Rating = EXCLUDED.app_Rating,
				app_Result = EXCLUDED.app_Result,
				app_Price = EXCLUDED.app_Price,
				Result = EXCLUDED.Result,
				DocPublishDate = EXCLUDED.DocPublishDate,
				Link = EXCLUDED.Link
			WHERE n.pnum = EXCLUDED.pnum
				AND n.LotNumber = EXCLUDED.LotNumber
				AND n.Version = EXCLUDED.Version
				AND n.app_Number = EXCLUDED.app_Number;
		`)

		_, err := db.Exec(placement_exec)
		if err != nil {
			log.Println("ERROR Insert NotificationPlacement:", placement_exec)
			g_err = err
		} else {
			log.Println(`Insert/Update NotificationPlacement  PurchaseNumber: "` + p.PurchaseNumber + `"`)
		}
	}

	return g_err
}

func ParsePlacementResult(data []byte, db *sql.DB) error {
	var export PlacementResult

	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Println(err)
		return err
	}

	ch, err := export.Validate()
	if !ch {
		log.Println("Not valid!", err, export.PurchaseNumber)
		return err
	} else {
		if err = export.Save(db); err != nil {
			log.Println("PlacementResult Err. on", export.PurchaseNumber, err)
		}
	}

	return err
}