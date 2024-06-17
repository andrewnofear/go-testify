package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenNotHaveCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=unkonwn", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	val, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	require.Equal(t, string(val), "wrong city value")
}

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	val, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	require.NotEqual(t, val, nil)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	val, err := io.ReadAll(responseRecorder.Body)
	require.NoError(t, err)
	require.NotEqual(t, val, nil)

	body := strings.Split(string(val), ",")
	require.Equal(t, len(body), totalCount)
}
