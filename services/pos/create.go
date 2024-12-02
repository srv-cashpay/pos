package pos

import (
	"encoding/json"

	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
	util "github.com/srv-cashpay/util/s"
)

func (s *posService) Create(req dto.PosRequest) (dto.PosResponse, error) {
	// Konversi produk ke JSON
	productsJSON, err := json.Marshal(req.Product)
	if err != nil {
		return dto.PosResponse{}, err
	}

	// Hitung total harga
	totalPrice := calculateTotalPrice(req.Product)

	// Buat struct Pos untuk disimpan
	pos := entity.Pos{
		ID:         util.GenerateRandomString(),
		UserID:     req.UserID,
		MerchantID: req.MerchantID,
		Product:    productsJSON,
		CreatedBy:  req.CreatedBy,
	}

	// Simpan ke database
	createdPos, err := s.Repo.Create(pos)
	if err != nil {
		return dto.PosResponse{}, err
	}

	// Konversi produk untuk response
	var responseProducts []dto.ProductResponse
	if err := json.Unmarshal(productsJSON, &responseProducts); err != nil {
		return dto.PosResponse{}, err
	}

	return dto.PosResponse{
		ID:         createdPos.ID,
		UserID:     createdPos.UserID,
		MerchantID: createdPos.MerchantID,
		CreatedBy:  createdPos.CreatedBy,
		Product:    responseProducts,
		TotalPrice: totalPrice,
	}, nil
}

func calculateTotalPrice(products []dto.ProductRequest) int {
	total := 0
	for _, product := range products {
		price := product.Price

		total += price * product.Quantity
	}
	return total
}
