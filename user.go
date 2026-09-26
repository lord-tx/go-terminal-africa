package terminal

import (
	"context"
	"fmt"
	"net/http"
)

// UserService handles user profile and wallet operations.
type UserService struct {
	client *Client
}

// User represents user account details.
type User struct {
	ID        string `json:"user_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

// WalletInfo represents detailed wallet info.
type WalletInfo struct {
	WalletID  string  `json:"wallet_id,omitempty"`
	UserID    string  `json:"user_id,omitempty"`
	Balance   float64 `json:"balance,omitempty"`
	Currency  string  `json:"currency,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
}

// WalletBalance represents current wallet balance.
type WalletBalance struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency,omitempty"`
}

// GetUser fetches details for a user.
func (s *UserService) GetUser(ctx context.Context, userID string) (*APIResponse[User], error) {
	path := fmt.Sprintf("/users/%s", userID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[User]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// WalletInfo fetches wallet information for a given user ID.
func (s *UserService) WalletInfo(ctx context.Context, userID string) (*APIResponse[WalletInfo], error) {
	path := fmt.Sprintf("/users/wallet?user_id=%s", userID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[WalletInfo]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// WalletBalance fetches wallet balance for a user.
func (s *UserService) WalletBalance(ctx context.Context, userID string) (*APIResponse[WalletBalance], error) {
	path := fmt.Sprintf("/users/wallet-balance?user_id=%s", userID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[WalletBalance]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Carriers fetches activated carriers for the user account.
func (s *UserService) Carriers(ctx context.Context) (*APIResponse[[]Carrier], error) {
	req, err := s.client.newRequest(ctx, http.MethodGet, "/users/carriers", nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Carrier]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
