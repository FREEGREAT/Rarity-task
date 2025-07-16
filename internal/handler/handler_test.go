package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"

	"test_task/internal/dto"
	"test_task/internal/entity"
	"test_task/internal/handler"
	"test_task/internal/handler/middleware"
	"test_task/internal/service"
)

type mockLogger struct{}

func (m *mockLogger) Named(name string) middleware.Logger {
	return m
}

func (m *mockLogger) With(args ...interface{}) middleware.Logger {
	return m
}

func (m *mockLogger) WithContext(ctx context.Context) middleware.Logger {
	return m
}

func (m *mockLogger) Debug(msg string, args ...interface{}) {}
func (m *mockLogger) Info(msg string, args ...interface{})  {}
func (m *mockLogger) Warn(msg string, args ...interface{})  {}
func (m *mockLogger) Error(msg string, args ...interface{}) {}
func (m *mockLogger) Fatal(msg string, args ...interface{}) {}

type mockApiClient struct {
	mock.Mock
}

func (m *mockApiClient) GetNftTraitsRarity(ctx context.Context, collectionID string, properties []dto.PropertyKV) (*dto.TraitRarityResponse, error) {
	args := m.Called(ctx, collectionID, properties)
	resp, _ := args.Get(0).(*dto.TraitRarityResponse)
	return resp, args.Error(1)
}

func (m *mockApiClient) GetOwnershipByID(ctx context.Context, id string) (*entity.Ownership, error) {
	args := m.Called(ctx, id)
	resp, _ := args.Get(0).(*entity.Ownership)
	return resp, args.Error(1)
}

func setupRouter(api service.RaribleApiClient) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	h := handler.NewHandler(&mockLogger{}, api)
	return h.InitRoutes(r)
}

func TestGetOwnershipById_Success(t *testing.T) {
	mockAPI := new(mockApiClient)
	router := setupRouter(mockAPI)

	expected := map[string]interface{}{"id": "123"}
	mockAPI.On("GetOwnershipByID", mock.Anything, "123").Return(expected, nil)

	req, _ := http.NewRequest("GET", "/nft/ownerships/123", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockAPI.AssertExpectations(t)
}

func TestGetOwnershipById_MissingID(t *testing.T) {
	router := setupRouter(new(mockApiClient))

	req, _ := http.NewRequest("GET", "/nft/ownerships/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestGetTraitsRarity_Success(t *testing.T) {
	mockAPI := new(mockApiClient)
	router := setupRouter(mockAPI)

	body := dto.TraitRarityRequest{
		CollectionId: "ETHEREUM:0x60e4...",
		Properties: []dto.PropertyKV{
			{Key: "Hat", Value: "Halo"},
			{Key: "Color", Value: "Red"},
		},
	}

	mockAPI.On("GetNftTraitsRarity", mock.Anything, body.CollectionId, body.Properties).
		Return(map[string]interface{}{"rarity": 0.5}, nil)

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/nft/trait-rarities", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockAPI.AssertExpectations(t)
}

func TestGetTraitsRarity_InvalidRequest(t *testing.T) {
	router := setupRouter(new(mockApiClient))

	req, _ := http.NewRequest("POST", "/nft/trait-rarities", bytes.NewBufferString(`{invalid json}`))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
