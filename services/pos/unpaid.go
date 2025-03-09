package pos

import (
	"encoding/json"
	"fmt"

	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
	util "github.com/srv-cashpay/util/s"
)

func (s *posService) Unpaid(req dto.PosRequest) (dto.PosResponse, error) {
	// Konversi produk ke JSON
	productsJSON, err := json.Marshal(req.Product)
	if err != nil {
		return dto.PosResponse{}, err
	}

	// Hitung total harga
	totalPrice := calculateTotalPriceUnpaid(req.Product)

	// Buat struct Pos untuk disimpan
	pos := entity.Pos{
		ID:            util.GenerateRandomString(),
		UserID:        req.UserID,
		MerchantID:    req.MerchantID,
		StatusPayment: req.StatusPayment,
		Product:       productsJSON,
		CreatedBy:     req.CreatedBy,
		Description:   fmt.Sprintf("%s telah mendapatkan total harga penjualan sebesar %d", req.CreatedBy, totalPrice),
	}

	// Simpan ke database
	createdPos, err := s.Repo.Unpaid(pos)
	if err != nil {
		return dto.PosResponse{}, err
	}

	// Konversi produk untuk response
	var responseProducts []dto.ProductResponse
	if err := json.Unmarshal(productsJSON, &responseProducts); err != nil {
		return dto.PosResponse{}, err
	}

	return dto.PosResponse{
		ID:            createdPos.ID,
		UserID:        createdPos.UserID,
		MerchantID:    createdPos.MerchantID,
		StatusPayment: createdPos.StatusPayment,
		CreatedBy:     createdPos.CreatedBy,
		Product:       responseProducts,
		TotalPrice:    totalPrice,
		Description:   createdPos.Description,
	}, nil
}

func calculateTotalPriceUnpaid(products []dto.ProductRequest) int {
	total := 0
	for _, product := range products {
		price := product.Price

		total += price * product.Quantity
	}
	return total
}
