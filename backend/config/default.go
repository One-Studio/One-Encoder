package config

/// 默认设置

var defaultCFG = CFG{
	AppVersion: "0.1.0 alpha",
	SrcPath:    "",
	DstPath:    "",
	Tools: map[string]tools{
		"ffmpeg": {
			Path:            "./tools/ffmpeg.exe",
			Exist:           false,
			Version:         "",
			InputPrefix:     "-i",
			OutputPrefix:    "",
			ProgressPattern: "",
			API: API{
				win:   "",
				mac:   "",
				linux: "",
			},
			CdnAPI: API{
				win:   "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Win64@master/api.json",
				mac:   "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Mac-master@master/api.json",
				linux: "https://cdn.jsdelivr.net/gh/One-Studio/FFmpeg-Linux64-master@master/api.json",
			},
			Param: "-vcodec libx264 -crf 17 -preset slower",
		},
		"x264": {
			Path:            "./tools/x264/x264.exe",
			Exist:           false,
			Version:         "",
			InputPrefix:     "",
			OutputPrefix:    "--output",
			ProgressPattern: "",
			API: API{
				win:   "",
				mac:   "",
				linux: "",
			},
			CdnAPI: API{
				win:   "",
				mac:   "",
				linux: "",
			},
			Param: "--crf 17 --preset slower",
		},
		"x265": {
			Path:            "./tools/x265/x265.exe",
			Exist:           false,
			Version:         "",
			InputPrefix:     "",
			OutputPrefix:    "--output",
			ProgressPattern: "",
			API: API{
				win:   "",
				mac:   "",
				linux: "",
			},
			CdnAPI: API{
				win:   "",
				mac:   "",
				linux: "",
			},
			Param: "",
		},
		"mediainfo": {
			Path:            "./tools/mediainfo/mediainfo.exe",
			Exist:           false,
			Version:         "",
			InputPrefix:     "",
			OutputPrefix:    "",
			ProgressPattern: "",
			API: API{
				win:   "",
				mac:   "",
				linux: "",
			},
			CdnAPI: API{
				win:   "",
				mac:   "",
				linux: "",
			},
			Param: "",
		},
	},
	Init: false,
}
