package terminal

import (
	"context"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	// DefaultBaseURL is the default API URL for Terminal Africa.
	DefaultBaseURL = "https://api.terminal.africa/v1"
)

// Client is the main Terminal Africa API client.
type Client struct {
	secretKey  string
	baseURL    string
	httpClient *http.Client

	// Services
	Addresses    *AddressService
	Carriers     *CarrierService
	Claims       *ClaimService
	Insurance    *InsuranceService
	Locations    *LocationService
	Packagings   *PackagingService
	Parcels      *ParcelService
	Rates        *RateService
	Shipments    *ShipmentService
	Transactions *TransactionService
	Users        *UserService
	Webhooks     *WebhookService
}

var (
	defaultClient *Client
	defaultOnce   sync.Once
)

// NewClient creates a new Terminal Africa API client with the given secret key and options.
func NewClient(secretKey string, opts ...ClientOption) *Client {
	c := &Client{
		secretKey: secretKey,
		baseURL:   DefaultBaseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.baseURL = strings.TrimRight(c.baseURL, "/")

	// Initialize sub-services
	c.Addresses = &AddressService{client: c}
	c.Carriers = &CarrierService{client: c}
	c.Claims = &ClaimService{client: c}
	c.Insurance = &InsuranceService{client: c}
	c.Locations = &LocationService{client: c}
	c.Packagings = &PackagingService{client: c}
	c.Parcels = &ParcelService{client: c}
	c.Rates = &RateService{client: c}
	c.Shipments = &ShipmentService{client: c}
	c.Transactions = &TransactionService{client: c}
	c.Users = &UserService{client: c}
	c.Webhooks = &WebhookService{client: c}

	return c
}

// NewClientFromEnv creates a client using env vars TERMINAL_AFRICA_SECRET_KEY and TERMINAL_AFRICA_URL.
func NewClientFromEnv(opts ...ClientOption) *Client {
	secretKey := os.Getenv("TERMINAL_AFRICA_SECRET_KEY")
	urlEnv := os.Getenv("TERMINAL_AFRICA_URL")

	var envOpts []ClientOption
	if urlEnv != "" {
		if !strings.HasSuffix(urlEnv, "/v1") && !strings.Contains(urlEnv, "/v1") {
			urlEnv = strings.TrimRight(urlEnv, "/") + "/v1"
		}
		envOpts = append(envOpts, WithBaseURL(urlEnv))
	}
	envOpts = append(envOpts, opts...)

	return NewClient(secretKey, envOpts...)
}

// Init sets the default global client instance for package-level function access.
func Init(secretKey string, opts ...ClientOption) *Client {
	defaultClient = NewClient(secretKey, opts...)
	return defaultClient
}

// GetDefaultClient returns the global default client instance or creates one from environment variables.
func GetDefaultClient() *Client {
	defaultOnce.Do(func() {
		if defaultClient == nil {
			defaultClient = NewClientFromEnv()
		}
	})
	return defaultClient
}

// Convenience package-level functions matching JS static TerminalAfrica methods

// CreateAddress creates an address using the default client.
func CreateAddress(ctx context.Context, params *CreateAddressParams) (*APIResponse[Address], error) {
	return GetDefaultClient().Addresses.CreateAddress(ctx, params)
}

// UpdateAddress updates an address using the default client.
func UpdateAddress(ctx context.Context, addressID string, params *UpdateAddressParams) (*APIResponse[Address], error) {
	return GetDefaultClient().Addresses.UpdateAddress(ctx, addressID, params)
}

// GetAddresses fetches addresses using the default client.
func GetAddresses(ctx context.Context, params *ListParams) (*APIResponse[[]Address], error) {
	return GetDefaultClient().Addresses.GetAddresses(ctx, params)
}

// GetAddress fetches a single address using the default client.
func GetAddress(ctx context.Context, addressID string) (*APIResponse[Address], error) {
	return GetDefaultClient().Addresses.GetAddress(ctx, addressID)
}

// ValidateAddress validates an address using the default client.
func ValidateAddress(ctx context.Context, params *ValidateAddressParams) (*APIResponse[AddressValidationResult], error) {
	return GetDefaultClient().Addresses.ValidateAddress(ctx, params)
}

// SetDefaultSenderAddress sets the default sender address.
func SetDefaultSenderAddress(ctx context.Context, addressID string) (*APIResponse[DefaultSenderAddressResponse], error) {
	return GetDefaultClient().Addresses.SetDefaultSenderAddress(ctx, addressID)
}

// GetDefaultSenderAddress gets the default sender address.
func GetDefaultSenderAddress(ctx context.Context) (*APIResponse[Address], error) {
	return GetDefaultClient().Addresses.GetDefaultSenderAddress(ctx)
}

// GetCarriers fetches carriers using default client.
func GetCarriers(ctx context.Context, params *CarrierListParams) (*APIResponse[[]Carrier], error) {
	return GetDefaultClient().Carriers.GetCarriers(ctx, params)
}

// GetCarrier fetches a single carrier.
func GetCarrier(ctx context.Context, carrierID string) (*APIResponse[Carrier], error) {
	return GetDefaultClient().Carriers.GetCarrier(ctx, carrierID)
}

// CreateShipment creates a shipment using default client.
func CreateShipment(ctx context.Context, params *CreateShipmentParams) (*APIResponse[Shipment], error) {
	return GetDefaultClient().Shipments.CreateShipment(ctx, params)
}

// CreateQuickShipment creates a quick shipment using default client.
func CreateQuickShipment(ctx context.Context, params *QuickShipmentParams) (*APIResponse[Shipment], error) {
	return GetDefaultClient().Shipments.CreateQuickShipment(ctx, params)
}

// ListParams specifies general pagination params (page, perPage).
type ListParams struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"perPage,omitempty"`
}
