package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

const tableContracts string = ` CREATE table IF NOT EXISTS contracts_%s (
row					SERIAL		PRIMARY KEY,
RegNum				text		UNIQUE,
PurchaseNumber		text		,
Published       	timestamp	,
Signed       		timestamp	,
Exec				timestamp	,
Budget				text		,
OKTMO				text		,
Suppliers			text		,
Customer			text		,
Placer				text		,
Price				float		,
Currency			text		,
Rate				float		,
Paid				float		,
Link				text		,
Status				varchar(5)  ,
Updated				timestamp	,
Version				integer	,
NMCK 				decimal,
Ntname			text
);`

const tableProducts string = ` CREATE table IF NOT EXISTS products_%s (
row					SERIAL		PRIMARY KEY,
RegNum				text		,
PurchaseNumber		text		,
Number				text		,
Name				text		,
OKPD		      	text		,
OKPDInfo	      	text		,
Units		      	text		,
OKEI				text		,
Quantity			float		,
Price				float		,
PriceRu				float		,
Sum					float		,
SumRu				float		,
Rate				float		,
Currency			text		,
Published       	timestamp	,
Signed       		timestamp	,
Exec				timestamp	,
Budget				text		,
Customer			text		,
CustomerRegion     	integer		,
CustomerINN     	text		,
CustomerKPP     	text		,
SupplierINN     	text		,
SupplierKPP     	text		,
SupplierAddress    	text		,
SupplierPhone		text		,
Supplier			text		,
ContractPrice		float		,
Paid				float		,
Status				varchar(5)	,
PlacingWay			text		,
BudgetSource		text		,
Updated				timestamp	,
Version				integer
);`

const tableContractsFiles = `CREATE table IF NOT EXISTS contracts_files_%s (
    row             SERIAL PRIMARY KEY,
    regnum  text,
    description		text,
    name			text,
    url				text,
    size			text,
    data			text,
    ignore			bool,
		downloaded	bool
    );`
const tableContractsProc = `CREATE table IF NOT EXISTS contracts_proc_%s (
    row             SERIAL PRIMARY KEY,
    regnum          text,
    published       timestamp,
    version         integer,
    final           bool,
    stage           text,
    applied         bool,
		paid			decimal,
    shard           text
);`

const tableNotifications string = `CREATE table IF NOT EXISTS notifications_%s (
    pnum            bigint         PRIMARY KEY,
    purchaseNumber  text           UNIQUE,
    objectInfo      text,
    url			    text,
    org_regnum	    text,
    org_inn         text,
    org_name        text,
    org_kpp         text,
    org_role        text,
    org_indx        bigint,
    org_region      int,
    type            text,
    published       timestamp,
    maxPrice        decimal,
    canceled		bool		  default false,
    finished        bool          default false,
		Updated			timestamp
);`
const tableNotificationsLots string = `CREATE table IF NOT EXISTS notifications_lots_%s (
    row             SERIAL        PRIMARY KEY,
    lotnum          integer,
    regnum          text,
    pnum            bigint,
    maxPrice        decimal,
    currency        varchar(5),
    finance_source  text,
	canceled		bool
);`
const tableNotificationsCustomers string = `CREATE table IF NOT EXISTS notifications_customers_%s (
    row             SERIAL       PRIMARY KEY,
    lotnum          integer,
    pnum            bigint,
    maxPrice        decimal,
    customer_name   text,
    customer_inn    text,
    customer_kpp    text,
    customer_indx   bigint,
    customer_regnum text,
    customer_region int,
		regnum					text,
		ctprice					decimal
);`
const tableNotificationsObjects string = `CREATE table IF NOT EXISTS notifications_objects_%s (
    row             SERIAL      PRIMARY KEY,
    pnum            bigint,
    lotnum			integer,
    customer        bigint,
    name			text,
    okpd			text,
    okei			text,
    quantity		decimal,
    countable       bool        default false,
    price			decimal,
    sum				decimal
);`
const tableNotificationsFiles = `CREATE table IF NOT EXISTS notifications_files_%s (
	row             SERIAL     PRIMARY KEY,
	purchaseNumber  bigint,
	description		text,
	name			text,
	url				text,
	size			text,
	data			text,
	ignore			bool,
	downloaded bool
 );`

