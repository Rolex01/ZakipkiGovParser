package main

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"log"

	//"bitbucket.org/crmsib/kraken2/shards"
	"bitbucket.org/crmsib/parser_gov/shards"
	//"sync"
)

// Contract scheme according to Protocol v. 5.0 for 2015 y.
type ContractProcedure struct {
	XMLName     xml.Name `xml:"export"`
	Id          uint64   `xml:"contractProcedure>id"`
	ExternalId  string   `xml:"contractProcedure>externalId"`
	PublishDate string   `xml:"contractProcedure>publishDate"`
	RegNum      string   `xml:"contractProcedure>regNum"`
	IKU         string   `xml:"contractProcedure>regNumIKU"`
	Executions  struct {
		EndDate string `xml:"endDate"`
		Year    string `xml:"stage>year"`
		Month   string `xml:"stage>month"`
		OldY    string `xml:"oldStage>year"`
		OldM    string `xml:"oldStage>month"`
	} `xml:"contractProcedure>executions"`
	Number      int         `xml:"ordinalNumber"`
	IsFinal     bool        `xml:"finalStageExecution"`
	D           []Execution `xml:"execution"`
	Termination struct {
		Paid            float64 `xml:"paid"`
		TerminationDate string  `xml:"terminationDate"`
		Reason          string  `xml:"reason"`
	} `xml:"contractProcedure>termination"`
	Stage              string `xml:"contractProcedure>currentContractStage"`
	ModificationReason string `xml:"contractProcedure>modificationReason"`
}

type Execution struct {
	Currency string  `xml:"currency>code"`
	Paid     float64 `xml:"paid"`
	Rate     float64 `xml:"currencyRate>rate"`
	PaidR    float64 `xml:"paidRUR"`
	Product  string  `xml:"product"`
}

func (p *ContractProcedure) GetPaid() float64 {
	var paid float64
	if p.Termination.Paid != 0 {
		log.Println("Termination Paid:", p.Termination)
		return p.Termination.Paid
	}
	for _, v := range p.D {
		log.Println("Paid:", v.Paid, p.D)
		paid += v.Paid
	}
	log.Println(p)
	return paid
}

func (p *ContractProcedure) Validate() bool {
	check := true
	var err error

	if p.RegNum == "" {
		check = false
		err = errors.New("ContractProcedure(null) is not valid: null RegNum!")
		log.Println(err)
	}

	if p.Stage == "" {
		check = false
		err = errors.New("ContractProcedure(" + p.RegNum + ") is not valid: null Stage!")
		log.Println(err)
	}

	if p.PublishDate == "" {
		check = false
		err = errors.New("ContractProcedure(" + p.RegNum + ") is not valid: null PublishDate!")
		log.Println(err)
	}
	return check
}
func UpdateContract(tx *sql.Tx, regnum, shard, stage, termination, updated string, paid float64) error {
	var err error
	log.Println("UpdatingProc: ", regnum, shard, stage, termination, updated, paid)
	if termination != "" {
		_, err := tx.Exec("UPDATE contracts_"+shard+" SET status=$2, paid=$3, updated=$5, exec=$4 WHERE regnum=$1", regnum, stage, paid, termination, updated)
		if err != nil {
			return err
		}
		_, err = tx.Exec("UPDATE products_"+shard+" SET status=$2, paid=$3, updated=$5, exec=$4 WHERE regnum=$1", regnum, stage, paid, termination, updated)
		if err != nil {
			return err
		}
	} else {
		_, err := tx.Exec("UPDATE contracts_"+shard+" SET status=$2, paid=$3, updated=$4 WHERE regnum=$1", regnum, stage, paid, updated)
		if err != nil {
			return err
		}
		_, err = tx.Exec("UPDATE products_"+shard+" SET status=$2, paid=$3, updated=$4 WHERE regnum=$1", regnum, stage, paid, updated)
		if err != nil {
			return err
		}
	}
	return err
}

func CancelContract(tx *sql.Tx, regnum, shard, stage, updated string) error {
	log.Println("CancelProc: ", regnum, shard, stage, updated)
	_, err := tx.Exec("UPDATE contracts_"+shard+" SET status=$2, updated=$3, exec=$3 WHERE regnum=$1", regnum, stage, updated)
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE products_"+shard+" SET status=$2, updated=$3, exec=$3 WHERE regnum=$1", regnum, stage, updated)
	if err != nil {
		return err
	}
	return err
}

func (p *ContractProcedure) Save(db *sql.DB) error {
	var shard string

	var date string
	if date = p.PublishDate; date == "" {
		date = p.Termination.TerminationDate
	}

	_, err := db.Exec(`INSERT INTO contracts_proc_`+shards.GetByRegnum(p.RegNum)+` (regnum, published, stage, final, paid, applied) VALUES ($1, to_timestamp(translate($2, 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'), $3, $4, $5, $6) RETURNING row;`,
		p.RegNum, p.PublishDate, p.Stage, p.IsFinal, p.GetPaid(), false)
	if err != nil {
		err = fmt.Errorf("E! R# %s SHARDEC %s %q %q %q %q", p.RegNum, shard, date, p.Executions, p, err)
		log.Println(err)
	}
	return err
}

