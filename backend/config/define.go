package config

/// 设置定义

type API struct {
	win   string
	mac   string
	linux string
}

type CFG struct {
	version       string
	srcPath       string
	dstPath       string
	param         string
	api           API
	ffmpegPath    string
	ffmpegVersion string
	ffmpegRegExp  string
	Init          bool
}

