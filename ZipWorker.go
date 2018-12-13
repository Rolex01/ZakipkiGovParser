package main

import (
	"archive/zip"
	"bytes"
	//	"fmt"
	"log"
	"path"
	//"regexp"
	//"runtime"
	//"runtime/debug"
	"strings"
	"database/sql"
)

// Archive describes zip file
// Period ...
// Num ...
// Type ...
// Tag ...
// Files ...
// <вид-документа_регион_начало-периода_конец-периода_дата-время-выгрузки_nnn.xml.zip>
type ZipFile struct {
	Period struct {
		Start string
		End   string
	}
	Num        string
	Name       string
	Type       string
	Region     string
	UploadDate string
	Files      []*XmlFile
}

const (
	manualType string = "[*]*_[0-9]{14}_[0-9]{1,4}.xml"
	autoType   string = "[*]*_[0-9]{10}_[0-9]{10}_[0-9]{1,4}.xml"
)

func (p *ZipFile) Init(fileName string) error {
	var (
		//matched bool
		err error
	)
	name := path.Base(strings.Replace(fileName, "\\", "/", -1))

	/*matched, err = regexp.MatchString(autoType, name)
	if err != nil {
		log.Println(err)
	}*/
	p.Name = name
	/*
		if matched {
			// Auto
			nameParts := strings.Split(name, "_")
			l := len(nameParts) - 1
			p.Type = nameParts[0]
			tmp := strings.Split(nameParts[l], ".")
			p.Num = tmp[0]
			p.Period.Start = nameParts[l-2]
			p.Period.End = nameParts[l-1]
			for i := 1; i < l-2; i++ {
				p.Region += nameParts[i]
			}
		} else {
			matched, err = regexp.MatchString(manualType, name)
			if err != nil {
				log.Println(err)
			}
			if matched {
				// Manual
				nameParts := strings.Split(name, "_")
				l := len(nameParts) - 1
				p.Type = nameParts[0]
				tmp := strings.Split(nameParts[l], ".")
				p.Num = tmp[0]
				p.Period.Start = nameParts[l-3]
				p.UploadDate = nameParts[l-2]
				p.Period.End = nameParts[l-1]
				for i := 1; i < l-3; i++ {
					p.Region += nameParts[i]
				}
			} else {
				err = fmt.Errorf("Not supported archive filename (%q)!", name)
				log.Println(err)
				// ? will it pass over the condition operator WARNING
			}
		}*/
	return err
}

func (p *ZipFile) Key() []byte {
	//return []byte("zip" + p.Period.Start + p.Period.End + p.Num)
	return []byte(p.Name)
}

func (p *ZipFile) Bucket() []byte {
	//return []byte("zip" + p.Type + p.Region)
	return []byte(p.Name)
}

// XmlFile •	document_regNum_documentId.xml
type XmlFile struct {
	RegNum string
	DocId  string
	Data   []byte
	Type   string
	Zip    string
	Tag    string
	Period string
	Name   string
	//error  error
}

func (p *XmlFile) Init(fileName string, tag []byte, val []byte, zip string) {
	p.Name = fileName
	p.Zip = zip
	name := strings.Split(fileName, "_")
	if len(name) < 2 {

	} else {
		p.RegNum = name[1]
		p.DocId = name[2]
		p.Type = name[0]
		p.Tag = string(tag)
		p.Period = string(val)
	}
}
func (p *XmlFile) LoadData(data []byte) error { return nil }

func (p *XmlFile) Key() []byte {
	//return []byte(p.RegNum)
	return []byte(p.Name)
}
func (p *XmlFile) Val() []byte {
	//return []byte(p.Period)
	return []byte(p.Name)
}

func (p *XmlFile) Bucket() []byte {
	return []byte(p.Tag)
}

func ZipWorker(xmlFile chan *XmlFile, db *sql.DB, flagIsCheck int) error {
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
						continue
					}

					buf := new(bytes.Buffer)
					buf.ReadFrom(rc)
					rc.Close()
					x.Data = buf.Bytes()

					xmlFile <- x
					//err = ParserWorker(db, x, flagIsSave, &counters)
					/*
					if err != nil {
						log.Fatal(err)
					} else {
						log.Println("GOOD! ZIP:", val.NameZip, "; XML:", f.Name)
					}
					*/
					//log.Println(x.Name, "\n\n\n", string(x.Data))
				}
			}

			err = r.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	}











	/*
	log.Println("ZipWorker :: Start | ", flagLinear)
	if flagLinear == false {
		sc := new(sync.WaitGroup)
		for file := range files {
			file := file // Create new instance of file for the goroutine.
			log.Println("ZipWorker :: Open (", file, ")")
			sc.Add(1)
			pool <- 1
			go func(sc *sync.WaitGroup, file string) {
				unzip(file, archives, db, flagLinear, linear, flagIsCheck)
				//put info into database
				<-pool
				sc.Done()
			}(sc, file)
		}
		sc.Wait()
		//done <- true
		//unzip(file, archives, db)

		//}
		//ToDo process errors
		close(archives)
		//close(fopens)
		//close(pool)
	} else {
		log.Println("ZipWorker :: Linear Mode")
		for file := range files {
			file := file // Create new instance of file for the goroutine.
			log.Println("ZipWorker :: Open (", file, ")")
			go unzip(file, archives, db, flagLinear, linear, flagIsCheck)
			log.Println("ZipWorker :: Unziped : ", file)
		WaitLoop:
			for {
				select {
				case <-linear:
					break WaitLoop
				default:
					time.Sleep(time.Microsecond)
				}
			}
		}
	}
	log.Println("ZipWorker :: End")
	return nil
	*/
	return nil
}
