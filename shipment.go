package terminal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ShipmentService handles shipment creation, tracking, pickup, and lifecycle operations.
type ShipmentService struct {
	client *Client
}

// Shipment represents a shipment object on Terminal Africa.
type Shipment struct {
	ID              string         `json:"shipment_id,omitempty"`
	AddressFrom     any            `json:"address_from,omitempty"`
	AddressTo       any            `json:"address_to,omitempty"`
	Parcel          any            `json:"parcel,omitempty"`
	Parcels         []any          `json:"parcels,omitempty"`
	ShipmentPurpose string         `json:"shipment_purpose,omitempty"`
	ShipmentType    string         `json:"shipment_type,omitempty"`
	Status          string         `json:"status,omitempty"`
	TrackingNumber  string         `json:"tracking_number,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
	CreatedAt       string         `json:"created_at,omitempty"`
}

// CreateShipmentParams parameters for creating a shipment.
type CreateShipmentParams struct {
	AddressFrom     string         `json:"address_from"`
	AddressTo       string         `json:"address_to"`
	Parcel          string         `json:"parcel,omitempty"`
	Parcels         []string       `json:"parcels,omitempty"`
	ShipmentPurpose string         `json:"shipment_purpose,omitempty"`
	ShipmentType    string         `json:"shipment_type,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
}

// QuickShipmentParams parameters for creating a quick shipment with inline addresses and parcel.
type QuickShipmentParams struct {
	PickupAddress   any            `json:"pickup_address"`
	DeliveryAddress any            `json:"delivery_address"`
	Parcel          any            `json:"parcel"`
	ShipmentPurpose string         `json:"shipment_purpose,omitempty"`
	ShipmentType    string         `json:"shipment_type,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
}

// UpdateShipmentParams parameters for updating a shipment.
type UpdateShipmentParams struct {
	AddressFrom     string         `json:"address_from,omitempty"`
	AddressTo       string         `json:"address_to,omitempty"`
	Parcel          string         `json:"parcel,omitempty"`
	Parcels         []string       `json:"parcels,omitempty"`
	ShipmentPurpose string         `json:"shipment_purpose,omitempty"`
	ShipmentType    string         `json:"shipment_type,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
}

// ShipmentListParams parameters for querying shipments.
type ShipmentListParams struct {
	Page     int    `json:"page,omitempty"`
	PerPage  int    `json:"perPage,omitempty"`
	Populate string `json:"populate,omitempty"`
	Status   string `json:"status,omitempty"`
}

// ArrangePickupParams parameters for arranging shipment pickup.
type ArrangePickupParams struct {
	RateID            string `json:"rate_id"`
	ShipmentID        string `json:"shipment_id,omitempty"`
	PurchaseInsurance bool   `json:"purchase_insurance,omitempty"`
	CashToCollect     int    `json:"cash_to_collect,omitempty"`
	ShipmentPurpose   string `json:"shipment_purpose,omitempty"`
}

// TrackingInfo shipment tracking updates.
type TrackingInfo struct {
	ShipmentID     string `json:"shipment_id,omitempty"`
	Status         string `json:"status,omitempty"`
	TrackingNumber string `json:"tracking_number,omitempty"`
	History        []any  `json:"history,omitempty"`
}

type shipmentIDPayload struct {
	ShipmentID string `json:"shipment_id"`
}

// CreateShipment creates a new shipment.
func (s *ShipmentService) CreateShipment(ctx context.Context, params *CreateShipmentParams) (*APIResponse[Shipment], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/shipments", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Shipment]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateQuickShipment creates a quick shipment with inline payload.
func (s *ShipmentService) CreateQuickShipment(ctx context.Context, params *QuickShipmentParams) (*APIResponse[Shipment], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/shipments/quick", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Shipment]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateShipment updates an existing shipment.
func (s *ShipmentService) UpdateShipment(ctx context.Context, shipmentID string, params *UpdateShipmentParams) (*APIResponse[Shipment], error) {
	path := fmt.Sprintf("/shipments/%s", shipmentID)
	req, err := s.client.newRequest(ctx, http.MethodPut, path, params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Shipment]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetShipments retrieves list of shipments.
func (s *ShipmentService) GetShipments(ctx context.Context, params *ShipmentListParams) (*APIResponse[[]Shipment], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
		if params.Populate != "" {
			queryParams["populate"] = params.Populate
		}
		if params.Status != "" {
			queryParams["status"] = params.Status
		}
	}
	path := buildURLWithQuery("/shipments", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Shipment]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetShipment fetches shipment details by ID.
func (s *ShipmentService) GetShipment(ctx context.Context, shipmentID string) (*APIResponse[Shipment], error) {
	path := fmt.Sprintf("/shipments/%s", shipmentID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Shipment]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// TrackShipment gets real-time tracking info for a shipment.
func (s *ShipmentService) TrackShipment(ctx context.Context, shipmentID string) (*APIResponse[TrackingInfo], error) {
	path := fmt.Sprintf("/shipments/track/%s", shipmentID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[TrackingInfo]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CancelShipment cancels an existing shipment.
func (s *ShipmentService) CancelShipment(ctx context.Context, shipmentID string) (*APIResponse[Shipment], error) {
	payload := shipmentIDPayload{ShipmentID: shipmentID}
	req, err := s.client.newRequest(ctx, http.MethodPost, "/shipments/cancel", payload)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Shipment]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteShipment deletes a shipment.
func (s *ShipmentService) DeleteShipment(ctx context.Context, shipmentID string) (*APIResponse[map[string]any], error) {
	payload := shipmentIDPayload{ShipmentID: shipmentID}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, s.client.baseURL+"/shipments", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if s.client.secretKey != "" {
		req.Header.Set("Authorization", "Bearer "+s.client.secretKey)
	}
	var res APIResponse[map[string]any]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DuplicateShipment duplicates a shipment.
func (s *ShipmentService) DuplicateShipment(ctx context.Context, shipmentID string) (*APIResponse[Shipment], error) {
	payload := shipmentIDPayload{ShipmentID: shipmentID}
	req, err := s.client.newRequest(ctx, http.MethodPost, "/shipments/duplicate", payload)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Shipment]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ArrangePickup schedules pickup and delivery for a shipment.
func (s *ShipmentService) ArrangePickup(ctx context.Context, params *ArrangePickupParams) (*APIResponse[Shipment], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/shipments/pickup", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Shipment]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
