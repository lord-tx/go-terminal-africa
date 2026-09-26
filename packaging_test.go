package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func TestPackagingService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/packaging":
			if r.Method == http.MethodPost {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Packaging]{
					Status: true,
					Data:   terminal.Packaging{ID: "PKG-1", Name: "Box A"},
				})
			} else {
				_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Packaging]{
					Status: true,
					Data:   []terminal.Packaging{{ID: "PKG-1"}},
				})
			}
		case "/packaging/default/terminal":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Packaging]{
				Status: true,
				Data:   terminal.Packaging{ID: "PKG-DEFAULT"},
			})
		default:
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Packaging]{
				Status: true,
				Data:   terminal.Packaging{ID: "PKG-1"},
			})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("CreatePackaging", func(t *testing.T) {
		res, err := client.Packagings.CreatePackaging(ctx, &terminal.CreatePackagingParams{
			Name: "Box A", Type: "box", Height: 10, Length: 20, Width: 15, Weight: 1, SizeUnit: "cm", WeightUnit: "kg",
		})
		if err != nil || !res.Status || res.Data.ID != "PKG-1" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("TerminalDefaultPackaging", func(t *testing.T) {
		res, err := client.Packagings.TerminalDefaultPackaging(ctx)
		if err != nil || !res.Status || res.Data.ID != "PKG-DEFAULT" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})
}
