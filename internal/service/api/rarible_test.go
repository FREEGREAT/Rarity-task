package api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"test_task/internal/dto"
	"test_task/internal/entity"
	"test_task/internal/service/api"
)

func TestRaribleApi_GetOwnershipByID_Success(t *testing.T) {
	expected := &entity.Ownership{
		ID:            "ownership-id",
		Blockchain:    "ETHEREUM",
		ItemID:        "item-id",
		Contract:      "0xabc123",
		Collection:    "collection-id",
		TokenID:       "token-id",
		Owner:         "0xowner",
		Value:         "1",
		CreatedAt:     "2024-01-01T00:00:00Z",
		LastUpdatedAt: "2024-01-02T00:00:00Z",
		Creators: []entity.Creator{
			{Account: "0xcreator", Value: 10000},
		},
		LazyValue:    "0",
		Pending:      []string{"tx1", "tx2"},
		OriginOrders: []string{"order1", "order2"},
		Version:      1,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/ownerships/") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := api.NewRaribleApiClient(server.URL, "dummy-key")
	resp, err := client.GetOwnershipByID(context.Background(), expected.ID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if fmt.Sprintf("%+v", resp) != fmt.Sprintf("%+v", expected) {
		t.Errorf("expected %+v, got %+v", expected, resp)
	}
}
func TestRaribleApi_GetOwnershipByID_ErrorStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	}))
	defer server.Close()

	client := api.NewRaribleApiClient(server.URL, "dummy-key")
	_, err := client.GetOwnershipByID(context.Background(), "nonexistent-id")
	if err == nil || !strings.Contains(err.Error(), "404") {
		t.Fatalf("expected 404 error, got: %v", err)
	}
}

func TestRaribleApi_GetNftTraitsRarity_Success(t *testing.T) {
	expected := &dto.TraitRarityResponse{
		Traits: []dto.ExtendedTraitProperty{
			{Key: "Hat", Value: "Halo", Rarity: "rare"},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.URL.Path != "/items/traits/rarity" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}

		var req dto.TraitRarityRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("failed to decode body: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	client := api.NewRaribleApiClient(server.URL, "dummy-key")
	req := dto.TraitRarityRequest{
		CollectionId: "col1",
		Properties: []dto.PropertyKV{
			{Key: "Hat", Value: "Halo"},
		},
	}
	resp, err := client.GetNftTraitsRarity(context.Background(), req.CollectionId, req.Properties)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(resp.Traits) != 1 || resp.Traits[0].Key != "Hat" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestRaribleApi_GetNftTraitsRarity_BadStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))
	defer server.Close()

	client := api.NewRaribleApiClient(server.URL, "dummy-key")
	_, err := client.GetNftTraitsRarity(context.Background(), "col", []dto.PropertyKV{
		{Key: "Hat", Value: "Halo"},
	})
	if err == nil || !strings.Contains(err.Error(), "500") {
		t.Fatalf("expected 500 error, got: %v", err)
	}
}
