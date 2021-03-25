package backend

import (
	"One-Encoder/backend/config"
	"fmt"
	pls "github.com/One-Studio/ptools/pkg"
	"github.com/wailsapp/wails"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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
		//a.runtime.Events.Emit("SetLog", err)	//TODO 这个时候页面还没有mounted 似乎会弹脚本错误
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

//设置后端 TODO
func (a *App) SetupBackend() string {
	if err := a.cfg.FFmpeg.Install(); err != nil {
		return err.Error()
	}
	a.noticeSuccess("FFmpeg安装/更新成功！")

	if err := a.cfg.FFprobe.Install(); err != nil {
		return err.Error()
	}
	a.noticeSuccess("FFprobe安装/更新成功！")

	if err := a.cfg.X264.Install(); err != nil {
		return err.Error()
	}
	a.noticeSuccess("x264安装/更新成功！")

	if err := a.cfg.X265.Install(); err != nil {
		return err.Error()
	}
	a.noticeSuccess("x265安装/更新成功！")

	if runtime.GOOS == "windows" {
		if err := a.cfg.Pssuspend.Install(); err != nil {
			return err.Error()
		}
		a.noticeSuccess("pssuspend成功！")
	}

	//if err := a.cfg.VapourSynth.Install(); err != nil {
	//	return err.Error()
	//}

	a.cfg.Init = true
	a.noticeSuccess("更新成功！")
	return ""
}

func (a *App) GenerateOutput(input string) (output string) {
	ext := filepath.Ext(input)
	noExt := strings.TrimSuffix(input, ext)
	output = noExt + "_encode" + ext
	for i := 1; pls.IsFileExisted(output); i++ {
		output = noExt + "_encode(" + strconv.Itoa(i) + ")" + ext
	}

	return output
}

//调用工具
//Param: Input Output Param压制参数 Tool用哪个工具
//Return: error
func (a *App) StartEncode(input, output, param, tool string) string {
	//Windows下检查pssuspend
	if runtime.GOOS == "windows" {
		if !a.cfg.Pssuspend.CheckExist() {
			a.noticeWarning("Pssuspend未正确安装，无法暂停压制")
		}
	}

	//工具 参数
	var sig = make(chan rune)
	var command string
	switch tool {
	case "ffmpeg":
		if !a.cfg.FFmpeg.CheckExist() {
			return "ffmpeg工具未正确安装"
		}
		command = a.cfg.FFmpeg.Path + " -i \"" + input + "\" " + param + " \"" + output + "\""
	case "x264":
		if !a.cfg.X264.CheckExist() {
			return "x264工具未正确安装"
		}
		command = a.cfg.X264.Path + " \"" + input + "\" " + param + " -o \"" + output + "\""
	case "x265":
		if !a.cfg.X265.CheckExist() {
			return "x265工具未正确安装"
		}
		command = a.cfg.X265.Path + " \"" + input + "\" " + param + " -o \"" + output + "\""
	case "ffprobe":
		if !a.cfg.FFprobe.CheckExist() {
			return "ffprobe工具未正确安装"
		}
		command = a.cfg.FFprobe.Path + " -v quiet  -print_format json -show_format \"" + input + "\""
	}

	//接受暂停/终止信号量
	go func() {
		a.runtime.Events.On("RealtimeSignal", func(data ...interface{}) {
			fmt.Println("收到信号:", data[0])
			sig <- data[0].(rune)
		})
	}()

	err := pls.ExecRealtimeControl(command, func(line string) {
		a.setProgress(66.57)
		a.setPerLog(line)
	}, sig, a.cfg.Pssuspend.Path)
	if err != nil {
		return err.Error()
	}

	return ""
}

func (a *App) ParseDragFiles() (string, error) {

	return "", nil
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
