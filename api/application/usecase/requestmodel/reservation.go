package requestmodel

// ReservationCreate 予約構造体
type ReservationCreate struct {
	AuthID       string `json:"-"`
	BeauticianID int64  `json:"beauticiaId"`
	SpaceID      int64  `json:"spaceId"`
	MenuID       int64  `json:"menuId"`
	Date         string `json:"date"`
	Time         string `json:"time"`
}
