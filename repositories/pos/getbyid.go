package pos

import (
	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
)

func (b *posRepository) GetById(req dto.GetByIdRequest) (*dto.PosUpdateResponse, error) {
	tr := entity.Pos{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.PosUpdateResponse{
		StatusPayment: tr.StatusPayment,
		UpdatedBy:     tr.StatusPayment,
	}

	return response, nil
}
