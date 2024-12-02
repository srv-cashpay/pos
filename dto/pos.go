package dto

type GetByIdRequest struct {
	ID string `param:"id" validate:"required"`
}

type PosRequest struct {
	ID         string           `json:"id"`
	UserID     string           `json:"user_id"`
	MerchantID string           `json:"merchant_id"`
	Quantity   int              `json:"quantity"`
	CreatedBy  string           `json:"created_by"`
	Product    []ProductRequest `json:"product"`
}

type ProductRequest struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type PosResponse struct {
	ID         string            `json:"id"`
	UserID     string            `json:"user_id"`
	MerchantID string            `json:"merchant_id"`
	CreatedBy  string            `json:"created_by"`
	Quantity   int               `json:"-"`
	Product    []ProductResponse `json:"product"`
	TotalPrice int               `json:"total_price"`
}

type ProductResponse struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
}

type PosUpdateRequest struct {
	ID            string `json:"id"`
	StatusPayment string `json:"status_payment"`
	UpdatedBy     string `json:"updated_by"`
}

type PosUpdateResponse struct {
	ID            string `json:"id"`
	StatusPayment string `json:"status_payment"`
	UpdatedBy     string `json:"updated_by"`
}
