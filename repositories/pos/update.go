package pos

import (
	dto "github.com/srv-cashpay/pos/dto"
	"github.com/srv-cashpay/pos/entity"
)

func (b *posRepository) Update(req dto.PosUpdateRequest) (dto.PosUpdateResponse, error) {
	tr := dto.GetByIdRequest{
		ID: req.ID,
	}

	request := entity.Pos{
		StatusPayment: req.StatusPayment,
		UpdatedBy:     req.UpdatedBy,
	}

	product, err := b.GetById(tr)
	if err != nil {
		return dto.PosUpdateResponse{}, err
	}

	err = b.DB.Where("ID = ?", req.ID).Updates(entity.Pos{
		StatusPayment: request.StatusPayment,
		UpdatedBy:     request.UpdatedBy,
	}).Error
	if err != nil {
		return dto.PosUpdateResponse{}, err
	}

	response := dto.PosUpdateResponse{
		StatusPayment: request.StatusPayment,
		UpdatedBy:     request.UpdatedBy,
		ID:            product.ID,
	}

	return response, nil
}
