package main

import (
	//"encoding/json"
	//"fmt"
	"log"
	"os"
	//"runtime"
	//	"sync"
	"flag"
	"strings"
	"time"

	"bitbucket.org/crmsib/aftertools"
	//"bitbucket.org/crmsib/parser_gov/shards"
	"path/filepath"
	"sort"
	"bytes"
	"archive/zip"
	//"sync"
	//"sync"
	"fmt"
	"database/sql"
)

type ZipArray struct {
	PathZip		string
	NameZip		string
	ModTime		time.Time
}
/*
func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(8)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

var (
	MAXPROC int = MaxParallelism()
)
*/
/*
func getPath(path string, folder2 string, folder3 string, parsemod int) []string {
	var p []string
	if path == "" {
		return append(p, "./data")
	}
	if parsemod == 0 || parsemod == 1 {
		p = strings.Split(path, ",")
	}
	if parsemod == 0 || parsemod == 2 {
		p = append(p, strings.Split(path, ",")[0] + folder2)
	}
	if parsemod == 0 || parsemod == 3 {
		p = append(p, strings.Split(path, ",")[0] + folder3)
	}
	log.Println("getPath", path, folder2, folder3, parsemod, p)
	//time.Sleep(time.Second * time.Duration(10))
	return p
}
//*/

var ZipList []ZipArray

