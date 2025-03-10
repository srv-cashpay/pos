package pos

import (
	"encoding/json"
	"fmt"

	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
	util "github.com/srv-cashpay/util/s"
)

func (s *posService) Paid(req dto.PosRequest) (dto.PosResponse, error) {
	// Validasi daftar produk tidak boleh kosong
	if len(req.Product) == 0 {
		return dto.PosResponse{}, fmt.Errorf("daftar produk tidak boleh kosong")
	}

	// Hitung total harga produk
	totalPrice := calculateTotalPrice(req.Product)

	// Validasi apakah `pay` cukup untuk membayar total harga
	if req.Pay < totalPrice {
		return dto.PosResponse{}, fmt.Errorf("jumlah pembayaran tidak cukup, total harga: %d, pembayaran: %d", totalPrice, req.Pay)
	}

	// Konversi produk ke JSON
	productsJSON, err := json.Marshal(req.Product)
	if err != nil {
		return dto.PosResponse{}, fmt.Errorf("gagal mengonversi produk ke JSON: %w", err)
	}

	// Buat struct Pos untuk disimpan
	pos := entity.Pos{
		ID:            util.GenerateRandomString(),
		UserID:        req.UserID,
		MerchantID:    req.MerchantID,
		StatusPayment: req.StatusPayment,
		Product:       productsJSON,
		Pay:           req.Pay,
		CreatedBy:     req.CreatedBy,
		Description:   "Terima Kasih",
	}

	// Simpan ke database
	createdPos, err := s.Repo.Paid(pos)
	if err != nil {
		return dto.PosResponse{}, fmt.Errorf("gagal menyimpan transaksi: %w", err)
	}

	// Konversi produk untuk response
	var responseProducts []dto.ProductResponse
	if err := json.Unmarshal(productsJSON, &responseProducts); err != nil {
		return dto.PosResponse{}, fmt.Errorf("gagal mengurai JSON produk untuk response: %w", err)
	}

	// Return response yang lebih lengkap
	return dto.PosResponse{
		ID:            createdPos.ID,
		UserID:        createdPos.UserID,
		MerchantID:    createdPos.MerchantID,
		StatusPayment: createdPos.StatusPayment,
		CreatedBy:     createdPos.CreatedBy,
		Product:       responseProducts,
		Pay:           createdPos.Pay,
		TotalPrice:    totalPrice,
		Description:   createdPos.Description,
	}, nil
}

// Fungsi untuk menghitung total harga produk
func calculateTotalPrice(products []dto.ProductRequest) int {
	total := 0
	for _, product := range products {
		total += product.Price * product.Quantity
	}
	return total
}
