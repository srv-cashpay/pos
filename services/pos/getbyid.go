package pos

import (
	dto "github.com/srv-cashpay/pos/dto"
)

func (b *posService) GetById(req dto.GetByIdRequest) (*dto.PosUpdateResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
