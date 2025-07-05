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
	ID            string             `json:"id"`
	UserID        string             `json:"user_id"`
	StatusPayment string             `json:"status_payment"`
	MerchantID    string             `json:"merchant_id"`
	MerchantName  string             `json:"merchant_name"`
	Discount      []DiscountResponse `json:"discount"`
	Address       string             `json:"address"`
	Country       string             `json:"country"`
	City          string             `json:"city"`
	CreatedBy     string             `json:"created_by"`
	Quantity      int                `json:"-"`
	Product       []ProductResponse  `json:"product"`
	TotalPrice    int                `json:"total_price"`
	Pay           int                `json:"pay"`
	Change        int                `json:"change"`
	Description   string             `json:"description"`
	Account       AccountResponse    `json:"account"`
}
type AccountResponse struct {
	StatusAccount  bool      `json:"status_account"`
	AccountExpired time.Time `json:"account_expired"`
}

type DiscountResponse struct {
	DiscountPercentage uint `json:"discount_percentage"`
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
		TotalPrice string `json:"total_price"`
		Pay        string `json:"pay"`
		Change     string `json:"change"`
	}{
		Alias:      (*Alias)(&r),
		TotalPrice: formatRupiah(r.TotalPrice),
		Pay:        formatRupiah(r.Pay),
		Change:     formatRupiah(r.Change),
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
