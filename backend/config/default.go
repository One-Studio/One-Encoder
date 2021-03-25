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
			Name: "ffmpeg",
			Path: "./tools/ffmpeg.exe",
			TakeOver: true,
			VersionApi: "https://www.gyan.dev/ffmpeg/builds/release-version",
			VersionApiCDN: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Win64@master/version",
			DownloadLink: "https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-essentials.7z",
			DownloadLinkCDN: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Win64@master/dist/ffmpeg-release-essentials.7z",
			IsGitHub: false,
			IsCLI: true,
			VersionRegExp: "ffmpeg version (\\S+)-essentials_build-www.gyan.dev",
			Fetch: "ffmpeg.exe",
		}
		c.FFprobe = pls.Tool{
			Name: "ffprobe",
			Path: "./tools/ffprobe.exe",
			TakeOver: true,
			VersionApi: "https://www.gyan.dev/ffmpeg/builds/release-version",
			VersionApiCDN: "",
			DownloadLink: "https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-essentials.7z",
			DownloadLinkCDN: "",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "ffprobe.exe",
		}
		c.X264 = pls.Tool{
			Name: "x264",
			Path: "./tools/x264.exe",
			TakeOver: true,
			VersionApi: "",
			VersionApiCDN: "https://cdn.jsdelivr.net/gh/One-Studio/x264-t_mod-posix@master/version",
			DownloadLink: "",
			DownloadLinkCDN: "https://cdn.jsdelivr.net/gh/One-Studio/x264-t_mod-posix@master/dist/x264-t_mod-posix.7z",
			GithubRepo: "jpsdr/x264",
			IsGitHub: true,
			IsCLI: true,
			Fetch: "x264_64.exe",
		}
		c.X265 = pls.Tool{
			Name: "x265",
			Path: "./tools/x265.exe",
			TakeOver: true,
			VersionApi: "",
			VersionApiCDN: "https://cdn.jsdelivr.net/gh/One-Studio/x265-Yuuki@master/version",
			DownloadLink: "",
			DownloadLinkCDN: "https://cdn.jsdelivr.net/gh/One-Studio/x265-Yuuki@master/dist/x265-Yuuki.7z",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "x265-gcc-multilib-full.exe",
		}
		c.Pssuspend = pls.Tool{
			Name: "pssuspend",
			Path: "./tools/pssuspend.exe",
			TakeOver: false,
			VersionApi: "",
			VersionApiCDN: "",
			DownloadLink: "",
			DownloadLinkCDN: "",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "",
		}
	case "darwin":
		//macos下参数
		//usr, err := user.Current()
		//if err != nil {
		//	panic("获取应用配置目录失败: " + err.Error())
		//}
		//path := fmt.Sprintf("%s/.%s", usr.HomeDir, "One Encoder")
		c.FFmpeg = pls.Tool{
			Name: "ffmpeg",
			Path: "ffmpeg",
			//Path: path + "/tools/ffmpeg",
			TakeOver: false,
			//TakeOver: true,
			VersionApi: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Mac-master@master/version",	//TODO 官方version api需要解析
			VersionApiCDN: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Mac-master@master/version",
			DownloadLink: "https://evermeet.cx/ffmpeg/get",
			DownloadLinkCDN: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Mac-master@master/dist/ffmpeg-mac-master.7z",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "ffmpeg",
		}
		c.FFprobe = pls.Tool{
			Name: "ffprobe",
			Path: "ffprobe",
			//Path: path + "/tools/ffprobe",
			TakeOver: false,
			//TakeOver: true,
			VersionApi: "https://evermeet.cx/ffmpeg/info/binary/version",
			VersionApiCDN: "",
			DownloadLink: "https://evermeet.cx/ffmpeg/get/ffprobe",
			DownloadLinkCDN: "",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "ffprobe",
		}
		c.X264 = pls.Tool{
			Name: "x264",
			Path: "x264",
			//Path: "./tools/x264",
			TakeOver: false,
			VersionApi: "",
			VersionApiCDN: "",
			DownloadLink: "",
			DownloadLinkCDN: "",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "x264",
		}
		c.X265 = pls.Tool{
			Name: "x265",
			Path: "x265",
			//Path: "./tools/x265",
			TakeOver: false,
			VersionApi: "",
			VersionApiCDN: "",
			DownloadLink: "",
			DownloadLinkCDN: "",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "x265",
		}
	default:
		//linux等其他系统下参数
		c.FFmpeg = pls.Tool{
			Name: "ffmpeg",
			Path: "./tools/ffmpeg",
			TakeOver: true,
			VersionApi: "",
			VersionApiCDN: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Linux64-master@master/version",
			DownloadLink: "",
			DownloadLinkCDN: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Linux64-master@master/dist/ffmpeg-linux64-master.7z",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "ffmpeg",
		}
		c.FFprobe = pls.Tool{
			Name: "ffmpeg",
			Path: "./tools/ffprobe",
			TakeOver: true,
			VersionApi: "",
			VersionApiCDN: "",
			DownloadLink: "",
			DownloadLinkCDN: "",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "ffprobe",
		}
		c.X264 = pls.Tool{
			Name: "x264",
			Path: "x264",
			//Path: "./tools/x264",
			TakeOver: false,
			VersionApi: "",
			VersionApiCDN: "",
			DownloadLink: "",
			DownloadLinkCDN: "",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "x264",
		}
		c.X265 = pls.Tool{
			Name: "x265",
			Path: "x265",
			//Path: "./tools/x265",
			TakeOver: false,
			VersionApi: "",
			VersionApiCDN: "",
			DownloadLink: "",
			DownloadLinkCDN: "",
			IsGitHub: false,
			IsCLI: true,
			Fetch: "x265",
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
