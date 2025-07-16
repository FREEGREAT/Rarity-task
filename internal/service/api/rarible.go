package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"test_task/internal/dto"
	"test_task/internal/entity"
	"test_task/internal/service"
)

type raribleApi struct {
	baseUrl    string
	apiKey     string
	httpClient *http.Client
}

func NewRaribleApiClient(baseUrl, apiKey string) service.RaribleApiClient {
	return &raribleApi{
		baseUrl:    baseUrl,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c *raribleApi) GetOwnershipByID(ctx context.Context, id string) (*entity.Ownership, error) {
	url := fmt.Sprintf("%s/ownerships/%s", c.baseUrl, id)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Set("X-API-KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rarible api error: %s", resp.Status)
	}

	var ownership entity.Ownership
	if err := json.NewDecoder(resp.Body).Decode(&ownership); err != nil {
		return nil, err
	}

	return &ownership, nil
}

// GetNftTraitsRarity implements service.RaribleApiClient.
func (c *raribleApi) GetNftTraitsRarity(ctx context.Context, collectionId string, properties []dto.PropertyKV) (*dto.TraitRarityResponse, error) {
	url := fmt.Sprintf("%s/items/traits/rarity", c.baseUrl)

	body := dto.TraitRarityRequest{
		CollectionId: collectionId,
		Properties:   properties,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-API-KEY", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("Rarible API status: %s, body: %s", resp.Status, string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rarible api error: %s", resp.Status)
	}

	var traitsResp dto.TraitRarityResponse
	if err := json.Unmarshal(bodyBytes, &traitsResp); err != nil {
		return nil, err
	}

	return &traitsResp, nil
}
