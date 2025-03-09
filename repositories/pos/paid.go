package pos

import (
	"github.com/srv-cashpay/pos/entity"
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
