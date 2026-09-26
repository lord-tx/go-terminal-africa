# go-terminal-africa

Official Golang client SDK for the [Terminal Africa API](https://terminal.africa).

Terminal Africa is a shipping and logistics aggregation API supporting carriers such as DHL, FedEx, GIG Logistics, Kwik, Aramex, Dellyman, UPS, and more across Africa.

## Features

- Complete coverage of all 12 Terminal Africa API modules (Addresses, Carriers, Claims, Insurance, Locations, Packagings, Parcels, Rates, Shipments, Transactions, Users, Webhooks).
- Idiomatic Go design with `context.Context` support for all calls.
- Functional client options (`WithBaseURL`, `WithHTTPClient`, `WithTimeout`).
- Package-level helper API for fast, global client operations.

## Installation

```bash
go get github.com/lord-tx/go-terminal-africa
```

## Quick Start

### 1. Initialize the Client

```go
package main

import (
	"context"
	"fmt"
	"log"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func main() {
	// Initialize client using secret key
	client := terminal.NewClient("sk_live_your_secret_key")

	// Or load automatically from environment variable TERMINAL_AFRICA_SECRET_KEY
	// client := terminal.NewClientFromEnv()

	ctx := context.Background()

	// Get Countries
	countriesRes, err := client.Locations.Countries(ctx)
	if err != nil {
		log.Fatalf("Error fetching countries: %v", err)
	}

	fmt.Printf("Status: %t, Total Countries: %d\n", countriesRes.Status, len(countriesRes.Data))
}
```

### 2. Create an Address

```go
addrRes, err := client.Addresses.CreateAddress(ctx, &terminal.CreateAddressParams{
	FirstName:     "Daniel",
	LastName:      "Ozeh",
	Email:         "danielozeh@gmail.com",
	Phone:         "+2348134277988",
	IsResidential: true,
	Line1:         "The Phillipi centre, Oluwalogbon House",
	Line2:         "Obafemi Awolowo way, Alausa",
	City:          "Ikeja",
	State:         "Lagos",
	Country:       "NG",
	Zip:           "100123",
})
if err != nil {
	log.Fatal(err)
}
fmt.Println("Address ID:", addrRes.Data.ID)
```

### 3. Get Shipping Rates & Create Quick Shipment

```go
// Fetch Rates
ratesRes, err := client.Rates.GetShipmentRates(ctx, &terminal.ShipmentRateParams{
	PickupAddress:   "AD-I3ZNBM5Q69P8UQKV",
	DeliveryAddress: "AD-P2TQ9TQEZPETKP7U",
	ParcelID:        "PC-J9FTT8W3FZ7XOSSX",
})
if err != nil {
	log.Fatal(err)
}

for _, rate := range ratesRes.Data {
	fmt.Printf("Carrier: %s | Amount: %s %.2f | Duration: %s\n", rate.CarrierName, rate.Currency, rate.Amount, rate.Duration)
}

// Create Quick Shipment
shipmentRes, err := client.Shipments.CreateQuickShipment(ctx, &terminal.QuickShipmentParams{
	PickupAddress: map[string]any{
		"first_name": "Sender",
		"last_name":  "User",
		"line1":      "123 Main St",
		"city":       "Lagos",
		"state":      "Lagos",
		"country":    "NG",
	},
	DeliveryAddress: map[string]any{
		"first_name": "Recipient",
		"last_name":  "User",
		"line1":      "456 Market St",
		"city":       "Abuja",
		"state":      "FCT",
		"country":    "NG",
	},
	Parcel: map[string]any{
		"description": "Laptop Box",
		"weight_unit": "kg",
		"items": []map[string]any{
			{"name": "MacBook Pro", "weight": 2.5, "quantity": 1, "value": 1200, "currency": "USD"},
		},
	},
})
if err != nil {
	log.Fatal(err)
}
fmt.Println("Created Shipment ID:", shipmentRes.Data.ID)
```

## Available Modules

| Module | Description |
| :--- | :--- |
| `Addresses` | Create, update, fetch, validate addresses & set default sender address |
| `Carriers` | List carriers, enable/disable individual or multiple carriers |
| `Claims` | Fetch and file shipping insurance claims |
| `Insurance` | Query insurance details, calculate premiums, and purchase policies |
| `Locations` | Fetch supported countries, states, and cities |
| `Packagings` | Create and manage custom or terminal default packaging sizes |
| `Parcels` | Create, update, and manage parcels & items |
| `Rates` | Get shipment rates, quick quotes, and multi-piece shipment rates |
| `Shipments` | Create, quick shipment, update, track, cancel, duplicate, and arrange pickup |
| `Transactions` | Retrieve transaction history and wallet ledger entries |
| `Users` | Retrieve user profile, wallet balance, and activated carriers |
| `Webhooks` | Register, update, list, delete, enable, and disable webhooks |

## License

[ISC License](LICENSE)
