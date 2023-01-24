package models

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//connection user
func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(192.168.138.139:3306)/fathanah"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	DB = db
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(192.168.138.139:3306)/fathanah")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// func ConnectDatabase() {
// 	db, err := gorm.Open(mysql.Open("root:@tcp(192.168.138.139:3306)/mydb"))
// 	if err != nil {
// 		fmt.Println("Gagal koneksi database")
// 	}

// 	DB = db
// }

//192.168.138.139
//localhost
