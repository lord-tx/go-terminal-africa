package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func TestClaimService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/claims":
			if r.Method == http.MethodPost {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Claim]{
					Status: true,
					Data:   terminal.Claim{ID: "CLM-1", InsuranceID: "INS-1"},
				})
			} else {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Claim]{
					Status: true,
					Data:   []terminal.Claim{{ID: "CLM-1"}},
				})
			}
		case "/claims/CLM-1":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Claim]{
				Status: true,
				Data:   terminal.Claim{ID: "CLM-1"},
			})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("FileClaim", func(t *testing.T) {
		res, err := client.Claims.FileClaim(ctx, &terminal.FileClaimParams{
			InsuranceID: "INS-1", ClaimData: map[string]any{"reason": "lost"},
		})
		if err != nil || !res.Status || res.Data.ID != "CLM-1" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("GetClaims", func(t *testing.T) {
		res, err := client.Claims.GetClaims(ctx, &terminal.ClaimListParams{Page: 1})
		if err != nil || !res.Status || len(res.Data) != 1 {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})
}
