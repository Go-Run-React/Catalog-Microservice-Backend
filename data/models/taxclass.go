package models

import (
	"gorm.io/gorm"
)

type TaxClass struct {
	gorm.Model
	Type string
	Code string
	Rate int8
	Description string
}

func(r *Repository) MigrateTaxClass() {
	err := r.DB.AutoMigrate(&TaxClass{})
	if err != nil {
		r.l.Fatal("Failed to migrate Tax Class")
	}
}
var tc TaxClass
var tclist []TaxClass

func(r *Repository) GetTaxClasses() []TaxClass {
	result := r.DB.Find(&tclist)
	if result.Error !=nil {
		r.l.Printf("TaxClass list fetch error: %s", result.Error)
		return nil
	}
	return tclist
}

func(r *Repository) AddTaxClass(tc *TaxClass) {
	result := r.DB.Create(tc)
	if result.Error !=nil {
		r.l.Printf("TaxClass creation failed: %s", result.Error)
	} else {
		r.l.Printf("TaxClass successfully created with id: %d", tc.ID)
	}
}


