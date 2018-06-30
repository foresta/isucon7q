package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func init() {
	db_host := os.Getenv("ISUBATA_DB_HOST")
	if db_host == "" {
		db_host = "127.0.0.1"
	}
	db_port := os.Getenv("ISUBATA_DB_PORT")
	if db_port == "" {
		db_port = "3306"
	}
	db_user := os.Getenv("ISUBATA_DB_USER")
	if db_user == "" {
		db_user = "root"
	}
	db_password := os.Getenv("ISUBATA_DB_PASSWORD")
	if db_password != "" {
		db_password = ":" + db_password
	}

	dsn := fmt.Sprintf("%s%s@tcp(%s:%s)/isubata?parseTime=true&loc=Local&charset=utf8mb4",
		db_user, db_password, db_host, db_port)

	log.Printf("Connecting to db: %q", dsn)
	db, _ = sqlx.Connect("mysql", dsn)
	for {
		err := db.Ping()
		if err == nil {
			break
		}
		log.Println(err)
		time.Sleep(time.Second * 3)
	}

	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Printf("Succeeded to connect db.")

}

func main() {

	filepath := "/home/isucon/isubata/webapp/public/icons/"

	var filename string
	var data []byte

	// DBからアイコンを抽出
	rows, err := db.Query("SELECT name, data FROM image")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&filename, &data)
		if err != nil {
			continue
		}
		log.Printf("file %s prosessing...", filename)

		file, err := os.Create(filepath + filename)
		if err != nil {
			continue
		}
		defer file.Close()

		buffer := bytes.NewBuffer(data)
		io.Copy(file, buffer)
	}
}
