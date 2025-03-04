package history

import (
	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/pos/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) GetById(c echo.Context) error {
	var req dto.GetByIdRequest

	// idStr := c.Param("id")

	idUint, err := res.IsNumber(c, "id")
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	req.ID = idUint

	transaction, err := b.serviceHistory.GetById(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.NotFound, err).Send(c)

	}

	return res.SuccessResponse(transaction).Send(c)

}
