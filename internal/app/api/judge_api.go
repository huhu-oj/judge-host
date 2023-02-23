package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/service"
	"github.com/menggggggg/go-web-template/pkg/errors"
)

var JudgeSet = wire.NewSet(wire.Struct(new(JudgeAPI), "*"))

// UserAPI ...
type JudgeAPI struct {
	JudgeSrv *service.JudgeSrv
}

// Judge
// @Tags 判题
// @Summary 判题
// @Param problemId query int false "问题id"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/judge [get]
func (a *JudgeAPI) Judge(c *gin.Context) {
	request, _ := strconv.ParseInt(c.Query("problemId"), 10, 64)
	resp, err := a.JudgeSrv.Judge(c, request)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.WrapWithInternalServerError(err))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// func (a *JudgeAPI) Test(c *gin.Context) {
// 	request := model.OjStandardIo{}
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.AbortWithError(http.StatusBadRequest, errors.WrapWithBadRequest(err))
// 		return
// 	}

// 	resp, err := a.JudgeSrv.Get(c, &request)
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, errors.WrapWithInternalServerError(err))
// 		return
// 	}
// 	println(resp.ID)
// 	println(resp.Input)
// 	println(resp.Output)
// 	println(resp.ProblemID)

// 	c.JSON(http.StatusOK, resp)
// }
