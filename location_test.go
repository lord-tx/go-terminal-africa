package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/lord-tx/go-terminal-africa"
)

func TestLocationService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/countries":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Country]{
				Status: true,
				Data:   []terminal.Country{{Name: "Nigeria", Code: "NG"}},
			})
		case "/states":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.State]{
				Status: true,
				Data:   []terminal.State{{Name: "Lagos", Code: "LA"}},
			})
		case "/cities":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.City]{
				Status: true,
				Data:   []terminal.City{{Name: "Ikeja", Code: "IKJ"}},
			})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("Countries", func(t *testing.T) {
		res, err := client.Locations.Countries(ctx)
		if err != nil || !res.Status || len(res.Data) != 1 || res.Data[0].Code != "NG" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("States", func(t *testing.T) {
		res, err := client.Locations.States(ctx, "NG")
		if err != nil || !res.Status || len(res.Data) != 1 || res.Data[0].Code != "LA" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})

	t.Run("Cities", func(t *testing.T) {
		res, err := client.Locations.Cities(ctx, "NG", "LA")
		if err != nil || !res.Status || len(res.Data) != 1 || res.Data[0].Name != "Ikeja" {
			t.Fatalf("unexpected result: res=%v, err=%v", res, err)
		}
	})
}
