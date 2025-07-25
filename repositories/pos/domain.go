package pos

import (
	auth "github.com/srv-cashpay/auth/entity"
	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Paid(req entity.Pos) (entity.Pos, error)
	Unpaid(pos entity.Pos) (entity.Pos, error)
	Update(req dto.PosUpdateRequest) (dto.PosUpdateResponse, error)
	GetById(req dto.GetByIdRequest) (*dto.PosUpdateResponse, error)
	GetUserVerified(userID string) (auth.UserVerified, error)
	Requirement(req dto.RequirementRequest) (dto.RequirementResponse, error)
}

type posRepository struct {
	DB *gorm.DB
}

func NewPosRepository(DB *gorm.DB) DomainRepository {
	return &posRepository{
		DB: DB,
	}
}
