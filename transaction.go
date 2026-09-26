package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// TransactionService handles financial & transaction history endpoints.
type TransactionService struct {
	client *Client
}

// Transaction represents a financial transaction record on Terminal Africa.
type Transaction struct {
	ID        string  `json:"transaction_id,omitempty"`
	WalletID  string  `json:"wallet_id,omitempty"`
	Amount    float64 `json:"amount,omitempty"`
	Currency  string  `json:"currency,omitempty"`
	Type      string  `json:"type,omitempty"`
	Status    string  `json:"status,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
}

// TransactionListParams parameters for querying transactions.
type TransactionListParams struct {
	Page     int    `json:"page,omitempty"`
	PerPage  int    `json:"perPage,omitempty"`
	WalletID string `json:"wallet,omitempty"`
}

// GetTransactions fetches list of transactions.
func (s *TransactionService) GetTransactions(ctx context.Context, params *TransactionListParams) (*APIResponse[[]Transaction], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
		if params.WalletID != "" {
			queryParams["wallet"] = params.WalletID
		}
	}
	path := buildURLWithQuery("/transactions", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Transaction]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetTransaction retrieves transaction details by ID.
func (s *TransactionService) GetTransaction(ctx context.Context, transactionID string) (*APIResponse[Transaction], error) {
	path := fmt.Sprintf("/transactions/%s", transactionID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Transaction]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
