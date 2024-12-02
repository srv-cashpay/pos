package history

import (
	"fmt"

	"github.com/labstack/echo/v4"
	dto "github.com/srv-cashpay/pos/dto"
)

func (s *historyService) Get(context echo.Context, req dto.PaginationRequest) dto.PaginationResponse {
	if req.Page < 1 {
		req.Page = 1
	}

	operationResult, totalPages := s.Repo.Get(req)

	// Set page 1-based untuk pagination link
	urlPath := context.Request().URL.Path
	searchQueryParams := ""

	for _, search := range req.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	// URL dengan base 1 untuk first page, previous page, next page, dan last page
	// data := operationResult.Result.(dto.PaginationRequest)
	data := operationResult
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=1&sort=%s", urlPath, req.Limit, req.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, req.Limit, totalPages, req.Sort) + searchQueryParams

	if req.Page > 1 {
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, req.Limit, req.Page-1, req.Sort) + searchQueryParams
	}
	if req.Page < totalPages {
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, req.Limit, req.Page+1, req.Sort) + searchQueryParams
	}

	return dto.PaginationResponse{
		Limit:        data.Limit,
		Page:         data.Page,
		Sort:         data.Sort,
		TotalRows:    data.TotalRows,
		TotalPages:   data.TotalPages,
		FirstPage:    data.FirstPage,
		PreviousPage: data.PreviousPage,
		NextPage:     data.NextPage,
		LastPage:     data.LastPage,
		FromRow:      data.FromRow,
		ToRow:        data.ToRow,
		Data:         data.Data,
		Searchs:      data.Searchs,
	}
}
