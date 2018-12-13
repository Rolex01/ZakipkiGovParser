package notifications

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"log"

	common "bitbucket.org/crmsib/parser_gov/notifications"
	_ "github.com/lib/pq"
)

//NotificationCancel - Извещение об отмене аукциона.
type NotificationCancel struct { // Извещение о проведении ЭА (электронный аукцион); внесение изменений
	XMLName        xml.Name `xml:"export"`
	ID             string   `xml:"fcsNotificationCancel>id"`
	ExternalID     string   `xml:"fcsNotificationCancel>externalId"`
	PurchaseNumber string   `xml:"fcsNotificationCancel>purchaseNumber"` // Идентификатор документа! Супер важная инфа
	DocPublishDate string   `xml:"fcsNotificationCancel>docPublishDate"`
}

// Save notification to database
func (p *NotificationCancel) Save(db *sql.DB) error {
	log.Println("Cancelling: ", p.PurchaseNumber)

	err := common.CancelNotification(db, p.PurchaseNumber, p.DocPublishDate, "TRUE")

	if err != nil {
		log.Println("Error on (notification cancel): ", p.PurchaseNumber, err)
	}

	return err
}

// Identify prints uid of notification
func (p *NotificationCancel) Identify() string {
	log.Println("Cancel: ", p.PurchaseNumber)
	return p.PurchaseNumber
}

// ParseNotificationCancel parse xml raw data, fill the model and validate it
func ParseNotificationCancel(data []byte, db *sql.DB) error {
	var export NotificationCancel
	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Println(err)
		return err
	}
	ch, err := export.Validate()
	if !ch {
		log.Println("Not valid!", err)
	} else {
		export.Save(db)
		export.Identify()
	}
	//log.Println(export.RegNum)
	return nil
}

// Validate the model
func (p *NotificationCancel) Validate() (bool, error) {
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
