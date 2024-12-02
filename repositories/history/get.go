package history

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
)

func (r *historyRepository) Get(req dto.PaginationRequest) (dto.PaginationResponse, int) {
	var products []entity.Pos

	var totalRows int64
	totalPages, fromRow, toRow := 0, 0, 0

	// Ubah offset agar sesuai dengan page yang dimulai dari 1
	offset := (req.Page - 1) * req.Limit

	// Ambil data sesuai limit, offset, dan urutan
	find := r.DB.Where("user_id = ?", req.UserID).Limit(req.Limit).Offset(offset).Order(req.Sort)

	// Generate where query untuk search
	if req.Searchs != nil {
		for _, value := range req.Searchs {
			column := value.Column
			action := value.Action
			query := value.Query

			switch action {
			case "equals":
				find = find.Where(fmt.Sprintf("%s = ?", column), query)
			case "contains":
				find = find.Where(fmt.Sprintf("%s LIKE ?", column), "%"+query+"%")
			case "in":
				find = find.Where(fmt.Sprintf("%s IN (?)", column), strings.Split(query, ","))
			}
		}
	}

	find = find.Find(&products)

	// Periksa jika ada error saat pengambilan data
	if errFind := find.Error; errFind != nil {
		return dto.PaginationResponse{}, totalPages
	}

	if errCount := r.DB.Model(&entity.Pos{}).Where("user_id = ?", req.UserID).Count(&totalRows).Error; errCount != nil {
		return dto.PaginationResponse{}, totalPages
	}

	req.TotalRows = int(totalRows)

	// Hitung total halaman berdasarkan limit
	totalPages = int(math.Ceil(float64(totalRows) / float64(req.Limit)))
	req.TotalPages = totalPages
	// Hitung `fromRow` dan `toRow` untuk page saat ini
	if req.Page == 1 {
		// Untuk halaman pertama
		fromRow = 1
		toRow = req.Limit
	} else {
		if req.Page <= totalPages {
			fromRow = (req.Page-1)*req.Limit + 1
			toRow = req.Page * req.Limit
		}
	}

	// Pastikan `toRow` tidak melebihi `totalRows`
	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	// Set hasil akhir
	req.FromRow = fromRow
	req.ToRow = toRow
	// req.Rows = products
	var posResponses []dto.PosResponse
	for _, pos := range products {
		var productResponses []dto.ProductResponse
		if err := json.Unmarshal(pos.Product, &productResponses); err != nil {
			continue
		}
		posResponse := dto.PosResponse{
			ID:         pos.ID,
			UserID:     pos.UserID,
			CreatedBy:  pos.CreatedBy,
			Product:    productResponses,
			TotalPrice: calculateTotalPrice(productResponses),
		}
		posResponses = append(posResponses, posResponse)
	}

	response := dto.PaginationResponse{
		Limit:        req.Limit,
		Page:         req.Page,
		Sort:         req.Sort,
		TotalRows:    req.TotalRows,
		TotalPages:   req.TotalPages,
		FirstPage:    req.FirstPage,
		PreviousPage: req.PreviousPage,
		NextPage:     req.NextPage,
		LastPage:     req.LastPage,
		FromRow:      req.FromRow,
		ToRow:        req.ToRow,
		Data:         posResponses,
		Searchs:      req.Searchs,
	}

	return response, totalPages
}

func calculateTotalPrice(products []dto.ProductResponse) int {
	total := 0
	for _, product := range products {
		total += product.Price * product.Quantity
	}
	return total
}
