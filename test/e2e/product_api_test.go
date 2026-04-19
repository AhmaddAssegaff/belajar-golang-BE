package e2e

import (
	"net/http"
	"testing"

	cfg "belajar-go-be/test/e2e/config"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	_, _, cleanup := StartTestServer()
	defer cleanup()

	resp, err := http.Get(cfg.BaseURL + "/products")

	assert.NoError(t, err)

	defer resp.Body.Close()

	assert.Equalf(t, http.StatusOK, resp.StatusCode, "unexpected status code")
}
