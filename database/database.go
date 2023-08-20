package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Document{}) // AutoMigrate to create tables
	return &Database{DB: db}
}

func (db *Database) SaveDocument(doc *Document) error {
	result := db.DB.Create(doc)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *Database) FindDocumentByID(id uint, token string) (*Document, error) {
	var document Document
	result := db.DB.First(&document, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &document, nil
}

