package httpclient

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := Config{Timeout: 10}
	want := &httpClient{
		client: &http.Client{
			Transport: &http.Transport{},
			Timeout:   time.Duration(cfg.Timeout) * time.Second,
		},
	}

	got := New(cfg)
	assert.IsType(t, want, got)
}

func Test_httpClient_Get(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected GET request")
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"), "Unexpected Content-Type header")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET response"}`))
	}))
	defer mockServer.Close()

	client := New(Config{})
	url := mockServer.URL

	resp, err := client.Get(url, http.Header{"Content-Type": []string{"application/json"}})
	assert.NoError(t, err, "Expected no error for GET request")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Error reading response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code")
	assert.Equal(t, `{"message": "GET response"}`, string(body), "Unexpected response body")

	url = "://invalid-url"
	resp, err = client.Get(url, http.Header{"Content-Type": []string{"application/json"}})
	assert.Errorf(t, err, "Expected error for GET request")
}

func Test_httpClient_Post(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected POST request")
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"), "Unexpected Content-Type header")
		body, _ := io.ReadAll(r.Body)
		assert.Equal(t, `{"data": "POST data"}`, string(body), "Unexpected request body")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "POST response"}`))
	}))
	defer mockServer.Close()

	client := New(Config{})
	url := mockServer.URL
	b := bytes.NewBuffer([]byte(`{"data": "POST data"}`))

	resp, err := client.Post(url, http.Header{"Content-Type": []string{"application/json"}}, b)
	assert.NoError(t, err, "Expected no error for POST request")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Error reading response body")
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Unexpected status code")
	assert.Equal(t, `{"message": "POST response"}`, string(body), "Unexpected response body")

	url = "://invalid-url"
	resp, err = client.Post(url, http.Header{"Content-Type": []string{"application/json"}}, b)
	assert.Errorf(t, err, "Expected error for POST request")
}

func Test_httpClient_Put(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method, "Expected PUT request")
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"), "Unexpected Content-Type header")
		body, _ := io.ReadAll(r.Body)
		assert.Equal(t, `{"data": "PUT data"}`, string(body), "Unexpected request body")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "PUT response"}`))
	}))
	defer mockServer.Close()

	client := New(Config{})
	url := mockServer.URL
	b := bytes.NewBuffer([]byte(`{"data": "PUT data"}`))

	resp, err := client.Put(url, http.Header{"Content-Type": []string{"application/json"}}, b)
	assert.NoError(t, err, "Expected no error for PUT request")
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err, "Error reading response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Unexpected status code")
	assert.Equal(t, `{"message": "PUT response"}`, string(body), "Unexpected response body")

	url = "://invalid-url"
	resp, err = client.Put(url, http.Header{"Content-Type": []string{"application/json"}}, b)
	assert.Errorf(t, err, "Expected error for PUT request")
}

func Test_httpClient_Delete(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method, "Expected DELETE request")
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"), "Unexpected Content-Type header")
		body, _ := io.ReadAll(r.Body)
		assert.Equal(t, `{"data": "DELETE data"}`, string(body), "Unexpected request body")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer mockServer.Close()

	client := New(Config{})
	url := mockServer.URL
	b := bytes.NewBuffer([]byte(`{"data": "DELETE data"}`))

	resp, err := client.Delete(url, http.Header{"Content-Type": []string{"application/json"}}, b)
	assert.NoError(t, err, "Expected no error for DELETE request")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNoContent, resp.StatusCode, "Unexpected status code")

	url = "://invalid-url"
	resp, err = client.Delete(url, http.Header{"Content-Type": []string{"application/json"}}, b)
	assert.Error(t, err, "Expected error for DELETE request")
}