const tableNsiBudget = `CREATE table IF NOT EXISTS nsiBudget (
    	row					SERIAL		PRIMARY KEY,
    	code				text		UNIQUE,
    	parent				text		,
    	name		     	text		,
    	actual				bool
     );`

const tableNsiOKTMO = `CREATE table IF NOT EXISTS nsiOKTMO (
     row					SERIAL		PRIMARY KEY,
     code				text		UNIQUE,
     parent				text		,
     section				text		,
     fullName			text		,
     actual				bool
     );`
const tableNsiFO = `CREATE table IF NOT EXISTS nsiFO (
 	row					SERIAL		PRIMARY KEY,
	shortname			text		,
	fullname			text		,
	code				integer		,
	region				text
	);`
const tableNsiOKEI = `CREATE table IF NOT EXISTS nsiOKEI (
row					SERIAL		PRIMARY KEY,
Code				text		UNIQUE,
FullName			text,
SectionCode			text,
SectionName			text,
GroupId				int,
GroupName			text,
LocalName			text,
InternationalName	text,
LocalSymbol			text,
InternationalSymbol	text,
Actual				bool);`

const tableNsiOKPD = `CREATE table IF NOT EXISTS nsiOKPD (
row					SERIAL		PRIMARY KEY,
code				text		UNIQUE,
parent				text		,
name		     	text		,
actual				bool
)`

const tableNsiOrg = ` CREATE table IF NOT EXISTS nsiOrg (
   row             SERIAL         PRIMARY KEY,
   RegNum          text           UNIQUE,
   ShortName       text,
   FullName        text,
   OKVED           text,
   Url             text,
   INN             text,
   OGRN            text,
   KPP             text,
   OKTMO           text,
   PostalAddress   text,
   ContactPerson   text,
   SubordinationType   text,
   Actual          bool
 );`

// InitSchema creates tables structure if the schema is empty
func InitSchema(db *sql.DB) error {
	// NSI tables

	if err := runQuery(tableNsiOrg, db); err != nil {
		log.Fatalln(err)
	}
	if err := runQuery(tableNsiOKPD, db); err != nil {
		log.Fatalln(err)
	}
	if err := runQuery(tableNsiOKEI, db); err != nil {
		log.Fatalln(err)
	}
	if err := runQuery(tableNsiOKTMO, db); err != nil {
		log.Fatalln(err)
	}
	if err := runQuery(tableNsiBudget, db); err != nil {
		log.Fatalln(err)
	}
	// Contracts tables
	if err := generateShards(2010, 2016, tableContracts, db); err != nil {
		log.Fatalln(err)
	}
	if err := generateShards(2010, 2016, tableProducts, db); err != nil {
		log.Fatalln(err)
	}
	if err := generateShards(2010, 2016, tableContractsFiles, db); err != nil {
		log.Fatalln(err)
	}
	if err := generateShards(2010, 2016, tableContractsProc, db); err != nil {
		log.Fatalln(err)
	}
	// Notifications tables
	if err := generateShards(2010, 2016, tableNotifications, db); err != nil {
		log.Fatalln(err)
	}
	if err := generateShards(2010, 2016, tableNotificationsFiles, db); err != nil {
		log.Fatalln(err)
	}
	if err := generateShards(2010, 2016, tableNotificationsLots, db); err != nil {
		log.Fatalln(err)
	}
	if err := generateShards(2010, 2016, tableNotificationsObjects, db); err != nil {
		log.Fatalln(err)
	}
	if err := generateShards(2010, 2016, tableNotificationsCustomers, db); err != nil {
		log.Fatalln(err)
	}
	//	log.Fatal("EBAT!")
	// Site tables
	// ...
	return nil
}

func generateShards(beginning, end int, query string, db *sql.DB) error {
	q := query
	for y := beginning; y < end+1; y++ {
		for m := 1; m < 13; m++ {
			shard := fmt.Sprintf("%d_%d", m, y)
			query = strings.Replace(q, "%s", shard, -1)
			//			log.Println(query)
			if err := runQuery(query, db); err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}

func runQuery(q string, db *sql.DB) error {
	_, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
