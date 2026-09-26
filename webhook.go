package terminal

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// WebhookService handles webhooks management operations.
type WebhookService struct {
	client *Client
}

// Webhook represents a webhook subscription object.
type Webhook struct {
	ID        string   `json:"webhook_id,omitempty"`
	Name      string   `json:"name,omitempty"`
	URL       string   `json:"url,omitempty"`
	Active    bool     `json:"active,omitempty"`
	Events    []string `json:"events,omitempty"`
	Live      bool     `json:"live,omitempty"`
	CreatedAt string   `json:"created_at,omitempty"`
}

// CreateWebhookParams parameters for registering a webhook.
type CreateWebhookParams struct {
	Name   string   `json:"name"`
	URL    string   `json:"url"`
	Active bool     `json:"active"`
	Events []string `json:"events"`
	Live   bool     `json:"live"`
}

// UpdateWebhookParams parameters for updating a webhook subscription.
type UpdateWebhookParams struct {
	Name   string   `json:"name,omitempty"`
	URL    string   `json:"url,omitempty"`
	Active *bool    `json:"active,omitempty"`
	Events []string `json:"events,omitempty"`
	Live   *bool    `json:"live,omitempty"`
}

// CreateWebhook creates a new webhook.
func (s *WebhookService) CreateWebhook(ctx context.Context, params *CreateWebhookParams) (*APIResponse[Webhook], error) {
	req, err := s.client.newRequest(ctx, http.MethodPost, "/webhooks", params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Webhook]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateWebhook updates an existing webhook.
func (s *WebhookService) UpdateWebhook(ctx context.Context, webhookID string, params *UpdateWebhookParams) (*APIResponse[Webhook], error) {
	path := fmt.Sprintf("/webhooks/%s", webhookID)
	req, err := s.client.newRequest(ctx, http.MethodPut, path, params)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Webhook]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetWebhooks lists webhooks.
func (s *WebhookService) GetWebhooks(ctx context.Context, params *ListParams) (*APIResponse[[]Webhook], error) {
	queryParams := make(map[string]string)
	if params != nil {
		if params.Page > 0 {
			queryParams["page"] = strconv.Itoa(params.Page)
		}
		if params.PerPage > 0 {
			queryParams["perPage"] = strconv.Itoa(params.PerPage)
		}
	}
	path := buildURLWithQuery("/webhooks", queryParams)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[[]Webhook]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetWebhook retrieves single webhook by ID.
func (s *WebhookService) GetWebhook(ctx context.Context, webhookID string) (*APIResponse[Webhook], error) {
	path := fmt.Sprintf("/webhooks/%s", webhookID)
	req, err := s.client.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Webhook]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteWebhook deletes a webhook by ID.
func (s *WebhookService) DeleteWebhook(ctx context.Context, webhookID string) (*APIResponse[map[string]any], error) {
	path := fmt.Sprintf("/webhooks/%s", webhookID)
	req, err := s.client.newRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[map[string]any]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// EnableWebhook enables a webhook by ID.
func (s *WebhookService) EnableWebhook(ctx context.Context, webhookID string) (*APIResponse[Webhook], error) {
	path := fmt.Sprintf("/webhooks/enable/%s", webhookID)
	req, err := s.client.newRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Webhook]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DisableWebhook disables a webhook by ID.
func (s *WebhookService) DisableWebhook(ctx context.Context, webhookID string) (*APIResponse[Webhook], error) {
	path := fmt.Sprintf("/webhooks/disable/%s", webhookID)
	req, err := s.client.newRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}
	var res APIResponse[Webhook]
	if err := s.client.do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
