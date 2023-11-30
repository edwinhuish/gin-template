package core

import (
	"fmt"
	"net/http"
	"time"

	"gin-template/global"
	"gin-template/routers"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func InitHttp() {
	if global.CONFIG.System.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	Router := routers.InitRouters()
	// Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	serv := &http.Server{
		Addr:           address,
		Handler:        Router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info("server run success on ", zap.String("address", address))

	global.LOG.Error(serv.ListenAndServe().Error())
}
