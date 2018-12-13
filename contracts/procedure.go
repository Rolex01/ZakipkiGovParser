package contracts

import (
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"

	log "github.com/Sirupsen/logrus"

	"bitbucket.org/crmsib/parser_gov/shards"
	//"sync"
)

// Contract scheme according to Protocol v. 5.0 for 2015 y.
type ContractProcedure struct {
	XMLName       xml.Name `xml:"export"`
	Id            uint64   `xml:"contractProcedure>id"`
	ExternalId    string   `xml:"contractProcedure>externalId"`
	PublishDate   string   `xml:"contractProcedure>publishDate"`
	VersionNumber int      `xml:"contractProcedure>versionNumber"` /* Номер редакции сведений */
	RegNum        string   `xml:"contractProcedure>regNum"`
	IKU           string   `xml:"contractProcedure>regNumIKU"`
	Executions    struct {
		EndDate string      `xml:"endDate"`
		Year    string      `xml:"stage>year"`
		Month   string      `xml:"stage>month"`
		OldY    string      `xml:"oldStage>year"`
		OldM    string      `xml:"oldStage>month"`
		IsFinal bool        `xml:"finalStageExecution"`
		D       []Execution `xml:"execution"`
	} `xml:"contractProcedure>executions"`
	Number int `xml:"ordinalNumber"`
	//IsFinal     bool        `xml:"finalStageExecution"`
	//D           []Execution `xml:"execution"`
	Termination struct {
		Paid            float64 `xml:"paid"`
		TerminationDate string  `xml:"terminationDate"`
		Reason          string  `xml:"reason"`
	} `xml:"contractProcedure>termination"`
	Stage              string `xml:"contractProcedure>currentContractStage"`
	ModificationReason string `xml:"contractProcedure>modificationReason"`
}

type Execution struct {
	Currency   string  `xml:"currency>code"`
	Paid       float64 `xml:"paid"`
	Rate       float64 `xml:"currencyRate>rate"`
	PaidR      float64 `xml:"paidRUR"`
	Product    string  `xml:"product"`
	Name       string  `xml:"docExecution>name"`
	Date       string  `xml:"docExecution>documentDate"`
	Number     string  `xml:"docExecution>documentNum"`
	PayDocName string  `xml:"payDoc>documentName"`
	PayDocDate string  `xml:"payDoc>documentDate"`
	PayDocNum  string  `xml:"payDoc>documentNum"`
	IsFinal    bool    `xml:"finalStageExecution"`
}

func (p *Execution) Normalize() {

	if p.Name == "" && p.PayDocName != "" {
		p.Name = p.PayDocName
	}

	if p.Date == "" && p.PayDocDate != "" {
		p.Date = p.PayDocDate
	}

	if p.Number == "" && p.PayDocNum != "" {
		p.Number = p.PayDocNum
	}
}

