package backend

import (
	"One-Encoder/backend/config"
	"fmt"
	pls "github.com/One-Studio/ptools/pkg"
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
	if err = a.cfg.ReadConfig("./config.json"); err != nil {
		a.runtime.Events.Emit("SetLog", err)	//TODO 这个时候页面还没有mounted 似乎会弹脚本错误
		log.Println(err)
		return err
	}

	return nil
}

//Wails结束前
func (a *App) WailsShutdown() {
	//结束前：
	fmt.Println("Wails结束")
	err := a.cfg.SaveConfig("./config.json")
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

//调用工具
//Param: Input Output Param压制参数 Tool用哪个工具
//Return: error
func (a *App) Encode(input, output, param, tool string) error {
	//Windows下检查pssuspend

	//检查工具是否能正常使用

	//工具 参数
	var sig = make(chan rune)
	var command string
	switch tool {
	case "ffmpeg":
		command = a.cfg.FFmpeg.Path + " -i \"" + input + "\" " + param + " \"" + output + "\""
	case "x264":
		command = a.cfg.X264.Path + " \"" + input + "\" " + param + " -o \"" + output + "\""
	case "x265":
		command = a.cfg.X265.Path + " \"" + input + "\" " + param + " -o \"" + output + "\""
	case "ffprobe":
		command = a.cfg.FFprobe.Path + " -v quiet  -print_format json -show_format \"" + input + "\""
	}

	//接受暂停/终止信号量
	go func() {
		a.runtime.Events.On("RealtimeSignal", func(data ...interface{}) {
			fmt.Println("收到信号:", data[0])
		})
	}()

	err := pls.ExecRealtimeControl(command, func(line string) {
		fmt.Println(line)
	}, sig, "")

	return err
}



func (a *App) SelectSrcPath() string {
	return a.SelectFileTitle("选择输入路径")
}

func (a *App) SelectDstPath() string {
	return a.SelectFileTitle("选择输出路径")
}

func (a *App) ParseDragFiles() (string, error) {

	return "", nil
}

func (a *App) checkTools() {

}

func (a *App) StartEncoding(name string) error {
	//t := a.cfg.Tools[name]
	//判断存在
	//if ok, err := tool.IsFileExisted(t.Path); err != nil || !ok {
	//	return errors.New("tool " + name + " does not exist")
	//}

	//path := "C:/Users/Purp1e/go/src/github.com/One-Studio/One-Encoder/build/tools/ffmpeg.exe"
	//input := "C:/Users/Purp1e/Videos/测试.mp4"
	//output := "C:/Users/Purp1e/Desktop/测试One-Encoder.mp4"
	//param := "-vcodec libx264 -crf 20 -preset slow"

	//command := path + " -i " + input + " " + param + " " + output + " -y"
	//arg := "-i " + input + " " + param + " " + output + " -y"
	//得到指令
	//command := t.Path + " " + t.InputPrefix + " " + a.cfg.Input + " " + t.Param + " " + t.OutputPrefix + " " + a.cfg.Output

	//执行
	//tool.CmdRealtime(path, arg, func(progress float64) {
	//	a.setProgress(progress)
	//})
	//if output, err := tool.Cmd(command); err != nil {
	//	a.noticeError(output)
	//	return err
	//}
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
