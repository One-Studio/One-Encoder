package backend

import (
	"One-Encoder/backend/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	pls "github.com/One-Studio/ptools/pkg"
	"github.com/wailsapp/wails"
)

///// app.go 存放backend包与frontend交互的大部分操作

//App设置
type App struct {
	runtime *wails.Runtime //初始化Runtime需要
	cfg     config.CFG     //程序设置
	sig     chan rune      //控制压制暂停/继续/结束的channel
	cfgPath string         //配置文件位置
}

//Wails初始化
func (a *App) WailsInit(rt *wails.Runtime) error {
	a.runtime = rt
	//初始化后：
	if runtime.GOOS == "windows" {
		a.cfgPath = "./config.json"
	} else {
		cfgDir, err := os.UserConfigDir()
		if err != nil {
			panic("获取应用配置目录失败: " + err.Error())
		}

		a.cfgPath = cfgDir + "/One Studio/One Encoder/config.json"
	}

	if err := a.cfg.ReadConfig(a.cfgPath); err != nil {
		//a.rt.Events.Emit("SetLog", err)	//TODO 这个时候页面还没有mounted 似乎会弹脚本错误
		log.Println(err)
		return err
	}

	return nil
}

//Wails结束前
func (a *App) WailsShutdown() {
	//结束前：
	fmt.Println("Wails结束")
	err := a.cfg.SaveConfig(a.cfgPath)
	if err != nil {
		log.Println(err)
	}
	return
}

//设置后端 TODO: 启动检查改进
func (a *App) SetupBackend() string {
	//TODO: 测试
	// a.runtime.Events.On("RealtimeSignal", func(data ...interface{}) {
	// 	fmt.Println("收到信号:", data[0])
	// 	a.sig <- data[0].(rune)
	// })

	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err == nil {
		a.setPerLog(dir)
	}

	//FFmpeg检查
	if !a.cfg.FFmpeg.CheckExist() {
		found := false
		if !a.cfg.AutoUpdate {
			//检查环境变量
			found = a.cfg.FFmpeg.CheckEnvPath()
		}

		if !found {
			if err := a.cfg.FFmpeg.Install(); err != nil {
				return err.Error()
			}
		}
	}
	//x264检查
	if !a.cfg.X264.CheckExist() {
		found := false
		if !a.cfg.AutoUpdate {
			//检查环境变量
			found = a.cfg.X264.CheckEnvPath()
		}

		if !found {
			if err := a.cfg.X264.Install(); err != nil {
				return err.Error()
			}
		}
	}
	//x265检查
	if !a.cfg.X265.CheckExist() {
		found := false
		if !a.cfg.AutoUpdate {
			//检查环境变量
			found = a.cfg.X265.CheckEnvPath()
		}

		if !found {
			if err := a.cfg.X265.Install(); err != nil {
				return err.Error()
			}
		}
	}
	//MediaInfo检查
	if !a.cfg.MediaInfo.CheckExist() {
		found := false
		if !a.cfg.AutoUpdate {
			//检查环境变量
			found = a.cfg.MediaInfo.CheckEnvPath()
		}

		if !found {
			if err := a.cfg.MediaInfo.Install(); err != nil {
				return err.Error()
			}
		}
	}
	//Pssuspend检查
	if !a.cfg.Pssuspend.CheckExist() {
		found := false
		if !a.cfg.AutoUpdate {
			//检查环境变量
			found = a.cfg.Pssuspend.CheckEnvPath()
		}

		if !found && runtime.GOOS == "windows" {
			if err := a.cfg.Pssuspend.Install(); err != nil {
				return err.Error()
			}
		}
	}
	//VapourSynth检查
	//if !a.cfg.VapourSynth.CheckExist() {
	//	found := false
	//	if !a.cfg.AutoUpdate {
	//		//检查环境变量
	//		found = a.cfg.VapourSynth.CheckEnvPath()
	//	}
	//
	//	if !found && runtime.GOOS == "windows" {
	//		if err := a.cfg.VapourSynth.Install(); err != nil {
	//			return err.Error()
	//		}
	//	}
	//}

	if a.cfg.Init {
		a.noticeSuccess("工具更新成功！")
	} else {
		a.cfg.Init = true
		a.noticeSuccess("工具安装成功！")
	}

	return ""
}

