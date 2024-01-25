package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func response(url string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, url, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	expectedCount := 4
	expectedCode := http.StatusOK

	res := response("/cafe?city=moscow&count=5")

	assert.Equal(t, expectedCode, res.Code)
	require.NotEmpty(t, res.Body)
	body := strings.Split(res.Body.String(), ",")
	require.Equal(t, expectedCount, len(body))
}

func TestMainHandlerWhenStatusOk(t *testing.T) {
	expectedCode := http.StatusOK

	res := response("/cafe?city=moscow&count=1")

	assert.Equal(t, expectedCode, res.Code)
	assert.NotEmpty(t, res.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	expectedCode := http.StatusBadRequest
	expectedBody := "wrong count value"

	res := response("/cafe?city=tver&count=3")

	assert.Equal(t, expectedCode, res.Code)
	require.NotEmpty(t, res.Body)
	assert.Equal(t, expectedBody, res.Body.String())
}
