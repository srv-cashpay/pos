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

	if err := b.DB.Where("id = ?", tr.ID).
		Preload("Merchant").
		Preload("Discount").
		Preload("Tax").
		Take(&tr).Error; err != nil {
		return nil, err
	}

	var products []dto.ProductResponse
	if err := json.Unmarshal(tr.Product, &products); err != nil {
		return nil, err
	}

	var discounts []dto.DiscountResponse
	var discountPercents []uint
	for _, d := range tr.Discount {
		discounts = append(discounts, dto.DiscountResponse{
			DiscountName:       d.DiscountName,
			DiscountPercentage: d.DiscountPercentage,
		})
		discountPercents = append(discountPercents, d.DiscountPercentage)
	}

	var taxs []dto.TaxResponse
	var taxPercents []uint
	for _, t := range tr.Tax {
		taxs = append(taxs, dto.TaxResponse{
			Tax:           t.Tax,
			TaxPercentage: t.TaxPercentage,
		})
		taxPercents = append(taxPercents, t.TaxPercentage)
	}

	// Hitung harga
	totalPrice := calculateTotalPrice(products)
	totalAfterDiscount := calculateDiscountedTotal(totalPrice, discountPercents)
	taxAmount := calculateTaxAmount(totalAfterDiscount, taxPercents)
	totalFinal := totalAfterDiscount + taxAmount

	response := &dto.PosResponse{
		ID:                 tr.ID,
		UserID:             tr.UserID,
		StatusPayment:      tr.StatusPayment,
		MerchantID:         tr.MerchantID,
		MerchantName:       tr.Merchant.MerchantName,
		Address:            tr.Merchant.Address,
		City:               tr.Merchant.City,
		Country:            tr.Merchant.Country,
		CreatedBy:          tr.CreatedBy,
		Product:            products,
		Discount:           discounts,
		Tax:                taxs,
		TaxAmount:          taxAmount,
		TotalPrice:         totalPrice,
		TotalAfterDiscount: totalAfterDiscount,
		TotalWithTax:       totalFinal,
		Pay:                tr.Pay,
		Change:             tr.Pay - totalFinal,
		Description:        tr.Description,
	}

	return response, nil
}

func calculateTotalPrice(products []dto.ProductResponse) int {
	total := 0
	for _, p := range products {
		total += p.Price * p.Quantity
	}
	return total
}

func calculateDiscountedTotal(total int, discounts []uint) int {
	var maxDiscount uint
	for _, d := range discounts {
		if d > maxDiscount {
			maxDiscount = d
		}
	}
	if maxDiscount > 100 {
		maxDiscount = 100
	}
	return total - (total * int(maxDiscount) / 100)
}

func calculateTaxAmount(total int, taxPercents []uint) int {
	var totalTaxPercent uint
	for _, t := range taxPercents {
		totalTaxPercent += t
	}
	if totalTaxPercent > 100 {
		totalTaxPercent = 100
	}
	return total * int(totalTaxPercent) / 100
}
