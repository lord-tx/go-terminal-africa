package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// RateService handles shipment rate & quote calculations.
type RateService struct {
	client *Client
}

// Rate represents a shipping rate object.
type Rate struct {
	ID                 string  `json:"rate_id,omitempty"`
	CarrierID          string  `json:"carrier_id,omitempty"`
	CarrierName        string  `json:"carrier_name,omitempty"`
	CarrierLogo        string  `json:"carrier_logo,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	Currency           string  `json:"currency,omitempty"`
	DeliveryDate       string  `json:"delivery_date,omitempty"`
	DeliveryTime       string  `json:"delivery_time,omitempty"`
	Distance           float64 `json:"distance,omitempty"`
	ChargeableWeight   float64 `json:"chargeable_weight,omitempty"`
	Duration           string  `json:"duration,omitempty"`
	CODSupported       bool    `json:"cash_on_delivery,omitempty"`
	InsuranceSupported bool    `json:"insurance_supported,omitempty"`
}

// ShipmentRateParams query options for fetching shipment rates.
type ShipmentRateParams struct {
	Currency        string `json:"currency,omitempty"`
	DeliveryAddress string `json:"delivery_address,omitempty"`
	PickupAddress   string `json:"pickup_address,omitempty"`
	ParcelID        string `json:"parcel_id,omitempty"`
	ShipmentID      string `json:"shipment_id,omitempty"`
	CashOnDelivery  *bool  `json:"cash_on_delivery,omitempty"`
	CarrierID       string `json:"carrier_id,omitempty"`
}

// QuotesParams body options for calculating quick quotes.
type QuotesParams struct {
	PickupAddress   any    `json:"pickup_address"`
	DeliveryAddress any    `json:"delivery_address"`
	Parcel          any    `json:"parcel"`
	CarrierID       string `json:"carrier_id,omitempty"`
	Currency        string `json:"currency,omitempty"`
	CashOnDelivery  any    `json:"cash_on_delivery,omitempty"`
}

// MultiPieceRateParams body options for multi-piece rates.
type MultiPieceRateParams struct {
	Currency        string   `json:"currency,omitempty"`
	DeliveryAddress string   `json:"delivery_address,omitempty"`
	PickupAddress   string   `json:"pickup_address,omitempty"`
	Parcels         []string `json:"parcels"`
	ShipmentID      string   `json:"shipment_id,omitempty"`
}

// GetShipmentRates fetches available shipping rates for given parameters.
func (s *RateService) GetShipmentRates(ctx context.Context, params *ShipmentRateParams) (*APIResponse[[]Rate], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Currency != "" {
			queryParams["currency"] = params.Currency
		}
		if params.DeliveryAddress != "" {
			queryParams["delivery_address"] = params.DeliveryAddress
		}
		if params.PickupAddress != "" {
			queryParams["pickup_address"] = params.PickupAddress
		}
		if params.ParcelID != "" {
			queryParams["parcel_id"] = params.ParcelID
		}
		if params.ShipmentID != "" {
			queryParams["shipment_id"] = params.ShipmentID
		}
		if params.CashOnDelivery != nil {
			queryParams["cash_on_delivery"] = strconv.FormatBool(*params.CashOnDelivery)
		}
		if params.CarrierID != "" {
			queryParams["carrier_id"] = params.CarrierID
		}
	}
	path := buildURLWithQuery("/rates/shipment", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Rate]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetQuotesForShipment fetches quick shipping quotes.
func (s *RateService) GetQuotesForShipment(ctx context.Context, params *QuotesParams) (*APIResponse[[]Rate], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/rates/shipment/quotes", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Rate]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetMultiPieceShipmentRates calculates rates for multi-piece shipments.
func (s *RateService) GetMultiPieceShipmentRates(ctx context.Context, params *MultiPieceRateParams) (*APIResponse[[]Rate], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/rates/multi/shipment", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Rate]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRates fetches generated rates.
func (s *RateService) GetRates(ctx context.Context, params *ListParams) (*APIResponse[[]Rate], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
	}
	path := buildURLWithQuery("/rates", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Rate]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRate fetches single rate by ID.
func (s *RateService) GetRate(ctx context.Context, rateID string) (*APIResponse[Rate], error) {
	path := fmt.Sprintf("/rates/%s", rateID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Rate]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
