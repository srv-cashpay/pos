package history

import (
	"github.com/labstack/echo/v4"
	m "github.com/srv-cashpay/middlewares/middlewares"
	dto "github.com/srv-cashpay/pos/dto"

	r "github.com/srv-cashpay/pos/repositories/history"
)

type HistoryService interface {
	Get(context echo.Context, req dto.PaginationRequest) dto.PaginationResponse
	GetById(req dto.GetByIdRequest) (*dto.PosResponse, error)
}

type historyService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewHistoryService(Repo r.DomainRepository, jwtS m.JWTService) HistoryService {
	return &historyService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