func main() {
	start := time.Now()

	var counters Stats
	var flagFile bool
	var flagPath string
	var flagDBHost string
	var flagWait int
	var flagDBCons int
	var flagPool int
	//var flagDBRedis int
	//var flagRedisHost string
	//var flagRedisPass string
	var flagDBName string
	var flagDBUser string
	var flagDBPassword string
	var flagAfterTools string
	var flagLinear bool
	var flagInitIndxes bool
	var flagForce bool
	var flagPattern string

	var flagIsSave int
	var flagIsCheck int
	var flagIsMssql int

	var f_gmp int
	var f_sleep int
	var f_mod int
	var f_zipdate_min string
	var f_zipdate_max string

	flag.BoolVar(&flagFile, "log2file", false, "Record log to file")
	flag.BoolVar(&flagInitIndxes, "init-indexes", false, "Init indexes (first run only!)")
	flag.BoolVar(&flagLinear, "linear", false, "Process only one zip per time (for nsi files)")
	flag.StringVar(&flagPath, "path", "", "Location(s) to parse. Use ',' separator for several locations")
	flag.StringVar(&flagDBName, "dbname", "good", "Use custom database name")
	//flag.StringVar(&flagRedisHost, "hstredis", "", "Use custom redis host")
	//flag.StringVar(&flagRedisPass, "pswdredis", "", "Use custom redis pass")
	flag.StringVar(&flagDBUser, "dbuser", "", "Use custom user name")
	flag.StringVar(&flagDBPassword, "dbpassword", "", "Use custom user password")
	flag.StringVar(&flagDBHost, "dbhost", "localhost", "Use custom database host")
	flag.StringVar(&flagAfterTools, "aftertools", "", "Enable specific post-processing tool")
	flag.StringVar(&flagPattern, "pattern", "*.zip", "Custom pattern for file scan")
	flag.BoolVar(&flagForce, "force", false, "Force the actions")
	flag.IntVar(&flagWait, "wait", 60, "Delay time in seconds before auto kill. (For testing)")
	flag.IntVar(&flagDBCons, "dbcons", 1, "Count of database connections")
	//flag.IntVar(&flagDBRedis, "dbredis", 0, "Set custom redis instance")
	flag.IntVar(&flagPool, "pool", 1, "Count of goroutines")
	flag.StringVar(&f_zipdate_min, "zd_min", "20131231", "Minimum file (zip) change date")
	flag.StringVar(&f_zipdate_max, "zd_mzx", "20991231", "Maximum file (zip) change date")

	flag.IntVar(&flagIsSave, "issave", 1, "Saved? (1 - Yes, 0 - No)")
	flag.IntVar(&flagIsCheck, "ischeck", 1, "Checked? (1 - Yes, 0 - No)")
	flag.IntVar(&flagIsMssql, "ismssql", 0, "MSSQL? (1 - Yes, 0 - No)")

	flag.IntVar(&f_gmp, "gomaxprocs", 16, "Define max procs for go")
	flag.IntVar(&f_sleep, "sleep", 30, "Define sleep second after [mod] files download")
	flag.IntVar(&f_mod, "mod", 50, "Define mod count for sleep activate")

	flag.Parse()

	//runtime.GOMAXPROCS(f_gmp)

	
	

	if flagFile {
		logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer logFile.Close()
		log.Println("Writing log to file")
		log.SetOutput(logFile)
	}

	//ToDo load database params from settings file?
	//_db, _ := InitPSQL(&Database{"", "", "localhost", "", flagDBName, flagDBCons, flagDBCons})
	//db := deferstats.NewDB(_db)
	db, _ := InitPSQL(&Database{flagDBUser, flagDBPassword, flagDBHost, "", flagDBName, flagDBCons, int(flagDBCons / 2)})
	/*var err error
	if flagIsMssql == 0 {
		db, err = sql.Open("postgres",
			fmt.Sprintf(
				"host=%s port=%d user=%s password=%s dbname=%s connect_timeout=60 sslmode=disable",
				flagDBHost, 5432, flagDBUser, flagDBPassword, flagDBName,
			),
		)
		if err != nil {
			fmt.Println("Error in connect DB1")
			log.Fatal(err)
		}
	} else {
		db, err = sql.Open ("mssql",
			fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;",
				"localhost", "test", "test", "testdb",
			),
		)
		if err != nil {
			fmt.Println("Error in connect DB2")
			log.Fatal(err)
		}
	}*/
	defer db.Close()
	// ToDo create InitStructure function for next bunch of code:

	/*
	shards.InitSchema(db)
	if flagInitIndxes {
		shards.InitIndexes(db)
	}
	//*/

	//log.Println("CPU Num usage:", MAXPROC)
	switch flagAfterTools {
		case "budget":
			log.Println("Applying Budget")
			aftertools.ApplyBudget(db)
			log.Println("Total WT:", time.Since(start))
		case "pnames":
			log.Println("Applying ProductNames")
			aftertools.ApplyProductNames(db)
			log.Println("Total WT:", time.Since(start))
		case "nmck":
			log.Println("Applying NMCK")
			aftertools.ApplyNMCK(flagForce, db)
			log.Println("Total WT:", time.Since(start))
		case "ncustomers":
			log.Println("Applying NCustomers")
			aftertools.ApplyNCustomers(flagForce, db)
			log.Println("Total WT:", time.Since(start))
		case "pcfix":
			log.Println("Applying ProductsCountFix")
			aftertools.ApplyProductsCountFix(flagForce, db)
			log.Println("Total WT:", time.Since(start))
		case "paidfix":
			aftertools.ApplyPaidFix(flagForce, db)
		case "":

			//	Получаем список Архивов

			walk := func(path string, info os.FileInfo, err error) error {
				if err != nil {
					log.Fatalln(err)
				}

				matched, err := filepath.Match(flagPattern, info.Name())
				if err != nil {
					log.Fatalln(err)
				}

				if !info.IsDir() && matched && f_zipdate_min < info.ModTime().Format("20060102") && f_zipdate_max > info.ModTime().Format("20060102") {
					//log.Println(path, info.Size(), info.ModTime())
					ZipList = append(ZipList, ZipArray{path, info.Name(), info.ModTime()})
					//log.Println("From1:", info.Name(), !info.IsDir() && matched, len(ch))
				}

				return nil
			}

			filepath.Walk(flagPath, walk)

			sort.Slice(ZipList, func(i, j int) bool {
				// return ZipList[i].ModTime.Sub(ZipList[j].ModTime) < 0
				return (ZipList[i].ModTime.Sub(ZipList[j].ModTime) < 0) || (ZipList[i].ModTime.Sub(ZipList[j].ModTime) == 0 && ZipList[i].NameZip < ZipList[j].NameZip)
				// return (p[i].X < p[j].X) || (p[i].X == p[j].X && p[i].Y < p[j].Y)
			})



			//go ZipWorker(zip_pool, files, archives, db, flagLinear, linear, done, flagIsCheck)

			//Перебираем список архивов, получаем список xml
			for i, val := range ZipList {
				log.Println(i, val.PathZip, val.ModTime)

				log.Println("ZipWorker :: Open (", val.PathZip, ")")

				z := new(ZipFile)
				if &val.NameZip == nil {
					panic("Empty zip name")
				}
				z.Init(val.PathZip)

				r, err := zip.OpenReader(val.PathZip)
				if err != nil {
					log.Println("Trouble with zip file: ", val.PathZip, ";  error:", err)
					continue
				}

				i := 0
				xmlLenght := len(r.File)
				if xmlLenght != 0 {
					for k, f := range r.File {
						//if k > 0 { break }

						if r.File[k] == nil {
							panic("Empty XML file")
						}

						if strings.Contains(f.Name, ".sig") {
							log.Println(f.Name, "Skipping... SIG")
							continue
						}
						x := new(XmlFile)

						x.Init(f.Name, z.Bucket(), z.Key(), val.PathZip)

						xmlExist := false

						if flagIsCheck == 1 {
							xmlExist, err = HaveXml(x, db)
							if err != nil {
								log.Println(err)
							}
						}
						//log.Println("CHEK!", x, xmlExist)

						if !xmlExist {
							i++
							rc, err := f.Open()
							if err != nil {
								log.Fatal(err)
							}

							buf := new(bytes.Buffer)
							buf.ReadFrom(rc)
							rc.Close()
							x.Data = buf.Bytes()


							err = ParserWorker(db, x, flagIsSave, &counters)

							if err != nil {
								log.Fatal(err)
							} else {
								//if flagIsSave == 1 {
								//	err = SaveXml(x, db)
								//}
								log.Println("GOOD! ZIP:", val.NameZip, "; XML:", f.Name)
							}
							//log.Println(x.Name, "\n\n\n", string(x.Data))
						}
					}

					err = r.Close()
					if err != nil {
						log.Fatal(err)
					}
				}
			}

			log.Println("Pre Verification WT:", time.Since(start))

			// Part 4: Verification
			log.Println("Total WT:", time.Since(start))
			//log.Println("TGoroutines", runtime.NumGoroutine())
			log.Println("Stats| Contracts: ", counters.contractswithErrors, "\\", counters.contracts, " ContractProc: ", counters.contractsProcwithErrors, "\\", counters.contractsProc, ";", " ContractCancel: ", counters.contractsCancel, ";")

			log.Println("nsiOrganizationType: ", counters.nsiOrganizationTypewithErrors, counters.nsiOrganizationType)
			log.Println("nsiOrganization: ", counters.nsiOrganizationwithErrors, counters.nsiOrganization)
			log.Println("nsiBudget: ", counters.nsiBudgetwithErrors, counters.nsiBudget)
			log.Println("nsiOKEI: ", counters.nsiOKEIwithErrors, counters.nsiOKEI)
			log.Println("nsiOKPD: ", counters.nsiOKPDwithErrors, counters.nsiOKPD)
			log.Println("nsiOKPD2: ", counters.nsiOKPD2withErrors, counters.nsiOKPD2)
			log.Println("nsiOKTMO: ", counters.nsiOKTMOwithErrors, counters.nsiOKTMO)
			log.Println("nsiOffBudget: ", counters.nsiOffBudgetwithErrors, counters.nsiOffBudget)
			log.Println("nsiOffBudgetType: ", counters.nsiOffBudgetTypewithErrors, counters.nsiOffBudgetType)
			log.Println("nsiPurchasePreference: ", counters.nsiPurchasePreferencewithErrors, counters.nsiPurchasePreference)
			log.Println("nsiETPs: ", counters.nsiETPswithErrors, counters.nsiETPs)
			log.Println("nsiPlacingWayList: ", counters.nsiPlacingWayListwithErrors, counters.nsiPlacingWayList)
			log.Println("nsiFarmDrugsDictionary: ", counters.nsiFarmDrugsDictionarywithErrors, counters.nsiFarmDrugsDictionary)

			log.Println("notificationEF: ", counters.notificationEFwithErrors, counters.notificationEF)
			log.Println("notificationOK: ", counters.notificationOKwithErrors, counters.notificationOK)
			log.Println("notificationEP: ", counters.notificationEPwithErrors, counters.notificationEP)
			log.Println("notificationOKD: ", counters.notificationOKDwithErrors, counters.notificationOKD)
			log.Println("notificationOKOU: ", counters.notificationOKOUwithErrors, counters.notificationOKOU)
			log.Println("notificationOKU: ", counters.notificationOKUwithErrors, counters.notificationOKU)
			log.Println("notificationPOT: ", counters.notificationPOTwithErrors, counters.notificationPOT)
			log.Println("notificationPO: ", counters.notificationPOwithErrors, counters.notificationPO)
			log.Println("notificationZK: ", counters.notificationZKwithErrors, counters.notificationZK)
			log.Println("notificationZKB: ", counters.notificationZKBwithErrors, counters.notificationZKB)
			log.Println("notificationZP: ", counters.notificationPOwithErrors, counters.notificationZP)
			log.Println("notificationISM: ", counters.notificationISMwithErrors, counters.notificationISM)
			log.Println("notificationINM111: ", counters.notificationINM111withErrors, counters.notificationINM111)
			log.Println("notificationZKKU: ", counters.notificationZKKUwithErrors, counters.notificationZKKU)
			log.Println("notificationZKK: ", counters.notificationZKKwithErrors, counters.notificationZKK)
			log.Println("notificationZA: ", counters.notificationZAwithErrors, counters.notificationZA)

			log.Println("contracts223: ", counters.contracts223)
			log.Println("Errors/Total:", counters.withErrors, counters.total)
			log.Println("Unsupported:", counters.unsupported)
	}
}
