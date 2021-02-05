package entityx

import "github.com/myapp/noname/api/domain/entity"

// NewBeauticianSalon ファクトリ関数
func NewBeauticianSalon(beauticianID, salonID int64) *entity.BeauticianSalon {
	return &entity.BeauticianSalon{
		BeauticianID: beauticianID,
		SalonID:      salonID,
	}
}
