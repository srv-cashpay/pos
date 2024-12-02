package history

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/pos/helpers"
	res "github.com/srv-cashpay/util/s/response"
)

func (b *domainHandler) Get(c echo.Context) error {
	paginationDTO := helpers.GeneratePaginationRequest(c)

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	paginationDTO.UserID = userid

	err := c.Bind(&paginationDTO)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	users := b.serviceHistory.Get(c, paginationDTO)

	return res.SuccessResponse(users).Send(c)
}
