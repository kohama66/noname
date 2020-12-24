package responsemodel

import (
	"time"
)

// Beautician response構造体
type Beautician struct {
	LineID      *string   `json:"lineId"`
	InstagramID *string   `json:"instagramId"`
	Comment     *string   `json:"comment"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
	Beauticians []*User `json:"users"`
}
