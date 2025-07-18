package dto

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type GetByIdRequest struct {
	ID     string `param:"id" validate:"required"`
	UserID string `json:"user_id"`
}

type PosRequest struct {
	ID            string           `json:"id"`
	UserID        string           `json:"user_id"`
	MerchantID    string           `json:"merchant_id"`
	StatusPayment string           `json:"status_payment"`
	Quantity      int              `json:"quantity"`
	CreatedBy     string           `json:"created_by"`
	Product       []ProductRequest `json:"product"`
	DiscountApply uint             `json:"discount_apply"`
	TaxApply      uint             `json:"tax_apply"`
	Pay           int              `json:"pay"`
	Description   string           `json:"description"`
}

type ProductRequest struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type PosResponse struct {
	ID                 string             `json:"id"`
	UserID             string             `json:"user_id"`
	StatusPayment      string             `json:"status_payment"`
	MerchantID         string             `json:"merchant_id"`
	MerchantName       string             `json:"merchant_name"`
	Address            string             `json:"address"`
	City               string             `json:"city"`
	Country            string             `json:"country"`
	CreatedBy          string             `json:"created_by"`
	Product            []ProductResponse  `json:"product"`
	Discount           []DiscountResponse `json:"discount"`
	Tax                []TaxResponse      `json:"tax"`
	DiscountApply      uint               `json:"discount_apply"`
	TaxApply           uint               `json:"tax_apply"`
	TaxAmount          int                `json:"tax_amount"`
	TotalPrice         int                `json:"total_price"`
	TotalAfterDiscount int                `json:"total_after_discount"`
	TotalWithTax       int                `json:"total_with_tax"`
	Pay                int                `json:"pay"`
	Change             int                `json:"change"`
	Description        string             `json:"description"`
	Account            AccountResponse    `json:"account"`
}
type AccountResponse struct {
	StatusAccount  bool      `json:"status_account"`
	AccountExpired time.Time `json:"account_expired"`
}

type DiscountResponse struct {
	DiscountName       string `json:"discount_name"`
	DiscountPercentage uint   `json:"discount_percentage"`
}

type TaxResponse struct {
	Tax           string `json:"tax"`
	TaxPercentage uint   `json:"tax_percentage"`
}

type RequirementRequest struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	MerchantID string `json:"merchant_id"`
}
type RequirementResponse struct {
	Tax      []TaxResponse      `json:"tax_percentage"`
	Discount []DiscountResponse `json:"discount_percentage"`
}

type ProductResponse struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type PosUpdateRequest struct {
	ID            string `json:"id"`
	StatusPayment string `json:"status_payment"`
	UpdatedBy     string `json:"updated_by"`
	Description   string `json:"description"`
}

type PosUpdateResponse struct {
	ID            string `json:"id"`
	StatusPayment string `json:"status_payment"`
	UpdatedBy     string `json:"updated_by"`
	Description   string `json:"description"`
}

func (r PosResponse) MarshalJSON() ([]byte, error) {
	type Alias PosResponse
	return json.Marshal(&struct {
		*Alias
		TotalPrice         string `json:"total_price"`
		TotalAfterDiscount string `json:"total_after_discount"`
		TotalWithTax       string `json:"total_with_tax"`
		Pay                string `json:"pay"`
		Change             string `json:"change"`
	}{
		Alias:              (*Alias)(&r),
		TotalPrice:         formatRupiah(r.TotalPrice),
		TotalAfterDiscount: formatRupiah(r.TotalAfterDiscount),
		TotalWithTax:       formatRupiah(r.TotalWithTax),
		Pay:                formatRupiah(r.Pay),
		Change:             formatRupiah(r.Change),
	})
}

func formatRupiah(amount int) string {
	s := strconv.Itoa(amount)
	var result strings.Builder
	length := len(s)
	count := 0

	for i := length - 1; i >= 0; i-- {
		if count > 0 && count%3 == 0 {
			result.WriteString(".")
		}
		result.WriteByte(s[i])
		count++
	}

	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return "Rp " + string(runes)
}
