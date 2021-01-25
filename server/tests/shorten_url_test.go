package tests

import (
	"../router"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)


func TestShortenUrl(t *testing.T){
	b, _ := json.Marshal(map[string]string{"url": "test"})
	t.Run("POST", func(t *testing.T) {
		req := httptest.NewRequest("POST", "http://localhost:8080/", bytes.NewReader(b))
		rec := httptest.NewRecorder()

		mux := router.Router()
		mux.ServeHTTP(rec, req)
		assert.Equal(t, 201, rec.Code)
	})
}
