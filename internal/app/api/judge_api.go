package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/model"
	"github.com/menggggggg/go-web-template/internal/app/service"
	"net/http"
)

var JudgeSet = wire.NewSet(wire.Struct(new(JudgeAPI), "*"))

// UserAPI ...
type JudgeAPI struct {
	JudgeSrv *service.JudgeSrv
}

// Judge
// @Tags 判题
// @Summary 判题
// @Param AnswerRecord body model.AnswerRecord false "提交记录"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/judge [post]
func (a *JudgeAPI) Judge(c *gin.Context) {
	// request, _ := strconv.ParseInt(c.Query("problemId"), 10, 64)
	answerRecord := &model.AnswerRecord{}
	c.BindJSON(&answerRecord)
	// fmt.Println(answerRecord)
	a.JudgeSrv.Judge(c, answerRecord)
	c.JSON(http.StatusOK, answerRecord)
}

// Test
// @Tags 自定义判题
// @Summary 自定义判题
// @Param AnswerRecord body model.AnswerRecord false "提交记录"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /api/v1/test [post]
func (a *JudgeAPI) Test(c *gin.Context) {
	// request, _ := strconv.ParseInt(c.Query("problemId"), 10, 64)
	answerRecord := &model.AnswerRecord{}
	c.BindJSON(&answerRecord)
	// fmt.Println(answerRecord)
	a.JudgeSrv.Test(answerRecord, c)

}