func (p *ContractProcedure) GetPaid() float64 {
	var paid float64
	if p.Termination.Paid != 0 {
		//log.Println("Termination Paid:", p.Termination)
		return p.Termination.Paid
	}
	for _, v := range p.Executions.D {
		//log.Println("Paid:", v.Paid, p.D)

		paid += v.Paid
	}
	//log.Println(p)
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

func UpdateContract(tx *sql.DB, regnum, shard, stage, termination, updated string, paid float64) error {
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

func CancelContract(tx *sql.DB, regnum, stage, updated string) error {
	log.Println("CancelProc: ", regnum, stage, updated)
	_, err := tx.Exec("UPDATE contracts_"+shards.GetByRegnum(regnum)+" SET status=$2, updated=$3, exec=$3 WHERE regnum=$1", regnum, stage, updated)
	if err != nil {
		return err
	}
	_, err = tx.Exec("UPDATE products_"+shards.GetByRegnum(regnum)+" SET status=$2, updated=$3, exec=$3 WHERE regnum=$1", regnum, stage, updated)
	if err != nil {
		return err
	}
	return err
}

func (p *ContractProcedure) Save(db *sql.DB, data []byte) error {

	var date string

	if date = p.PublishDate; date == "" {
		date = p.Termination.TerminationDate
	}

	//log.Infoln("SAV! - 0")

	var c int

	if p.RegNum == "" {
		log.Infoln("Not Valid data in Contract Proc!")
		return ParseContractCancel(data, db)
		//log.Fatal("Not Valid data!\n", string(data))
	}

	log.Info("Going to make bad things... ", fmt.Sprintf("SELECT COUNT(*) FROM contracts_proc_"+shards.GetByRegnum(p.RegNum)+" WHERE id=%d;", p.Id))

	err := db.QueryRow(`
		SELECT COUNT(*)
		FROM contracts_proc_`+shards.GetByRegnum(p.RegNum)+`
		WHERE id=$1;
	`, p.Id).Scan(&c)

	if err != nil {
		return err
	}

	p.SaveExecutions(db)

	var n0 int64 = 0

	if c == 0 {

		//log.Infoln("SAV! - 1")

		var paid float64
		var status string

		err := db.QueryRow(`
			SELECT status, paid
			FROM contracts_`+shards.GetByRegnum(p.RegNum)+`
			WHERE regnum=$1
		`, p.RegNum).Scan(&status, &paid)

		if err == nil && paid == 0 && p.GetPaid() == 0 && status == p.Stage {
			log.Info("Nothing to update here. ", p.RegNum)
		} else {

			//log.Infoln("SAV! - 2")

			res0, err := db.Exec(`
				UPDATE contracts_`+shards.GetByRegnum(p.RegNum)+`
				SET status=$2, paid=$3
				WHERE regnum=$1
			`, p.RegNum, p.Stage, paid+p.GetPaid())

			//log.Infoln("SAV! - 3", err, p.Stage, status, paid)

			if err != nil {

				log.Fatal(err)

				return fmt.Errorf(`
					UPDATE contracts_`+shards.GetByRegnum(p.RegNum)+`
					SET status='%s', paid=(
						SELECT paid
						FROM contracts_`+shards.GetByRegnum(p.RegNum)+`
						WHERE regnum='%s'
					) + %f
					WHERE regnum='%s'; With additional error: %v
				`, p.Stage, p.RegNum, p.GetPaid(), p.RegNum, err)

				/*res1, err := db.Exec(`Update products_`+shards.GetByRegnum(p.RegNum)+` set Status=$2, Paid=(select Paid from contracts_`+shards.GetByRegnum(p.RegNum)+` where regnum=$1) + $3 where regnum=$1;`, p.RegNum,
				p.Stage, p.GetPaid())*/
			}

			n0, _ = res0.RowsAffected()

			/*_, err = db.Exec(fmt.Sprintf(`Update products_`+shards.GetByRegnum(p.RegNum)+` set Status='%s', Paid=(select Paid from contracts_`+shards.GetByRegnum(p.RegNum)+` where regnum='%s') + %f where regnum='%s';`, p.Stage, p.RegNum, p.GetPaid(), p.RegNum))

			n0, _ = res0.RowsAffected()
			if err != nil {
				return fmt.Errorf(fmt.Sprintf(`Update products_`+shards.GetByRegnum(p.RegNum)+` set Status='%s', Paid=(select Paid from contracts_`+shards.GetByRegnum(p.RegNum)+` where regnum='%s') + %f where regnum='%s'; With error: %v`, p.Stage, p.RegNum, p.GetPaid(), p.RegNum, err))
			}*/

			//res1.RowsAffected()
		}

		_, err = db.Exec(`
			INSERT INTO contracts_proc_`+shards.GetByRegnum(p.RegNum)+` (regnum, id, published, stage, final, paid, applied, version)
			VALUES ($1, $2, to_timestamp(translate($3, 'T', ' '), 'YYYY-MM-DD HH24:MI:SS'), $4, $5, $6, $7, $8)
			RETURNING row
		`, p.RegNum, p.Id, p.PublishDate, p.Stage, p.Executions.IsFinal, p.GetPaid(), n0 > 0, p.VersionNumber)

		return err
	}

	return nil
}

func (p *ContractProcedure) SaveExecutions(db *sql.DB) {

	for _, execution := range p.Executions.D {

		execution.Normalize()

		_, err := db.Exec(`
			INSERT INTO executions_`+shards.GetByRegnum(p.RegNum)+` (regnum, name, number, date, currency, paid)
			VALUES ($1, $2, $3, to_timestamp($4, 'YYYY-MM-DD'), $5, $6)
		`, p.RegNum, execution.Name, execution.Number, execution.Date, execution.Currency, execution.Paid)

		if err != nil {
			log.Infof("Record '%s' '%s' '%s' '%s' already exists.", p.RegNum, execution.Number, execution.Name, execution.Date)
		}
	}
}

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

type ContractProcedureCancel struct {
	XMLName      xml.Name `xml:"export"`
	ProcID       int64    `xml:"cancelledProcedureId"` /* Идентификатор отменяемой информации об исполнении (расторжении) контракта */
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
	return export.Save(db, data)
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
