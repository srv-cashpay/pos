package pos

import (
	"github.com/srv-cashpay/pos/entity"
)

func (r *posRepository) Unpaid(pos entity.Pos) (entity.Pos, error) {
	if err := r.DB.Create(&pos).Error; err != nil {
		return entity.Pos{}, err
	}
	return pos, nil
}
