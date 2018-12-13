package common

import (
	"path/filepath"
	"os"
	"log"
	"flag"
	"time"
	"sort"
	"archive/zip"
	"strings"
	"bytes"
	"path"
	"fmt"
)

type ZipArray struct {
	PathZip		string
	NameZip		string
	ModTime		time.Time
}

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

func (p *ZipFile) Init(fileName string) error {
	var err error
	name := path.Base(strings.Replace(fileName, "\\", "/", -1))
	p.Name = name
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

func main() {

	qqqqq := fmt.Sprint("qwe", 23, "asd")
	fmt.Println(qqqqq)
	return


	var myflag string
	flag.StringVar(&myflag, "mydir", "", "path")

	flag.Parse()

	//	Получаем список Архивов
	var ZipList []ZipArray

	var cnt int64 = 0
	var cnt_all int64 = 0

	walk := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalln(err)
		}

		matched, err := filepath.Match("*.zip", info.Name())
		if err != nil {
			log.Fatalln(err)
		}
		cnt_all++
		if !info.IsDir() && matched && cnt < 1 {
			cnt++
			//log.Println(path, info.Size(), info.ModTime())

			ZipList = append(ZipList, ZipArray{path, info.Name(), info.ModTime()})

			//log.Println("From1:", info.Name(), !info.IsDir() && matched, len(ch))
		}

		return nil
	}
	filepath.Walk(myflag, walk)

	sort.Slice(ZipList, func(i, j int) bool {
		return ZipList[i].ModTime.Sub(ZipList[j].ModTime) < 0
	})




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
			log.Println("Trouble with zip file: ", val.PathZip, " ", err)
			continue
		}

		i := 0
		//j := 0
		if len(r.File) != 0 {
			for k, f := range r.File {
				if k > 0 { break }

				if r.File[k] == nil {
					panic("Empty XML file")
				}

				if strings.Contains(f.Name, ".sig") {
					log.Println(f.Name, "Skipping...")
					continue
				}
				x := new(XmlFile)

				x.Init(f.Name, z.Bucket(), z.Key(), val.PathZip)

				//xmlExist := false

				/*
				if flagIsCheck == 1 {
					xmlExist, err = HaveXml(x, db)
					if err != nil {
						log.Println(err)
					}
				}
				//log.Println("CHEK!", x, xmlExist)

				if xmlExist {
					j++
					if flagLinear {
						linear <- true
					}
				} else {
				//*/
					i++
					rc, err := f.Open()
					if err != nil {
						log.Fatal(err)
					}
					buf := new(bytes.Buffer)
					buf.ReadFrom(rc)
					rc.Close()
					x.Data = buf.Bytes()
					//z.Files = append(z.Files, x)

					///////////////////////////////ParserWorker(db *sql.DB, xmlName XmlFile, flagIsSave int)

					log.Println(x.Name, "\n\n\n", string(x.Data))

					/*
				}
				*/
			}

			//log.Println(val.PathZip, "New:", i, "Existed:", j)

			err = r.Close()
			if err != nil {
				log.Fatal(err)
			}
		}


	}

	log.Println("\n\nResult:", cnt, "/", cnt_all)
}

