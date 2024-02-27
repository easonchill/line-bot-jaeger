package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCallback(t *testing.T) {
	r := gin.Default()
	SetupRouter(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sendMassage", bytes.NewBufferString(`{"userid":"test001", "text":"Hello, World!"}`))
	req.Header.Set("Content-Type", "multipart/form-data")
	r.ServeHTTP(w, req)

	assert.EqualValues(t, w.Code, http.StatusInternalServerError)
}
