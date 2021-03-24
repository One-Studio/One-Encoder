package config

import pls "github.com/One-Studio/ptools/pkg"

/// 设置定义

type CFG struct {
	AppVersion string   //当前App版本号
	Input      string   //输入
	Output     string   //输出
	Init       bool     //是否已经初始化
	FFmpeg     pls.Tool //ffmpeg工具
	X264       pls.Tool //x264工具
	X265       pls.Tool //x265工具
	FFParam    []string //ffmpeg参数/预设
	X264Param  []string //x264参数/预设
	X265Param  []string //x265参数/预设
	Current    int      //当前使用的预设
}
