package pos

import (
	dto "github.com/srv-cashpay/pos/dto"
)

func (u *posService) Requirement(req dto.RequirementRequest) (dto.RequirementResponse, error) {
	// Validasi refresh token dan dapatkan user ID

	requirement, err := u.Repo.Requirement(req)
	if err != nil {
		return dto.RequirementResponse{}, err
	}

	return requirement, nil
}
