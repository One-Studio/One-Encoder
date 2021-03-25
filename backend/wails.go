package backend

import (
	"fmt"
	pls "github.com/One-Studio/ptools/pkg"
)

///// wails.go存放backend包对frontend细粒度交互的操作

//设置go变量
func (a *App) SetParam(name string, param string)  {
	//t := a.cfg.Tools[name]
	//t.Param = param
	//fmt.Println("参数", name, param)
}

//测试发送信息
func (a *App) SayHello() string {
	fmt.Println("Hello to Backend!")
	//发信息 事件.信号  Emit
	//a.runtime.Events.Emit("error", "Go发送的错误信息！", 6657)
	return "Hello to Frontend!"
}

func (a *App) Test(str1, str2, str3 string) (result string) {
	result = "Backend has got:\nstr1 = " + str1 + "\nstr2 = " + str2 + "\nstr3 = " + str3
	return
}

//设置进度条
func (a *App) setProgress(progress float64) {
	a.runtime.Events.Emit("SetProgess", progress)
}

//设置日志信息
func (a *App) setLog(log string) {
	a.runtime.Events.Emit("SetLog", log)
}

//设置单行日志
func (a *App) setPerLog(log string) {
	a.runtime.Events.Emit("SetPerLog", log)
}

//设置版本代号
func (a *App) setVersionCode(versionCode string) {
	a.runtime.Events.Emit("SetVersionCode", versionCode)
}

//设置App版本
func (a *App) setAppVersion(appVersion string) {
	a.runtime.Events.Emit("SetAppVersion", appVersion)
}

//通知成功
func (a *App) noticeSuccess(msg string) {
	a.runtime.Events.Emit("NoticeSuccess", msg)
}

//通知错误
func (a *App) noticeError(msg string) {
	a.runtime.Events.Emit("NoticeError", msg)
}

//通知警告
func (a *App) noticeWarning(msg string) {
	a.runtime.Events.Emit("NoticeWarning", msg)
}

//选择文件夹
func (a *App) SelectDirectory() string {
	directory := a.runtime.Dialog.SelectDirectory()
	if !pls.IsFileExisted(directory) {
		_ = pls.WriteFast("./cancel.txt", "取消安装")
		a.noticeError("文件夹不存在或者未选择 ")
		return ""
	}

	return directory
}

//选择文件
func (a *App) SelectFile() string {
	path := a.runtime.Dialog.SelectFile()
	if !pls.IsFileExisted(path) {
		a.noticeError("文件不存在或者未选择 ")
		return ""
	}

	return path
}

//选择文件，有标题
func (a *App) SelectFileTitle(Title string) string {
	path := a.runtime.Dialog.SelectFile(Title)
	if !pls.IsFileExisted(path) {
		a.noticeError("文件不存在或者未选择 ")
		return ""
	}

	return path
}

//选择文件，有标题和过滤文件
func (a *App) SelectFileTitleFilter(Title string, Filter string) string {
	path := a.runtime.Dialog.SelectFile(Title, Filter)
	if !pls.IsFileExisted(path) {
		a.noticeError("文件不存在或者未选择 ")
		return ""
	}

	return path
}

//选择要保存的文件，有标题
func (a *App) SelectSaveFileTitle(Title string) string {
	path := a.runtime.Dialog.SelectSaveFile(Title)

	return path
}