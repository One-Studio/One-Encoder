package config

/// 设置定义

type API struct {
	win   string
	mac   string
	linux string
}

type tools struct {
	Path            string //工具路径，存在环境变量中则直接用工具名
	Exist           bool   //是否存在
	Version         string //工具版本
	InputPrefix     string //指令中输入媒体前缀，如 -i
	OutputPrefix    string //指令中输出媒体前缀，如 -o
	ProgressPattern string //获取压制进度的正则表达式
	API             API    //官方源API
	CdnAPI          API    //CDN搬运源API
	Param           string //当前压制参数
}

type CFG struct {
	AppVersion string
	SrcPath    string
	DstPath    string
	Tools      map[string]tools
	Init       bool
}
