package main

import (
	"One-Encoder/backend"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	//绑定js和css
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	//创建app
	app := wails.CreateApp(&wails.AppConfig{
		Width:     600,
		Height:    360,
		Title:     "One Encoder",
		JS:        js,
		CSS:       css,
		Colour:    "#eaeaea",
		Resizable: false,
	})

	//绑定后端&运行
	app.Bind(&backend.App{})
	app.Run()
}
