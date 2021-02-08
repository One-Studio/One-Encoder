package backend

import (
	"One-Encoder/backend/config"
	tool "One-Encoder/backend/tool"
	"fmt"
	"github.com/wailsapp/wails"
	"log"
)

///// app.go 存放backend包与frontend交互的大部分操作

//App设置
type App struct {
	runtime *wails.Runtime //初始化Runtime需要
	cfg     config.CFG
	//zoom
}

//Wails初始化
func (a *App) WailsInit(runtime *wails.Runtime) error {
	a.runtime = runtime
	//初始化后：
	var err error
	if a.cfg, err = config.ReadConfig("./config.json"); err != nil {
		a.runtime.Events.Emit("SetLog", err)
		log.Println(err)
		return err
	}

	return nil
}

//Wails结束前
func (a *App) WailsShutdown() {
	//结束前：
	fmt.Println("Wails结束")
	err := config.SaveConfig(a.cfg, "./config.json")
	if err != nil {
		log.Println(err)
	}
	return
}

//设置前端变量
func (a *App) SetVar() {
	//a.setAppVersion(a.cfg.AppVersion)
	//a.setVersionCode(a.cfg.VersionCode)
}

func (a *App) SelectSrcPath() (string, error) {

	return "", nil
}

func (a *App) SelectDstPath() (string, error) {

	return "", nil
}

func (a *App) ParseDragFiles() (string, error) {

	return "", nil
}

func (a *App) StartEncoding(toolSelected string) error {
	//得到指令
	command := "1"

	//执行
	if output, err := tool.Cmd(command); err != nil {
		a.noticeError(output)
		return err
	}
	return nil
}

func (a *App) PauseEncoding() (string, error) {

	return "", nil
}
func (a *App) QuitEncoding() (string, error) {

	return "", nil
}

func (a *App) CheckUpdate() error {

	return nil
}

func (a *App) BrowseLog() (string, error) {

	return "", nil
}

func (a *App) OpenProgramDir() (string, error) {

	return "", nil
}
