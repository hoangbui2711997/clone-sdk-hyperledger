package webserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tkquestionpost/utils"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewGinEngine(t *testing.T) {
	viper.SetDefault("server.mode", "release")
	defer viper.Reset()
	app := NewGinEngine(nil)
	assert.NotNil(t, app)
}

func TestGinBasicAuth(t *testing.T) {
	viper.Set("basic_auth.username", "axiaoxin")
	viper.Set("basic_auth.password", "axiaoxin")
	defer viper.Reset()
	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/get", nil)
	auth := GinBasicAuth()
	auth(c)
	assert.Equal(t, c.Writer.Status(), http.StatusUnauthorized, "request without basic auth should return StatusUnauthorized")
}

func TestPubk2Addr(t *testing.T) {
	addr := utils.GetAddressByPubkey("0254fff0acdb6d6bbb94d811cdd3cc7eb7f443d288037c6120d1ec055a17244900")
	logrus.Infoln(addr)
}
