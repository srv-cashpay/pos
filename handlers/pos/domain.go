package pos

import (
	s "github.com/srv-cashpay/pos/services/pos"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Paid(c echo.Context) error
	Unpaid(c echo.Context) error
	Update(c echo.Context) error
	GetById(c echo.Context) error
}

type domainHandler struct {
	servicePos s.PosService
}

func NewPosHandler(service s.PosService) DomainHandler {
	return &domainHandler{
		servicePos: service,
	}
}
