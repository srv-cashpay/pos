package history

import (
	dto "github.com/srv-cashpay/pos/dto"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.PaginationRequest) (dto.PaginationResponse, int)
	GetById(req dto.GetByIdRequest) (*dto.PosResponse, error)
}

type historyRepository struct {
	DB *gorm.DB
}

func NewHistoryRepository(DB *gorm.DB) DomainRepository {
	return &historyRepository{
		DB: DB,
	}
}
