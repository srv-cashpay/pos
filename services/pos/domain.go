package pos

import (
	m "github.com/srv-cashpay/middlewares/middlewares"
	dto "github.com/srv-cashpay/pos/dto"

	r "github.com/srv-cashpay/pos/repositories/pos"
)

type PosService interface {
	Create(req dto.PosRequest) (dto.PosResponse, error)
	GetById(req dto.GetByIdRequest) (*dto.PosUpdateResponse, error)
	Update(req dto.PosUpdateRequest) (dto.PosUpdateResponse, error)
}

type posService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPosService(Repo r.DomainRepository, jwtS m.JWTService) PosService {
	return &posService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
