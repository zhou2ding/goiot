package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/gogf/gf/v2/os/gcmd"
	"goiot/internal/pkg/conf"
	"goiot/internal/pkg/version"
)

func main() {
	for k, v := range gcmd.GetOptAll() {
		switch k {
		case "h", "help":
			figure.NewFigure("Linkview", "", false).Print()
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

}
