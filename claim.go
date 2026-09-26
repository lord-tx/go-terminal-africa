package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// ClaimService handles communication with insurance claim endpoints.
type ClaimService struct {
	client *Client
}

// Claim represents a claim on Terminal Africa.
type Claim struct {
	ID          string         `json:"claim_id,omitempty"`
	InsuranceID string         `json:"insurance,omitempty"`
	ClaimData   map[string]any `json:"claim,omitempty"`
	Status      string         `json:"status,omitempty"`
	CreatedAt   string         `json:"created_at,omitempty"`
}

// ClaimListParams for listing claims.
type ClaimListParams struct {
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"perPage,omitempty"`
	Status  string `json:"status,omitempty"`
}

// FileClaimParams parameters for filing a new insurance claim.
type FileClaimParams struct {
	InsuranceID string         `json:"insurance"`
	ClaimData   map[string]any `json:"claim"`
}

// GetClaims fetches insurance claims.
func (s *ClaimService) GetClaims(ctx context.Context, params *ClaimListParams) (*APIResponse[[]Claim], error) {
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
	path := buildURLWithQuery("/claims", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Claim]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetClaim retrieves a specific claim by ID.
func (s *ClaimService) GetClaim(ctx context.Context, claimID string) (*APIResponse[Claim], error) {
	path := fmt.Sprintf("/claims/%s", claimID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Claim]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// FileClaim files a new claim for an insured shipment.
func (s *ClaimService) FileClaim(ctx context.Context, params *FileClaimParams) (*APIResponse[Claim], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/claims", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Claim]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
