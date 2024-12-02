package history

import (
	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
)

func (b *historyRepository) GetById(req dto.GetByIdRequest) (*dto.PosResponse, error) {
	tr := entity.Pos{
		ID: req.ID,
	}

	if err := b.DB.Where("id = ?", tr.ID).Take(&tr).Error; err != nil {
		return nil, err
	}

	response := &dto.PosResponse{
		UserID: tr.UserID,
	}

	return response, nil
}
