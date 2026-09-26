package terminal_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	terminal "github.com/danielozeh/go-terminal-africa"
)

func TestTransactionService(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/transactions":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[[]terminal.Transaction]{
				Status: true,
				Data:   []terminal.Transaction{{ID: "TX-1", Amount: 5000}},
			})
		case "/transactions/TX-1":
			_ = json.NewEncoder(w).Encode(terminal.APIResponse[terminal.Transaction]{
				Status: true,
				Data:   terminal.Transaction{ID: "TX-1", Amount: 5000},
			})
		}
	}))
	defer server.Close()

	client := terminal.NewClient("sk_test", terminal.WithBaseURL(server.URL))
	ctx := context.Background()

	t.Run("GetTransactions", func(t *testing.T) {
		res, err := client.Transactions.GetTransactions(ctx, &terminal.TransactionListParams{WalletID: "W-1"})
		if err != nil || !res.Status || len(res.Data) != 1 {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})

	t.Run("GetTransaction", func(t *testing.T) {
		res, err := client.Transactions.GetTransaction(ctx, "TX-1")
		if err != nil || !res.Status || res.Data.ID != "TX-1" {
			t.Fatalf("unexpected res: %v, err: %v", res, err)
		}
	})
}
