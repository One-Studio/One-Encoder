package config

import (
	pls "github.com/One-Studio/ptools/pkg"
	"runtime"
)

/// 默认设置
func (c *CFG) SetDefCFG()  {
	c.AppVersion = "v0.1.0"
	c.Init = false
	c.X264Param = nil
	c.X265Param = nil
	c.Current = 0
	switch runtime.GOOS {
	case "windows":
		//windows下参数
		c.FFmpeg = pls.Tool{

		}
		c.X264 = pls.Tool{

		}
		c.X265 = pls.Tool{

		}
	case "darwin":
		//macos下参数
		c.FFmpeg = pls.Tool{

		}
		c.X264 = pls.Tool{

		}
		c.X265 = pls.Tool{

		}
	default:
		//linux等其他系统下参数
		c.FFmpeg = pls.Tool{

		}
		c.X264 = pls.Tool{

		}
		c.X265 = pls.Tool{

		}
	}
		//	"ffmpeg": {
		//		Path:            "./tools/ffmpeg.exe",
		//		Exist:           false,
		//		Version:         "",
		//		InputPrefix:     "-i",
		//		OutputPrefix:    "",
		//		ProgressPattern: "",
		//		CdnAPI: API{
		//			win:   "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Win64@master/api.json",
		//			mac:   "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Mac-master@master/api.json",
		//			linux: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Linux64-master@master/api.json",
		//		},
		//		Param: "-vcodec libx264 -crf 17 -preset slower",
		//	},
		//	"x264": {
		//		Path:            "./tools/x264/x264.exe",
		//		Exist:           false,
		//		Version:         "",
		//		InputPrefix:     "",
		//		OutputPrefix:    "--output",
		//		ProgressPattern: "",
		//		Param: "--crf 17 --preset slower",
		//	},
		//	"x265": {
		//		Path:            "./tools/x265/x265.exe",
		//		Exist:           false,
		//		Version:         "",
		//		InputPrefix:     "",
		//		OutputPrefix:    "--output",
		//		ProgressPattern: "",
}
