package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// AddressService handles communication with address related endpoints.
type AddressService struct {
	client *Client
}

// Address represents a Terminal Africa address object.
type Address struct {
	ID            string         `json:"address_id,omitempty"`
	FirstName     string         `json:"first_name,omitempty"`
	LastName      string         `json:"last_name,omitempty"`
	Email         string         `json:"email,omitempty"`
	Phone         string         `json:"phone,omitempty"`
	IsResidential bool           `json:"is_residential,omitempty"`
	Line1         string         `json:"line1,omitempty"`
	Line2         string         `json:"line2,omitempty"`
	City          string         `json:"city,omitempty"`
	State         string         `json:"state,omitempty"`
	Country       string         `json:"country,omitempty"`
	Zip           string         `json:"zip,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
	CreatedAt     string         `json:"created_at,omitempty"`
	UpdatedAt     string         `json:"updated_at,omitempty"`
}

// CreateAddressParams represents parameters for creating a new address.
type CreateAddressParams struct {
	FirstName     string         `json:"first_name,omitempty"`
	LastName      string         `json:"last_name,omitempty"`
	Email         string         `json:"email,omitempty"`
	Phone         string         `json:"phone,omitempty"`
	IsResidential bool           `json:"is_residential,omitempty"`
	Line1         string         `json:"line1"`
	Line2         string         `json:"line2,omitempty"`
	City          string         `json:"city"`
	State         string         `json:"state"`
	Country       string         `json:"country"`
	Zip           string         `json:"zip,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
}

// UpdateAddressParams represents parameters for updating an address.
type UpdateAddressParams struct {
	FirstName     string         `json:"first_name,omitempty"`
	LastName      string         `json:"last_name,omitempty"`
	Email         string         `json:"email,omitempty"`
	Phone         string         `json:"phone,omitempty"`
	IsResidential bool           `json:"is_residential,omitempty"`
	Line1         string         `json:"line1,omitempty"`
	Line2         string         `json:"line2,omitempty"`
	City          string         `json:"city,omitempty"`
	State         string         `json:"state,omitempty"`
	Country       string         `json:"country,omitempty"`
	Zip           string         `json:"zip,omitempty"`
	Metadata      map[string]any `json:"metadata,omitempty"`
}

// ValidateAddressParams represents parameters for validating an address.
type ValidateAddressParams struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
	Zip     string `json:"zip,omitempty"`
}

// AddressValidationResult represents the result of validating an address.
type AddressValidationResult struct {
	IsValid bool   `json:"is_valid"`
	Message string `json:"message,omitempty"`
}

// SetDefaultSenderAddressParams payload.
type SetDefaultSenderAddressParams struct {
	AddressID string `json:"address_id"`
}

// DefaultSenderAddressResponse payload.
type DefaultSenderAddressResponse struct {
	AddressID string `json:"address_id,omitempty"`
	Status    bool   `json:"status,omitempty"`
}

// CreateAddress creates an address on Terminal Africa.
func (s *AddressService) CreateAddress(ctx context.Context, params *CreateAddressParams) (*APIResponse[Address], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/addresses", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Address]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateAddress updates an existing address on Terminal Africa.
func (s *AddressService) UpdateAddress(ctx context.Context, addressID string, params *UpdateAddressParams) (*APIResponse[Address], error) {
	path := fmt.Sprintf("/addresses/%s", addressID)
	req, err := s.client.newRequest(ctx, http.MethodPut, path, params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Address]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAddresses fetches a list of addresses.
func (s *AddressService) GetAddresses(ctx context.Context, params *ListParams) (*APIResponse[[]Address], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
	}
	path := buildURLWithQuery("/addresses", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Address]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAddress fetches details for a specific address.
func (s *AddressService) GetAddress(ctx context.Context, addressID string) (*APIResponse[Address], error) {
	path := fmt.Sprintf("/addresses/%s", addressID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Address]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ValidateAddress validates an address payload.
func (s *AddressService) ValidateAddress(ctx context.Context, params *ValidateAddressParams) (*APIResponse[AddressValidationResult], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/addresses/validate", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[AddressValidationResult]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SetDefaultSenderAddress sets an address as the default sender address.
func (s *AddressService) SetDefaultSenderAddress(ctx context.Context, addressID string) (*APIResponse[DefaultSenderAddressResponse], error) {
	payload := &SetDefaultSenderAddressParams{AddressID: addressID}
	req, err := s.client.newRequest(ctx, http.MethodPost, "/addresses/default/sender", payload)
	if err != nil {
		return nil, err
	}
	var res APIResponse[DefaultSenderAddressResponse]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetDefaultSenderAddress retrieves default sender address.
func (s *AddressService) GetDefaultSenderAddress(ctx context.Context) (*APIResponse[Address], error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/addresses/default/sender", nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Address]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
