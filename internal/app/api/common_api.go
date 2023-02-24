package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var CommonSet = wire.NewSet(wire.Struct(new(CommonAPI), "*"))

// UserAPI ...
type CommonAPI struct {
}

// Config
// @Tags 获取配置
// @Summary 通用
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/config [get]
func (a *CommonAPI) Config(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {

	})
}
func (a *CommonAPI) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {

	})
}
