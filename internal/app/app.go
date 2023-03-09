package app

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/menggggggg/go-web-template/internal/app/config"
	"github.com/menggggggg/go-web-template/internal/app/dao"
	"github.com/menggggggg/go-web-template/internal/app/model"
	"github.com/menggggggg/go-web-template/pkg/logger"
	//"github.com/spf13/viper"
	"github.com/robfig/cron"
)

func Init(ctx context.Context) (func(), error) {
	// 初始化配置文件
	config.LoadConfig()

	// 初始化日志
	InitLogger()
	//初始化swagger
	InitSwagger()
	// 初始化服务运行监控
	monitorCleanFunc := InitMonitor(ctx)

	// 初始化依赖注入器
	injector := BuildInjector()

	// 初始化HTTP服务
	httpServerCleanFunc := InitHTTPServer(ctx, injector.Engine)

	InitGen()
	dao.SetDefault(InitGormDB())
	//检测环境
	CheckAndInitCodeENV()
	//定时发送心跳包
	SendHealth()

	return func() {
		httpServerCleanFunc()
		monitorCleanFunc()
	}, nil
}

func CheckAndInitCodeENV() {
	var result []string
	for _, l := range config.C.Language {
		split := strings.Split(l, " ")
		code := execCode(split[0], split[1], "")
		if code == 0 {
			result = append(result, split[0])
		}
	}
	config.C.Info.SupportLanguage = strings.Join(result, ",")
	logger.Debug(strings.Join(result, ","))
	viper.Set("info", config.C.Info)
	viper.WriteConfig()
}
func execCode(executor string, execFilePath string, input string) int {
	cmd := exec.Command(executor, execFilePath)
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalln(err)
	}
	io.WriteString(stdinPipe, input+"\n")
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
	//log.Println(cmd.ProcessState.ExitCode())
	return cmd.ProcessState.ExitCode()
}
func SendHealth() {
	c := cron.New()
	c.AddFunc("*/5 * * * *", func() {
		//获取地址
		serverApi := config.C.API.ManagerServer
		//发送心跳
		configInfo := model.ConfigInfo{
			// Id: config.C.ConfigInfo.Id,
			Name:            config.C.Info.Name,
			SupportLanguage: config.C.Info.SupportLanguage,
			Enabled:         config.C.Info.Enabled,
			URL:             GetIntranetIp() + config.C.HTTP.Addr,
		}
		requestBody, _ := json.Marshal(configInfo)
		logger.Debug("准备发送心跳" + string(requestBody))

		_, err := http.Post(serverApi, "application/json", bytes.NewReader(requestBody))
		if err != nil {
			logger.Error("后端连接失败" + err.Error())
			return
		}
	})
	c.Start()

}
func GetIntranetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("ip:", ipnet.IP.String())
				return ipnet.IP.String()
			}

		}
	}
	return ""
}
func InitSwagger() {
	cmd := exec.Command("swag", "init")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug(out.String())
}

// InitHTTPServer 初始化http服务
func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.C.HTTP
	srv := &http.Server{
		Addr:         cfg.Addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		logger.Infof("HTTP server is running at %s.", cfg.Addr)

		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}

	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
	}
}

// Run 运行服务
func Run(ctx context.Context) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.Infof("Receive signal[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.Infof("Server exit")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
