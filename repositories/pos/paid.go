package pos

import (
	"errors"
	"fmt"

	auth "github.com/srv-cashpay/auth/entity"

	"github.com/srv-cashpay/pos/entity"
	"gorm.io/gorm"
)

func (r *posRepository) Paid(pos entity.Pos) (entity.Pos, error) {
	if err := r.DB.Create(&pos).Error; err != nil {
		return entity.Pos{}, err
	}
	return pos, nil
}

func (r *posRepository) GetByID(id string) (entity.Pos, error) {
	var pos entity.Pos
	if err := r.DB.Where("id = ?", id).First(&pos).Error; err != nil {
		return entity.Pos{}, err
	}
	return pos, nil
}

// func (r *posRepository) IsUserVerified(userID string) (bool, error) {
// 	var user auth.UserVerified
// 	// Ganti ini tergantung field relasi kamu: ID atau UserID
// 	if err := r.DB.First(&user, "user_id = ?", userID).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return false, fmt.Errorf("akun tidak ditemukan")
// 		}
// 		return false, err
// 	}
// 	return user.StatusAccount, nil
// }

func (r *posRepository) GetUserVerified(userID string) (auth.UserVerified, error) {
	var user auth.UserVerified
	if err := r.DB.First(&user, "user_id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, fmt.Errorf("akun tidak ditemukan")
		}
		return user, err
	}
	return user, nil
}
