package database

type Document struct {
	ID           uint   `gorm:"primaryKey"`
	Title        string
	// Tags         []string `gorm:"type:varchar(64)[]"`
	HashedToken  string
	ConverterType string
}
