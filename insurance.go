package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// InsuranceService handles insurance operations.
type InsuranceService struct {
	client *Client
}

// Insurance represents an insurance policy object on Terminal Africa.
type Insurance struct {
	ID         string  `json:"insurance_id,omitempty"`
	ShipmentID string  `json:"shipment_id,omitempty"`
	Amount     float64 `json:"amount,omitempty"`
	Status     string  `json:"status,omitempty"`
	CreatedAt  string  `json:"created_at,omitempty"`
}

// InsuranceListParams parameters for querying insurance records.
type InsuranceListParams struct {
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"perPage,omitempty"`
	Status  string `json:"status,omitempty"`
}

// PremiumResponse calculated insurance premium response.
type PremiumResponse struct {
	Premium float64 `json:"premium"`
	Value   float64 `json:"value,omitempty"`
}

// PurchaseInsuranceParams parameters for purchasing insurance.
type PurchaseInsuranceParams struct {
	ShipmentID string `json:"shipment"`
}

// GetInsuranceList retrieves a list of insurance policies.
func (s *InsuranceService) GetInsuranceList(ctx context.Context, params *InsuranceListParams) (*APIResponse[[]Insurance], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
		if params.Status != "" {
			queryParams["status"] = params.Status
		}
	}
	path := buildURLWithQuery("/insurance", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Insurance]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetInsurance fetches insurance details by insurance ID.
func (s *InsuranceService) GetInsurance(ctx context.Context, insuranceID string) (*APIResponse[Insurance], error) {
	path := fmt.Sprintf("/insurance/%s", insuranceID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Insurance]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetInsuranceUsingShipment fetches insurance details associated with a shipment.
func (s *InsuranceService) GetInsuranceUsingShipment(ctx context.Context, shipmentID string) (*APIResponse[Insurance], error) {
	path := fmt.Sprintf("/insurance/%s", shipmentID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Insurance]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetInsurancePremium calculates insurance premium for a parcel ID.
func (s *InsuranceService) GetInsurancePremium(ctx context.Context, parcelID string) (*APIResponse[PremiumResponse], error) {
	path := fmt.Sprintf("/insurance/premium?parcel=%s", parcelID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[PremiumResponse]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetInsurancePremiumUsingParcelValue calculates insurance premium using parcel value data.
func (s *InsuranceService) GetInsurancePremiumUsingParcelValue(ctx context.Context, parcelData any) (*APIResponse[PremiumResponse], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/insurance/premium", parcelData)
	if err != nil {
		return nil, err
	}
	var res APIResponse[PremiumResponse]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// PurchaseInsurance purchases insurance for a shipment.
func (s *InsuranceService) PurchaseInsurance(ctx context.Context, shipmentID string) (*APIResponse[Insurance], error) {
	payload := &PurchaseInsuranceParams{ShipmentID: shipmentID}
	req, err := s.client.newRequest(ctx, http.MethodPost, "/insurance/purchase", payload)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Insurance]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
