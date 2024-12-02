package dto

import (
	"github.com/srv-cashpay/pos/entity"
)

type PaginationRequest struct {
	UserID       string       `json:"-"`
	Limit        int          `json:"limit"`
	Page         int          `json:"page"`
	Sort         string       `json:"sort"`
	TotalRows    int          `json:"total_rows"`
	TotalPages   int          `json:"total_page"`
	FirstPage    string       `json:"first_page"`
	PreviousPage string       `json:"previous_page"`
	NextPage     string       `json:"next_page"`
	LastPage     string       `json:"last_page"`
	FromRow      int          `json:"from_row"`
	ToRow        int          `json:"to_row"`
	Rows         []entity.Pos `json:"rows"`
	Searchs      []Search     `json:"searchs"`
}

type PaginationResponse struct {
	Limit        int           `json:"limit"`
	Page         int           `json:"page"`
	Sort         string        `json:"sort"`
	TotalRows    int           `json:"total_rows"`
	TotalPages   int           `json:"total_page"`
	FirstPage    string        `json:"first_page"`
	PreviousPage string        `json:"previous_page"`
	NextPage     string        `json:"next_page"`
	LastPage     string        `json:"last_page"`
	FromRow      int           `json:"from_row"`
	ToRow        int           `json:"to_row"`
	Data         []PosResponse `json:"data"`
	Searchs      []Search      `json:"searchs"`
}

type GetAllRequest struct {
	Page     int      `query:"page"`
	Limit    int      `query:"limit"`
	Searchs  []Search `query:"searchs" json:"searchs"`
	SortBy   string   `query:"sort_by"`
	SortDesc bool     `query:"sort_desc"`
}

// type PaginationResponse struct {
// 	TotalData    int           `json:"total_data"`
// 	TotalRows    int           `json:"total_rows"`
// 	Limit        int           `json:"limit"`
// 	PreviousPage int           `json:"previous_page"`
// 	NextPage     int           `json:"next_page"`
// 	NextPageData int           `json:"next_page_data"`
// 	Data         []entity.Blog `json:"data"`
// 	Searchs      []Search      `json:"searchs"`
// }
