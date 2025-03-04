package history

import (
	"encoding/json"

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

	var products []dto.ProductResponse
	if err := json.Unmarshal(tr.Product, &products); err != nil {
		return nil, err
	}

	response := &dto.PosResponse{
		ID:            tr.ID,
		UserID:        tr.UserID,
		StatusPayment: tr.StatusPayment,
		MerchantID:    tr.MerchantID,
		CreatedBy:     tr.CreatedBy,
		Product:       products,
		TotalPrice:    calculateTotalPrice(products),
	}

	return response, nil
}
