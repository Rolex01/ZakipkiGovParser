package main

import (
	"encoding/xml"
	//	"fmt"
	"errors"
	"log"
)

type PlacementResult struct { // Извещение о проведении ЭА (электронный аукцион); внесение изменений
	XMLName        xml.Name `xml:"export"`
	PurchaseNumber string   `xml:"fcsPlacementResult>purchaseNumber"`
	LotNumber      int      `xml:"fcsPlacementResult>lotNumber"`
	DocPublishDate string   `xml:"fcsPlacementResult>docPublishDate"`
	Link           string   `xml:"fcsPlacementResult>href"`
	CreateDate     string   `xml:"fcsPlacementResult>createDate"`
	Failed         bool     `xml:"fcsPlacementResult>procedureFailed"`
	Reason         struct {
		Code       string `xml:"code"`
		ObjectName string `xml:"objectName"`
		Name       string `xml:"name"`
		Type       string `xml:"type"`
	} `xml:"fcsPlacementResult>abandonedReason"`

	Result string `xml:"fcsPlacementResult>result"`
}

type NotificationApplication struct {
	JournalNumber string  `xml:"journalNumber"`
	Result        string  `xml:"result"`
	Price         float64 `xml:"price"`
	Rating        string  `xml:"appRating"`
}

func ParsePlacementResult(data []byte) error {
	var export PlacementResult
	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Println(err)
		return err
	}
	ch, err := export.Validate()
	if !ch {
		log.Println("Not valid!", err)
		return err
	}
	export.Identify()
	return nil
}

func (p *PlacementResult) Identify() string {
	log.Println("PlacementResult: ", p.PurchaseNumber, p.Failed)
	return p.PurchaseNumber
}

func (p *PlacementResult) Validate() (bool, error) {
	var err error
	check := true
	if p.PurchaseNumber != "" {
		err = errors.New("PurchaseNumber is empty!")
		return check, err
	}
	if p.DocPublishDate != "" {
		err = errors.New("DocPublishDate is empty!")
		return check, err
	}
	if p.Link != "" {
		err = errors.New("Link is empty!")
		return check, err
	}
	if p.CreateDate != "" {
		err = errors.New("CreateDate is empty!")
		return check, err
	}
	if p.Result != "" {
		err = errors.New("Result is empty!")
		return check, err
	}

	return check, err
}
