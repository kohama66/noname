package responsemodel

import "time"

// Salon response構造体
type Salon struct {
	ID           int64     `json:"id"`
	RandID       string    `json:"randId"`
	Name         string    `json:"name"`
	PhoneNumber  string    `json:"phoneNumber"`
	OpeningHours string    `json:"openingHours"`
	ClosingHours string    `json:"closingHours"`
	PostalCode   string    `json:"postalCode"`
	Prefectures  string    `json:"prefectures"`
	City         string    `json:"city"`
	Town         string    `json:"town"`
	AddressCode  string    `json:"addressCode"`
	AddressOther string    `json:"addressOther"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// SalonFind 美容院検索 response構造体
type SalonFind struct {
	Salons []*Salon `json:"salons"`
}
