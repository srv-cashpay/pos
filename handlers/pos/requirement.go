package pos

import (
	"github.com/labstack/echo/v4"
	"github.com/srv-cashpay/pos/dto"
	res "github.com/srv-cashpay/util/s/response"
)

func (h *domainHandler) Requirement(c echo.Context) error {
	var req dto.RequirementRequest
	var resp dto.RequirementResponse

	userid, ok := c.Get("UserId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}

	merchantId, ok := c.Get("MerchantId").(string)
	if !ok {
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, nil).Send(c)
	}
	req.MerchantID = merchantId
	req.UserID = userid

	resp, err := h.servicePos.Requirement(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
