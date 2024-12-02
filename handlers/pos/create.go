package pos

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/pos/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) Create(c echo.Context) error {
	var req dto.PosRequest
	var resp dto.PosResponse

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	createdBy, ok := c.Get("CreatedBy").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.MerchantID = merchantId
	req.UserID = userid
	req.CreatedBy = createdBy

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	resp, err = h.servicePos.Create(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)

}
