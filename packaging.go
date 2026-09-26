package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// PackagingService handles packaging endpoints.
type PackagingService struct {
	client *Client
}

// Packaging represents packaging metadata on Terminal Africa.
type Packaging struct {
	ID         string  `json:"packaging_id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Type       string  `json:"type,omitempty"`
	Height     float64 `json:"height,omitempty"`
	Length     float64 `json:"length,omitempty"`
	Width      float64 `json:"width,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
	SizeUnit   string  `json:"size_unit,omitempty"`
	WeightUnit string  `json:"weight_unit,omitempty"`
	CreatedAt  string  `json:"created_at,omitempty"`
}

// CreatePackagingParams parameters for creating packaging.
type CreatePackagingParams struct {
	Name       string  `json:"name"`
	Type       string  `json:"type"`
	Height     float64 `json:"height"`
	Length     float64 `json:"length"`
	Width      float64 `json:"width"`
	Weight     float64 `json:"weight"`
	SizeUnit   string  `json:"size_unit"`
	WeightUnit string  `json:"weight_unit"`
}

// UpdatePackagingParams parameters for updating packaging.
type UpdatePackagingParams struct {
	Name       string  `json:"name,omitempty"`
	Type       string  `json:"type,omitempty"`
	Height     float64 `json:"height,omitempty"`
	Length     float64 `json:"length,omitempty"`
	Width      float64 `json:"width,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
	SizeUnit   string  `json:"size_unit,omitempty"`
	WeightUnit string  `json:"weight_unit,omitempty"`
}

// PackagingListParams parameters for querying packaging items.
type PackagingListParams struct {
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"perPage,omitempty"`
	Type    string `json:"type,omitempty"`
}

// CreatePackaging creates a packaging entry.
func (s *PackagingService) CreatePackaging(ctx context.Context, params *CreatePackagingParams) (*APIResponse[Packaging], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/packaging", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Packaging]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdatePackaging updates existing packaging details.
func (s *PackagingService) UpdatePackaging(ctx context.Context, packagingID string, params *UpdatePackagingParams) (*APIResponse[Packaging], error) {
	path := fmt.Sprintf("/packaging/%s", packagingID)
	req, err := s.client.newRequest(ctx, http.MethodPut, path, params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Packaging]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetPackagings fetches packagings.
func (s *PackagingService) GetPackagings(ctx context.Context, params *PackagingListParams) (*APIResponse[[]Packaging], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
		if params.Type != "" {
			queryParams["type"] = params.Type
		}
	}
	path := buildURLWithQuery("/packaging", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Packaging]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetPackaging fetches single packaging details.
func (s *PackagingService) GetPackaging(ctx context.Context, packagingID string) (*APIResponse[Packaging], error) {
	path := fmt.Sprintf("/packaging/%s", packagingID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Packaging]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// TerminalDefaultPackaging gets default terminal packaging configuration.
func (s *PackagingService) TerminalDefaultPackaging(ctx context.Context) (*APIResponse[Packaging], error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/packaging/default/terminal", nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Packaging]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
