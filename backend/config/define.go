package config

import (
	pls "github.com/One-Studio/ptools/pkg"
)

/// 设置定义

type CFG struct {
	AppVersion   string   //当前App版本号
	Init         bool     //是否已经初始化
	AutoUpdate   bool     //是否在程序开始的时候自动检查更新
	FFmpegRegExp string   //ffmpeg提取进度用正则表达式
	X264RegExp   string   //X264提取进度用正则表达式
	X265RegExp   string   //X265提取进度用正则表达式
	FFmpeg       pls.Tool //ffmpeg工具
	FFprobe      pls.Tool //ffprobe工具
	X264         pls.Tool //x264工具
	X265         pls.Tool //x265工具
	Pssuspend    pls.Tool //windows挂起进程所需工具
	VapourSynth  pls.Tool //VS工具 TODO 考虑如何使用
	FFmpegParam  []string //ffmpeg参数/预设
	FFprobeParam string   //ffprobe获取媒体信息的参数
	X264Param    []string //x264参数/预设
	X265Param    []string //x265参数/预设
}
