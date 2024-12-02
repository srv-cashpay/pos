package pos

import "github.com/srv-cashpay/pos/dto"

func (b *posService) Update(req dto.PosUpdateRequest) (dto.PosUpdateResponse, error) {
	request := dto.PosUpdateRequest{
		StatusPayment: req.StatusPayment,
		UpdatedBy:     req.UpdatedBy,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.PosUpdateResponse{
		StatusPayment: request.StatusPayment,
		UpdatedBy:     request.UpdatedBy,
	}

	return response, nil
}
