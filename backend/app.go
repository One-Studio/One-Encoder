package backend

import (
	"One-Encoder/backend/config"
	"fmt"
	"github.com/wailsapp/wails"
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
	fmt.Println("Wails初始化")
	config.ReadConfig(a.cfg)

	return nil
}

//Wails结束前
func (a *App) WailsShutdown() {
	//结束前：
	fmt.Println("Wails结束")
	config.SaveConfig(a.cfg)

	return
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

func (a *App) StartEncoding() (string, error) {

	return "", nil
}

func (a *App) PauseEncoding() (string, error) {

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



//具体的方法返回只能是 x 或者 (x, error) 下面是cantor中使用映射json的方式处理返回值
//// GetConfig 获取 git 配置和版本信息
//func (a *App) GetConfig() *configs.Resp {
//	resp := map[string]interface{}{
//		"config": a.Git,
//		"version": map[string]interface{}{
//			"current": configs.Version,
//			"last":    a.Git.LastVersion(),
//		},
//	}
//	a.Log.Info("GetConfig content: ", crypto.JsonEncode(resp))
//	return tools.Success(resp)
//}
//
//// config包：Resp ...
//type Resp struct {
//	Code int         `json:"code"`
//	Msg  string      `json:"msg"`
//	Data interface{} `json:"data"`
//}
