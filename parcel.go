package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// ParcelService handles parcel management operations.
type ParcelService struct {
	client *Client
}

// ParcelItem represents an item within a parcel.
type ParcelItem struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Value       float64 `json:"value,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
}

// Parcel represents a parcel payload on Terminal Africa.
type Parcel struct {
	ID          string         `json:"parcel_id,omitempty"`
	Description string         `json:"description,omitempty"`
	Packaging   string         `json:"packaging,omitempty"`
	WeightUnit  string         `json:"weight_unit,omitempty"`
	Items       []ParcelItem   `json:"items,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
	CreatedAt   string         `json:"created_at,omitempty"`
}

// CreateParcelParams parameters for creating a new parcel.
type CreateParcelParams struct {
	Description string         `json:"description,omitempty"`
	Packaging   string         `json:"packaging,omitempty"`
	WeightUnit  string         `json:"weight_unit"`
	Items       []ParcelItem   `json:"items"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

// UpdateParcelParams parameters for updating a parcel.
type UpdateParcelParams struct {
	Description string         `json:"description,omitempty"`
	Packaging   string         `json:"packaging,omitempty"`
	WeightUnit  string         `json:"weight_unit,omitempty"`
	Items       []ParcelItem   `json:"items,omitempty"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

// CreateParcel creates a new parcel.
func (s *ParcelService) CreateParcel(ctx context.Context, params *CreateParcelParams) (*APIResponse[Parcel], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/parcels", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Parcel]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateParcel updates an existing parcel.
func (s *ParcelService) UpdateParcel(ctx context.Context, parcelID string, params *UpdateParcelParams) (*APIResponse[Parcel], error) {
	path := fmt.Sprintf("/parcels/%s", parcelID)
	req, err := s.client.newRequest(ctx, http.MethodPut, path, params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Parcel]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetParcels fetches list of parcels.
func (s *ParcelService) GetParcels(ctx context.Context, params *ListParams) (*APIResponse[[]Parcel], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
	}
	path := buildURLWithQuery("/parcels", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Parcel]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetParcel fetches single parcel by ID.
func (s *ParcelService) GetParcel(ctx context.Context, parcelID string) (*APIResponse[Parcel], error) {
	path := fmt.Sprintf("/parcels/%s", parcelID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Parcel]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
