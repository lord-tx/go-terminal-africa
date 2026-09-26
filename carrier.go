package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// CarrierService handles communication with carrier endpoints.
type CarrierService struct {
	client *Client
}

// Carrier represents a logistics carrier object on Terminal Africa.
type Carrier struct {
	ID            string `json:"carrier_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Slug          string `json:"slug,omitempty"`
	Logo          string `json:"logo,omitempty"`
	Domestic      bool   `json:"domestic,omitempty"`
	Regional      bool   `json:"regional,omitempty"`
	International bool   `json:"international,omitempty"`
	Active        bool   `json:"active,omitempty"`
}

// CarrierListParams parameters for querying carriers.
type CarrierListParams struct {
	Page    int   `json:"page,omitempty"`
	PerPage int   `json:"perPage,omitempty"`
	Active  *bool `json:"active,omitempty"`
}

// CarrierToggleParams options for enabling or disabling domestic/regional/international service.
type CarrierToggleParams struct {
	Domestic      bool `json:"domestic"`
	Regional      bool `json:"regional"`
	International bool `json:"international"`
}

// CarrierItemToggle parameters for multiple carrier operations.
type CarrierItemToggle struct {
	CarrierID     string `json:"carrier_id"`
	Domestic      bool   `json:"domestic"`
	Regional      bool   `json:"regional"`
	International bool   `json:"international"`
}

// EnableMultipleCarriersParams payload for bulk carrier enable/disable.
type EnableMultipleCarriersParams struct {
	Carriers []CarrierItemToggle `json:"carriers"`
}

// GetCarriers fetches a list of carriers.
func (s *CarrierService) GetCarriers(ctx context.Context, params *CarrierListParams) (*APIResponse[[]Carrier], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
		if params.Active != nil {
			queryParams["active"] = strconv.FormatBool(*params.Active)
		}
	}
	path := buildURLWithQuery("/carriers", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Carrier]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCarrier fetches details of a single carrier by ID.
func (s *CarrierService) GetCarrier(ctx context.Context, carrierID string) (*APIResponse[Carrier], error) {
	path := fmt.Sprintf("/carriers/%s", carrierID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Carrier]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// EnableCarrier enables a specific carrier.
func (s *CarrierService) EnableCarrier(ctx context.Context, carrierID string, params *CarrierToggleParams) (*APIResponse[Carrier], error) {
	queryParams := make(map[string]string)
	if params != nil {
		queryParams["domestic"] = strconv.FormatBool(params.Domestic)
		queryParams["regional"] = strconv.FormatBool(params.Regional)
		queryParams["international"] = strconv.FormatBool(params.International)
	}
	path := buildURLWithQuery(fmt.Sprintf("/carriers/enable/%s", carrierID), queryParams)
	req, err := s.client.newRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Carrier]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DisableCarrier disables a specific carrier.
func (s *CarrierService) DisableCarrier(ctx context.Context, carrierID string, params *CarrierToggleParams) (*APIResponse[Carrier], error) {
	queryParams := make(map[string]string)
	if params != nil {
		queryParams["domestic"] = strconv.FormatBool(params.Domestic)
		queryParams["regional"] = strconv.FormatBool(params.Regional)
		queryParams["international"] = strconv.FormatBool(params.International)
	}
	path := buildURLWithQuery(fmt.Sprintf("/carriers/disable/%s", carrierID), queryParams)
	req, err := s.client.newRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Carrier]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// EnableMultipleCarriers enables multiple carriers at once.
func (s *CarrierService) EnableMultipleCarriers(ctx context.Context, payload *EnableMultipleCarriersParams) (*APIResponse[[]Carrier], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/carriers/multiple/enable", payload)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Carrier]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DisableMultipleCarriers disables multiple carriers at once.
func (s *CarrierService) DisableMultipleCarriers(ctx context.Context, payload *EnableMultipleCarriersParams) (*APIResponse[[]Carrier], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/carriers/multiple/disable", payload)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Carrier]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
