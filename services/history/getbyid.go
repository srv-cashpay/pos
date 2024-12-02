package history

import (
	dto "github.com/srv-cashpay/pos/dto"
)

func (b *historyService) GetById(req dto.GetByIdRequest) (*dto.PosResponse, error) {
	transaction, err := b.Repo.GetById(req)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
