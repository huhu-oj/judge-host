// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/menggggggg/go-web-template/internal/app/api"
	"github.com/menggggggg/go-web-template/internal/app/router"
	"github.com/menggggggg/go-web-template/internal/app/service"
)

// Injectors from wire.go:

// BuildInjector 生成注入器
func BuildInjector() *Injector {
	judgeSrv := &service.JudgeSrv{}
	judgeAPI := &api.JudgeAPI{
		JudgeSrv: judgeSrv,
	}
	routerRouter := &router.Router{
		JudgeAPI: judgeAPI,
	}
	engine := InitGinEngine(routerRouter)
	injector := &Injector{
		Engine: engine,
	}
	return injector
}
