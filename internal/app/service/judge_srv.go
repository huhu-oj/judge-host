package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/google/wire"
	"github.com/menggggggg/go-web-template/internal/app/config"
	"github.com/menggggggg/go-web-template/internal/app/dao"
	"github.com/menggggggg/go-web-template/internal/app/model"
)

var JudgeSet = wire.NewSet(wire.Struct(new(JudgeSrv), "*"))

type JudgeSrv struct {
}

const (
	cpp    = "cpp"
	java   = ""
	python = "python"
	golang = "go"
)

func (a *JudgeSrv) Judge(ctx context.Context, request *model.AnswerRecord) {
	//设置输入输出
	ioList, err := dao.OjStandardIo.Where(dao.OjStandardIo.ProblemID.Eq(request.ProblemId)).Find()
	if err != nil {
		log.Fatalln("读取数据库失败" + err.Error())
		return
	}
	//获取语言
	language, err := dao.OjLanguage.Where(dao.OjLanguage.ID.Eq(request.LanguageId)).First()
	if err != nil {
		log.Fatalln("获取语言出错" + err.Error())
		return
	}
	//获取执行者
	executor, err := getExecutor(language)
	if err != nil {
		log.Fatalln("获取执行者出错" + err.Error())
		return
	}
	//保存成临时文件夹
	dir, err := createTempDir("code")
	if err != nil {
		log.Fatalln("保存代码文件失败" + err.Error())
		return
	}
	defer func() {
		os.Remove(dir)
	}()
	//保存临时文件
	f1, err := saveTempFile(dir, "code", executor.Ext, request.Code)
	defer func() {
		f1.Close()
		os.Remove(f1.Name())
	}()

	//编译
	//compile(executor)
	//执行 loop
	//判断对错
	for _, standardio := range ioList {
		//执行
		out, stderr, _ := execCode(executor.Cmd, request.Input, insert(executor.Args, len(executor.Args), f1.Name())...)
		if stderr.Len() != 0 {
			//有错误
			request.Error = stderr.String()
			request.ExecuteResultId = 3
			return
		}
		formatOut := strings.ReplaceAll(out.String(), "\n", "")
		//对比输出
		if strings.Compare(formatOut, standardio.Output) == 0 {
			request.PassNum++
		} else {
			//不相等
			request.Error = fmt.Sprintf("输入：%v\n期望输出：%v\n实际输出：%v", standardio.Input, standardio.Output, formatOut)
			request.NotPassNum = len(ioList) - request.PassNum
			request.ExecuteResultId = 2
			return
		}
		request.ExecuteResultId = 1
	}
	//统计结果

}
func createTempDir(tempDirName string) (string, error) {
	workDir, err := os.MkdirTemp("", tempDirName)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	return workDir, nil
}
func getExecutor(language *model.OjLanguage) (*config.Executor, error) {
	//获取语言对应的executor
	for _, executor := range config.C.Executor {
		if executor.Cmd == language.Name {
			return &executor, nil
		}
	}
	return nil, errors.New("找不到对应的执行者")
}

//	func compile(e config.Executor, sourceFilePath string) (string, error) {
//		//var args [] string
//		//执行executor中的编译命令
//		c := e.Compile
//		for c != nil {
//			//append(c.Args,sourceFilePath)
//			_, stderr, err := execCode(c.Cmd, "", c.Args...)
//			if err != nil {
//				log.Fatalln(stderr, err.Error())
//				return "", err
//			}
//			c = c.Compile
//		}
//		//返回可执行文件的路径
//		return "", nil
//	}
func execCompile(e config.Executor) {

}
func (a *JudgeSrv) Test(request *model.AnswerRecord) {

	//获取语言
	language, err := dao.OjLanguage.Where(dao.OjLanguage.ID.Eq(request.LanguageId)).First()
	if err != nil {
		log.Fatalln("获取语言出错" + err.Error())
		return
	}
	//获取执行者
	executor, err := getExecutor(language)
	if err != nil {
		log.Fatalln("获取执行者出错" + err.Error())
		return
	}
	//保存成临时文件
	f1, err := saveTempFile("", "code", executor.Ext, request.Code)
	if err != nil {
		log.Fatalln("保存代码文件失败" + err.Error())
		return
	}
	//编译
	//javac file
	// f2, err1 := compile()
	// if err1 != nil {
	// 	log.Fatalln("生成可执行文件失败"+err1.Error())
	// 	return
	// }
	//执行
	// cmd := exec.Command("go", "run", f1.Name())
	out, stderr, err := execCode(executor.Cmd, request.Input, insert(executor.Args, len(executor.Args), f1.Name())...)
	//包装结果
	println(out.String(), stderr.String())
	request.Log = out.String()
	request.Error = stderr.String()
	//删除缓存文件
	defer func() {
		f1.Close()
		os.Remove(f1.Name())
	}()
}
func insert(slice []string, index int, value string) []string {
	// 创建一个新的 slice，长度为原始 slice 的长度+1
	result := make([]string, len(slice)+1)
	// 将原始 slice 中 index 之前的元素拷贝到新的 slice 中
	copy(result[:index], slice[:index])
	// 将要插入的元素插入到新 slice 的 index 位置上
	result[index] = value
	// 将原始 slice 中 index 之后的元素拷贝到新的 slice 中
	copy(result[index+1:], slice[index:])
	return result
}
func saveTempFile(tempdir string, filename string, ext string, content string) (*os.File, error) {
	file, err := os.CreateTemp(tempdir, fmt.Sprintf("%v.*.%v", filename, ext))
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(file.Name(), []byte(content), 0666)
	if err != nil {
		return nil, err
	}
	return file, nil

}
func execCode(executor string, input string, args ...string) (*bytes.Buffer, *bytes.Buffer, error) {
	fmt.Println(args)
	cmd := exec.Command(executor, args...)
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &out
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	io.WriteString(stdinPipe, input+"\n")
	if err := cmd.Run(); err != nil {
		log.Println("运行", err, stderr.String())
		return &out, &stderr, err
	}
	return &out, &stderr, nil
}
func judge0(standardIos []*model.OjStandardIo, path string) {
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