//根据输入路径自动得出输出路径
func (a *App) GenerateOutput(input string) (output string) {
	ext := filepath.Ext(input)
	noExt := strings.TrimSuffix(input, ext)
	output = noExt + "_encode" + ext
	for i := 1; pls.IsFileExisted(output); i++ {
		output = noExt + "_encode(" + strconv.Itoa(i) + ")" + ext
	}

	return output
}

func (a *App) GetMediaInfo(input string) string {
	if !a.cfg.MediaInfo.CheckExist() {
		a.noticeWarning("MediaInfo未正确安装")
	}

	cmdArgs := []string{a.cfg.MediaInfo.Path}
	cmdArgs = append(cmdArgs, strings.Fields("--OUTPUT=JSON")...)
	cmdArgs = append(cmdArgs, input)

	if output, err := pls.ExecArgs(cmdArgs); err != nil {
		return err.Error()
	} else {
		a.setMediaInfo(output)
		return ""
	}
}

//调用工具 TODO 进度获取 TODO2 暂停/结束功能
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
	params := strings.Fields(param)
	var cmdArgs []string
	switch tool {
	case "ffmpeg":
		if !a.cfg.FFmpeg.CheckExist() {
			return "ffmpeg工具未正确安装"
		}
		cmdArgs = []string{a.cfg.FFmpeg.Path, "-i", input}
		cmdArgs = append(cmdArgs, params...)
		cmdArgs = append(cmdArgs, output, "-y")
	case "x264":
		if !a.cfg.X264.CheckExist() {
			return "x264工具未正确安装"
		}
		cmdArgs = []string{a.cfg.X264.Path, input}
		cmdArgs = append(cmdArgs, params...)
		cmdArgs = append(cmdArgs, "--output", output)
	case "x265":
		if !a.cfg.X265.CheckExist() {
			return "x265工具未正确安装"
		}
		cmdArgs = []string{a.cfg.X265.Path, input}
		cmdArgs = append(cmdArgs, params...)
		cmdArgs = append(cmdArgs, "-o", output)
	}

	//接受暂停/终止信号量 TODO: 暂停/结束功能有bug
	var logs []string
	err := pls.ExecRealtimeControlArgs(cmdArgs, func(line string) {
		a.setProgress(66.57)
		a.setPerLog(line)
		logs = append(logs, line)
	}, a.sig, a.cfg.Pssuspend.Path)
	if err != nil {
		//a.noticeError(err.Error())
		return err.Error()
	}

	//output, err = pls.CMD("chmod u+x .")
	//if err != nil {
	//	a.noticeError(err.Error())
	//}
	//a.noticeWarning(output)
	//err = pls.WriteFast("./log.txt", strings.Join(logs, "\n"))
	//if err != nil {
	//	a.noticeError(err.Error())
	//}
	//err = pls.WriteFast("/Users/purp1e/Desktop/log.txt", strings.Join(logs, "\n"))
	//if err != nil {
	//	a.noticeError(err.Error())
	//}

	return ""
}

//func (a *App) ParseDragFiles() (string, error) {
//
//	return "", nil
//}

//暂停压制
func (a *App) PauseEncoding() {
	pls.Pause(a.sig)
}

//结束压制
func (a *App) QuitEncoding() {
	pls.Quit(a.sig)
}

//继续压制
func (a *App) ResumeEncoding() {
	pls.Resume(a.sig)
}

func (a *App) BrowseLog() (string, error) {

	return "", nil
}

func (a *App) OpenProgramDir() (string, error) {

	return "", nil
}
