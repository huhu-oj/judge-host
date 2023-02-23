package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var CommonSet = wire.NewSet(wire.Struct(new(JudgeAPI), "*"))

// UserAPI ...
type CommonApi struct {
}

// Judge
// @Tags 获取配置
// @Summary 通用
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/config [get]
func (a *JudgeAPI) Config(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {

	})
}
