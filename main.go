package main

import (
	"One-Encoder/backend"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"runtime"
)

func main() {
	//TODO 启动前读取CFG，包括w, h, scale

	w, h := 660, 408-14
	//无边框窗口特性 wails v2.0之前解决固定窗口高度问题 TODO linux 高度测试
	if runtime.GOOS == "darwin" {
		//h += 28
	}

	//绑定js和css
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	//创建app
	app := wails.CreateApp(&wails.AppConfig{
		Width:     w,
		Height:    h,
		Title:     "One Encoder",
		JS:        js,
		CSS:       css,
		Colour:    "#FEFEFE",
		Resizable: false,
	})

	//绑定后端&运行
	app.Bind(&backend.App{})
	_ = app.Run()
}
