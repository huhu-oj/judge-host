package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime"
	"sync"

	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/dao"
	"github.com/menggggggg/go-web-template/internal/app/model"
)

var JudgeSet = wire.NewSet(wire.Struct(new(JudgeSrv), "*"))

type JudgeSrv struct {
}

func (a *JudgeSrv) Judge(ctx context.Context, request int64) ([]*model.OjStandardIo, error) {
	ioList, err := dao.OjStandardIo.Where(dao.OjStandardIo.ProblemID.Eq(request)).Find()
	if err != nil {
		return nil, err
	}
	judge0(ioList,"main.go")
	
	return nil, nil
}

func judge0(standardIos []*model.OjStandardIo,path string) {
	// 答案错误的channel
	WA := make(chan int)
	// 超内存的channel
	OOM := make(chan int)
	// 编译错误的channel
	CE := make(chan int)
	// 答案正确的channel
	AC := make(chan int)
	// 非法代码的channel
	EC := make(chan struct{})

	// 通过的个数
	passCount := 0
	var lock sync.Mutex
	// 提示信息
	var msg string
	for _, testCase := range standardIos {
		testCase := testCase
			go func() {
				cmd := exec.Command("go", "run", path)
				cmd.Dir = "C:\\Users\\Administrator\\Desktop\\go-study\\gin-gorm-oj\\internal\\code\\code-user\\"
				var out, stderr bytes.Buffer
				cmd.Stderr = &stderr
				cmd.Stdout = &out
				stdinPipe, err := cmd.StdinPipe()
				if err != nil {
					log.Fatalln(err)
				}
				io.WriteString(stdinPipe, testCase.Input+"\n")

				var bm runtime.MemStats
				runtime.ReadMemStats(&bm)
				if err := cmd.Run(); err != nil {
					log.Println(err, stderr.String())
					if err.Error() == "exit status 2" {
						msg = stderr.String()
						CE <- 1
						return
					}
				}
				var em runtime.MemStats
				runtime.ReadMemStats(&em)

				// 答案错误
				if testCase.Output != out.String() {
					WA <- 1
					return
				}
				lock.Lock()
				passCount++
				if passCount == len(standardIos) {
					AC <- 1
				}
				lock.Unlock()
			}()
	}
		
	select {
	case <-EC:
		msg = "无效代码"
		// sb.Status = 6
	case <-WA:
		msg = "答案错误"
		// sb.Status = 2
	case <-OOM:
		msg = "运行超内存"
		// sb.Status = 4
	case <-CE:
		// sb.Status = 5
	case <-AC:
		msg = "答案正确"
		// sb.Status = 1
	default:
		if passCount == len(standardIos) {
			// sb.Status = 1
			msg = "答案正确"
		} else {
			// sb.Status = 3
			msg = "运行超时"
		}
	}
	fmt.Println(msg)
}