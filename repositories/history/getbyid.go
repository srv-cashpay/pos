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

	// Mengambil data Pos sekaligus memuat relasi Merchant
	if err := b.DB.Where("id = ?", tr.ID).Preload("Merchant").Preload("Discount").Take(&tr).Error; err != nil {
		return nil, err
	}

	var products []dto.ProductResponse
	if err := json.Unmarshal(tr.Product, &products); err != nil {
		return nil, err
	}

	var discounts []dto.DiscountResponse
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
		MerchantName:  tr.Merchant.MerchantName,
		Address:       tr.Merchant.Address,
		City:          tr.Merchant.City,
		Country:       tr.Merchant.Country,
		TotalPrice:    calculateTotalPrice(products),
		Discount:      discounts,
		Pay:           tr.Pay,
		Change:        tr.Pay - calculateTotalPrice(products),
		Description:   tr.Description,
	}

	return response, nil
}
