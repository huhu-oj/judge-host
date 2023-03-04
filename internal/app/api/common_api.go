package api

import (
	"encoding/json"
	"github.com/menggggggg/go-web-template/internal/app/model"
	"github.com/menggggggg/go-web-template/pkg/logger"
	"github.com/spf13/viper"
	"io"
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
	//接受后端新的配置
	responseConfig := model.ConfigInfo{}
	data, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(data, &responseConfig)
	logger.Debug(responseConfig)
	viper.Set("info", responseConfig)
	viper.WriteConfig()
	viper.WatchConfig()
	c.JSON(http.StatusOK, gin.H{})
}
func (a *CommonAPI) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
