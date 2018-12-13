package main

import (
	"log"
	"database/sql"
)

func HaveXml(x *XmlFile, db *sql.DB) (bool, error) {
	collection := string(x.Bucket())
	key := string(x.Key())
	check := false

	/*
	log.Println("z.Bucket()", string(x.Bucket()))
	log.Println("z.Key()", string(x.Key()))
	log.Println("z.Name", x.Name)
	log.Println("z.Type", x.Type)
	log.Println("z.RegNum", x.RegNum)
	log.Println("z.Val()", string(x.Val()))
	os.Exit(123)
	//*/

	var res bool
	err := db.QueryRow(`
		SELECT ARRAY[$1] <@ xmlnames
		FROM ftp_xmlnames
		WHERE zipname = $2
	`, key, collection).Scan(&res)

	println("File:", collection, key, "; Res:", res, err)

	if err != nil {
		log.Println("HaveXml:   ", collection, " ", key, " doesnt exist;", "res:", res, "err:", err)
		if err.Error() == "sql: no rows in result set" {
			err = nil
		}
	}

	check = res

	return check, err
}

func SaveXml(x *XmlFile, db *sql.DB) error {
	collection := string(x.Bucket())
	key := string(x.Key())

	_, err := db.Exec(`
		INSERT INTO ftp_xmlnames (zipname, typename) VALUES ($1, $2)
		ON CONFLICT (zipname) DO NOTHING
	`, collection, x.Type)

	if err != nil {
		log.Fatal("SaveXml INSERT:   ", err)
		return err
	}

	_, err = db.Exec(`
		UPDATE ftp_xmlnames
		SET xmlnames = array_append(xmlnames, t.val),
    		update_date = now()
		FROM (VALUES($1)) AS t(val)
		WHERE zipname = $2
  			AND NOT ARRAY[t.val] <@ xmlnames
	`, key, collection)

	if err != nil {
		log.Fatal("SaveXml UPDATE:   ", err)
		return err
	}

	log.Println("Saved", collection, x.Type, key, string(x.Val()))
	return err
}
