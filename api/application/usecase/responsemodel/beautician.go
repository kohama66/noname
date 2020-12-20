package responsemodel

import (
	"time"
)

// Beautician response構造体
type Beautician struct {
	// ID            int64             `json:"id"`
	// RandID        string            `json:"randId"`
	// FirstName     string            `json:"firstName"`
	// LastName      string            `json:"lastName"`
	// FirstNameKana string            `json:"firstNameKana"`
	// LastNameKana  string            `json:"lastNameKana"`
	// PhoneNumber   string            `json:"phoneNumber"`
	LineID      *string `json:"lineId"`
	InstagramID *string `json:"instagramId"`
	Comment     *string `json:"comment"`
	// Menus     []*BeauticianMenu `json:"menus"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BeauticianCreate response構造体
type BeauticianCreate struct {
	*Beautician `json:"beautician"`
}

// BeauticianGet response構造体
type BeauticianGet struct {
	*Beautician `json:"beautician"`
}

// BeauticianFind response構造体
type BeauticianFind struct {
	Beauticians []*User `json:"beauticians"`
}
