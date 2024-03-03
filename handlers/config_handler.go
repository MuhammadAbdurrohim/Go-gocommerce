package handlers

import (
	"log"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	log.Printf("Connecting to database...") // Menggunakan log.Printf untuk mencetak pesan log
	db, err = gorm.Open("mysql", "root@localhost:root@tcp(127.0.0.1:3306)/dbgorm")


	if err != nil {
		log.Fatalf("Failed to connect database: %v", err) // Menggunakan log.Fatalf untuk mencetak pesan log dan keluar dari program jika terjadi kesalahan
	}

	log.Printf("Connected to database successfully.")
}