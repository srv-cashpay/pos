package pos

import (
	"github.com/srv-cashpay/merchant/entity"
	dto "github.com/srv-cashpay/pos/dto"
)

func (u *posRepository) Requirement(req dto.RequirementRequest) (dto.RequirementResponse, error) {
	var d entity.Discount
	var t entity.Tax

	// Ambil discount
	if err := u.DB.
		Where("merchant_id = ? AND status = ?", req.MerchantID, 1). // misal hanya yg aktif
		First(&d).Error; err != nil {
		return dto.RequirementResponse{}, err
	}

	// Ambil tax
	if err := u.DB.
		Where("merchant_id = ? AND status = ?", req.MerchantID, 1).
		First(&t).Error; err != nil {
		return dto.RequirementResponse{}, err
	}

	resp := dto.RequirementResponse{
		Tax: []dto.TaxResponse{{
			Tax:           t.Tax,
			TaxPercentage: t.TaxPercentage,
		},
		},
		Discount: []dto.DiscountResponse{{
			DiscountName:       d.DiscountName,
			DiscountPercentage: d.DiscountPercentage,
		},
		},
	}

	return resp, nil
}
