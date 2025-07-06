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

	// Ambil data POS + relasi Merchant, Discount, Tax
	if err := b.DB.Where("id = ?", tr.ID).
		Preload("Merchant").
		Preload("Discount").
		Preload("Tax").
		Take(&tr).Error; err != nil {
		return nil, err
	}

	// Unmarshal product JSON
	var products []dto.ProductResponse
	if err := json.Unmarshal(tr.Product, &products); err != nil {
		return nil, err
	}

	// Siapkan discount
	var discounts []dto.DiscountResponse
	var discountPercents []uint
	for _, d := range tr.Discount {
		discounts = append(discounts, dto.DiscountResponse{
			DiscountPercentage: d.DiscountPercentage,
		})
		discountPercents = append(discountPercents, d.DiscountPercentage)
	}

	// Siapkan tax
	var taxs []dto.TaxResponse
	var taxPercents []uint
	for _, t := range tr.Tax {
		taxs = append(taxs, dto.TaxResponse{
			TaxPercentage: t.TaxPercentage,
		})
		taxPercents = append(taxPercents, t.TaxPercentage)
	}

	// Hitung harga
	totalPrice := calculateTotalPrice(products)
	discountedTotal := calculateDiscountedTotal(totalPrice, discountPercents)

	// Hitung tax total
	var totalTaxPercent uint
	for _, t := range taxPercents {
		totalTaxPercent += t
	}
	if totalTaxPercent > 100 {
		totalTaxPercent = 100
	}

	taxAmount := calculateTax(discountedTotal, totalTaxPercent)
	totalWithTax := discountedTotal + taxAmount

	// Bangun response
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
		Tax:                taxAmount,
		TotalPrice:         totalPrice,
		TotalAfterDiscount: discountedTotal,
		TotalWithTax:       totalWithTax,
		Pay:                tr.Pay,
		Change:             tr.Pay - totalWithTax,
		Description:        tr.Description,
	}

	return response, nil
}

// Hitung harga total sebelum diskon
func calculateTotalPrice(products []dto.ProductResponse) int {
	total := 0
	for _, p := range products {
		total += p.Price * p.Quantity
	}
	return total
}

// Hitung harga setelah diskon (ambil diskon terbesar)
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
	discounted := total - (total * int(maxDiscount) / 100)
	return discounted
}

// Hitung tax
func calculateTax(total int, taxPercent uint) int {
	return total * int(taxPercent) / 100
}
