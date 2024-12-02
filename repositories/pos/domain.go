package pos

import (
	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Create(req entity.Pos) (entity.Pos, error)
	Update(req dto.PosUpdateRequest) (dto.PosUpdateResponse, error)
	GetById(req dto.GetByIdRequest) (*dto.PosUpdateResponse, error)
}

type posRepository struct {
	DB *gorm.DB
}

func NewPosRepository(DB *gorm.DB) DomainRepository {
	return &posRepository{
		DB: DB,
	}
}
