package entityx

import "github.com/myapp/noname/api/domain/entity"

// NewBeauticianSalon ファクトリ関数
func NewBeauticianSalon(beauticianID, salonID int64) *entity.BeauticianSalon {
	return &entity.BeauticianSalon{
		BeauticianID: beauticianID,
		SalonID:      salonID,
	}
}

// NewUserSalons ファクトリ関数
func NewUserSalons(salon, user int64, role string) *entity.UserSalon {
	return &entity.UserSalon{
		SalonID: salon,
		UserID:  user,
		Role:    role,
	}
}

// UserSalonRoleEnum userSalons Role enum構造体
type UserSalonRoleEnum struct {
	value string
}

// UserSalonRoles userSalons Role enum実体
var UserSalonRoles = struct {
	Admin UserSalonRoleEnum
}{
	Admin: UserSalonRoleEnum{"admin"},
}

func (u UserSalonRoleEnum) String() string {
	if u.value == "" {
		return ""
	}
	return u.value
}