/*func (p *ContractProcedure) Save(db *sql.DB) error {
	var exist, fresh bool
	var shard string
	var err error
	log.Println("EBANIY STID!", p.GetPaid())

	exist, fresh, shard, err = searchContractByRegNum(p.RegNum, p.PublishDate, db)
	if err != nil {
		return err
	}
	if !exist {
		//err = errors.New("ContractProcedureCancel error - contract didn't found: " + p.RegNum + "; " + shard)
		log.Println("ContractProcedureCancel error - contract didn't found: " + p.RegNum + "; " + shard)
		return nil
	}
	if !fresh {
		log.Println("ContractProcedureCancel warning - contract is already updated: " + p.RegNum + "; " + p.PublishDate)
		return nil
	} else {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()
		err = UpdateContract(tx, p.RegNum, shard, p.Stage, p.Termination.TerminationDate, p.PublishDate, p.GetPaid())
		err = tx.Commit()
		return err
	}
	return err
}*/

func (p *ContractProcedureCancel) Save(db *sql.DB) error {
	//log.Println(ShardsID(p.CancelDate))
	_, err := db.Exec(`INSERT INTO contracts_proc_`+shards.GetByRegnum(p.RegNum)+` (regnum, published, stage, final, paid, applied) VALUES ($1, to_timestamp(translate($2, 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'), $3, $4, $5, $6) RETURNING row;`,
		p.RegNum, p.CancelDate, p.Stage, false, 0., false)
	if err != nil {
		err = fmt.Errorf("E!! R# %s %q %q", p.RegNum, p, err)
		log.Println(err)
	}
	return err
}

/*func (p *ContractProcedureCancel) Save(db *sql.DB) error {
	var exist, fresh bool
	var shard string
	var err error
	exist, fresh, shard, err = searchContractByRegNum(p.RegNum, p.CancelDate, db)
	if err != nil {
		return err
	}
	if !exist {
		//err = errors.New("ContractProcedure error - contract didn't found: " + p.RegNum + "; " + shard)
		log.Println("ContractProcedure error - contract didn't found: " + p.RegNum + "; " + shard)
		return nil
	}
	if !fresh {
		log.Println("ContractProcedure warning - contract is already updated: " + p.RegNum + "; " + p.CancelDate)
		return nil
	} else {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		defer tx.Rollback()
		err = CancelContract(tx, p.RegNum, shard, p.Stage, p.CancelDate)
		err = tx.Commit()
		return err
	}
	return err
}
*/
type ContractProcedureCancel struct {
	XMLName      xml.Name `xml:"export"`
	RegNum       string   `xml:"contractProcedureCancel>regNum"`
	CancelDate   string   `xml:"contractProcedureCancel>cancelDate"`
	DocumentBase string   `xml:"contractProcedureCancel>documentBase"`
	Stage        string   `xml:"contractProcedureCancel>currentContractStage"`
	/*
		E 	Исполнение;
		ET	Исполнение прекращено;
		EC	Исполнение завершено;
		IN	Aннулировано.
	*/
}

// Looking for contract between available shards. Return: 1. Exist 2. Updated 3. Shard
type searchContractByRegNumSt struct {
	err     error
	exist   bool
	updated bool
	shd     string
}

func searchContractByRegNum(regnum, published string, db *sql.DB) (bool, bool, string, error) {
	//var found bool
	log.Println("SearchContractByRegNum: ", regnum, published)
	//res := make(chan searchContractByRegNumSt, 24)
	//sc := new(sync.WaitGroup)
	var c int
	err := db.QueryRow("select count(*) from contracts_"+shards.GetByRegnum(regnum)+" where regnum = $1;", regnum).Scan(&c)
	if err != nil || c == 0 {

	} else {
		if c != 0 {
			c = 0
			err = db.QueryRow("select count(*) from contracts_"+shards.GetByRegnum(regnum)+" where regnum = $1 and published > $2;", regnum, published).Scan(&c)
			if c != 0 {
				return true, true, shards.GetByRegnum(regnum), err
			} else {
				return true, false, shards.GetByRegnum(regnum), err
			}
		}
	}

	return false, false, "", nil
}

func (p *ContractProcedureCancel) Validate() bool {
	check := false
	check = check || (p.RegNum != "")
	check = check && (p.CancelDate != "")
	check = check && (p.Stage != "")
	return check
}

func ParseContractProcedure(data []byte, db *sql.DB) error {
	var export ContractProcedure
	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Println(err)
		return err
	}
	return export.Save(db)
	//log.Println(export.RegNum)
}
func ParseContractCancel(data []byte, db *sql.DB) error {
	var export ContractProcedureCancel
	err := xml.Unmarshal(data, &export)
	if err != nil {
		log.Println(err)
		return err
	}
	return export.Save(db)
}
