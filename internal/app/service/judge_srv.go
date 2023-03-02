package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
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
const (
	cpp = "cpp"
	java = "java"
	python = "python"
	golang = "go"
)
func (a *JudgeSrv) Judge(ctx context.Context, request *model.AnswerRecord) ([]*model.OjStandardIo, error) {
	//设置输入输出
	ioList, err := dao.OjStandardIo.Where(dao.OjStandardIo.ProblemID.Eq(request.ProblemId)).Find()
	if err != nil {
		return nil, err
	}
	judge0(ioList,"main.go")

	//编译
	// complie()
	//执行 loop
		//判断对错
	
	//统计结果

	//返回结果
	return nil, nil
}
func  (a *JudgeSrv) Test(request *model.AnswerRecord){
	//保存成临时文件
	f1, err := saveTempFile("","code","py",request.Code)
	if err != nil {
		log.Fatalln("保存代码文件失败"+err.Error())
		return 
	}

	//编译
	
	//执行
	// cmd := exec.Command("go", "run", f1.Name())
	out,stderr := execCode(python,f1.Name(),request.Input)
	//包装结果
	request.Log = out.String()
	request.Error = stderr.String()
	//删除缓存文件
	defer func() {
		f1.Close()
        os.Remove(f1.Name())
	}()
}

func saveTempFile(tempdir string,filename string, ext string, content string) (*os.File,error) {
	file, err := ioutil.TempFile(tempdir, fmt.Sprintf("%v.*.%v",filename,ext))
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(file.Name(), []byte(content), 0666)
	if err != nil {
		return nil, err
	}
	return file, nil

}
func execCode(executor string, execFilePath string, input string) (bytes.Buffer,bytes.Buffer){
	cmd := exec.Command(executor, execFilePath)
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	io.WriteString(stdinPipe, input+"\n")
	if err := cmd.Run(); err != nil {
		log.Println(err, stderr.String())
	}
	return out,stderr
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
				cmd.Dir = "cmd/go-web-template/main.go"
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
