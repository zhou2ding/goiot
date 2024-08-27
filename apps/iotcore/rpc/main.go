package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gconv"
	"goiot/apps/iotcore/rpc/internal/config"
	"goiot/pkg/conf"
	"goiot/pkg/logger"
	"goiot/pkg/service"
	"goiot/pkg/utils"
	"goiot/pkg/version"
	"os"
)

func main() {
	for k, v := range gcmd.GetOptAll() {
		switch k {
		case "h", "help":
			figure.NewFigure("goiot", "", false).Print()
			return
		case "v", "version":
			version.PrintVersion()
			return
		case "c", "conf", "config":
			conf.InitConf(v)
		case "f":
			conf.InitRpcConf(v)
		}
	}

	configPath := "./configs/iotcore.yaml"
	if conf.Conf == nil {
		conf.InitConf(configPath)
	}
	rpcConfigPath := "./configs/iotcore-rpc.yaml"
	if config.RpcConf == nil {
		config.InitRpcConf(rpcConfigPath)
	}

	logger.InitLogger("iotcore-rpc")

	path, _ := os.Getwd()
	p := &Program{}
	sConfig := &service.Config{
		Name:             config.RpcConf.Name,
		DisplayName:      config.RpcConf.Name,
		WorkingDirectory: path,
	}
	s, err := service.New(p, sConfig)
	if err != nil {
		logger.Logger.Errorf("create new service error. %s", err)
		return
	}

	cmd := gconv.String(gcmd.GetArg(1))
	switch cmd {
	case "help":
		figure.NewFigure("goiot", "", false).Print()
		return
	case "version":
		version.PrintVersion()
		return
	case "install", "uninstall", "start", "stop", "restart":
		if err = service.Control(s, cmd); err != nil {
			logger.Logger.Errorf("%s\n", err)
			fmt.Printf("%s %s failed.\n", sConfig.DisplayName, cmd)
			utils.PauseExit()
		} else {
			fmt.Printf("%s %s success.\n", sConfig.DisplayName, cmd)
		}
		return
	default:
		if e := s.Run(); e != nil {
			logger.Logger.Errorf("service run error. %s", e)
			utils.PauseExit()
		}
	}
}
