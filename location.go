package terminal

import (
	"context"
	"net/http"
)

// LocationService handles location endpoints (countries, states, cities).
type LocationService struct {
	client *Client
}

// Country represents a country model.
type Country struct {
	Name string `json:"name,omitempty"`
	ISO2 string `json:"iso2,omitempty"`
	ISO3 string `json:"iso3,omitempty"`
	Code string `json:"code,omitempty"`
}

// State represents a state/region model.
type State struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

// City represents a city model.
type City struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

// Countries fetches supported countries.
func (s *LocationService) Countries(ctx context.Context) (*APIResponse[[]Country], error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/countries", nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Country]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// States fetches states for a country.
func (s *LocationService) States(ctx context.Context, countryCode string) (*APIResponse[[]State], error) {
	queryParams := map[string]string{"country_code": countryCode}
	path := buildURLWithQuery("/states", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]State]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Cities fetches cities for a country and optional state.
func (s *LocationService) Cities(ctx context.Context, countryCode, stateCode string) (*APIResponse[[]City], error) {
	queryParams := map[string]string{"country_code": countryCode}
	if stateCode != "" {
		queryParams["state_code"] = stateCode
	}
	path := buildURLWithQuery("/cities", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]City]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
