package notifications

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"log"

	common "bitbucket.org/crmsib/parser_gov/notifications"
	_ "github.com/lib/pq"
)

//NotificationLotCancel - Извещение о проведении ЭА (электронный аукцион); внесение изменений
type NotificationLotCancel struct { // Извещение о проведении ЭА (электронный аукцион); внесение изменений
	XMLName xml.Name `xml:"export"`
	// Информация о порядке проведения - неинтересно. 24.03.2015 - единогласно.
	ID             string `xml:"fcsNotificationLotCancel>id"`
	ExternalID     string `xml:"fcsNotificationLotCancel>externalId"`
	PurchaseNumber string `xml:"fcsNotificationLotCancel>purchaseNumber"` // Идентификатор документа! Супер важная инфа
	DocPublishDate string `xml:"fcsNotificationLotCancel>docPublishDate"`
	LotNumber      int    `xml:"fcsNotificationLotCancel>lot>lotNumber"`
}

// Save notification to database
func (p *NotificationLotCancel) Save(db *sql.DB) error {
	log.Println("Cancelling lot #", p.LotNumber, " at ", p.PurchaseNumber)
	err := common.CancelLot(db, p.PurchaseNumber, p.DocPublishDate, p.LotNumber)
	if err != nil {
		log.Println("Error on lotCancel: ", p.PurchaseNumber, err)
	}
	return err
}

// Identify prints uid of notification
func (p *NotificationLotCancel) Identify() string {
	log.Println(p.PurchaseNumber, "[", p.LotNumber, "]")
	return p.PurchaseNumber
}

// ParseNotificationLotCancel parse xml raw data, fill the model and validate it
func ParseNotificationLotCancel(data []byte, db *sql.DB) error {
	var export NotificationLotCancel
	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Println(err)
		return err
	}
	ch, err := export.Validate()
	if !ch {
		log.Println("Not valid!", err)
	} else {
		err = export.Save(db)
		export.Identify()
	}
	return err
}

// Validate the model
func (p *NotificationLotCancel) Validate() (bool, error) {
	check := false
	var err error
	check = check || (p.PurchaseNumber != "")
	if !check {
		err = errors.New("PurchaseNumber is empty!")
	}
	check = check && (p.DocPublishDate != "")
	if !check {
		err = errors.New("PublishDate is empty!")
	}

	return check, err
}
