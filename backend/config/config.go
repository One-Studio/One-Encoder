package config

import (
	"bytes"
	"encoding/json"
	pls "github.com/One-Studio/ptools/pkg"
	"os"
)

//读设置
func (c *CFG) ReadConfig(path string) error {
	//检查文件是否存在
	if exist := pls.IsFileExisted(path); exist {
		//存在则读取文件
		content, err := pls.ReadAll(path)
		if err != nil {
			return err
		}

		//初始化实例并解析JSON
		var CFGInst CFG
		err = json.Unmarshal([]byte(content), &CFGInst) //第二个参数要地址传递
		if err != nil {
			return err
		}

		//检查现有设置，做一定语法上的修正，处理版本更新带来的设置选项的变化 TODO
		//CFGInst, err = c.checkConfig(CFGInst)
		//if err != nil {
		//	return err
		//}

		c.SetCFG(CFGInst)
		return nil
	} else {
		//设置文件不存在则初始化
		c.SetDefCFG()
		return nil
	}
}

//写设置
func (c *CFG) SaveConfig(path string) error {
	//检查文件是否存在
	if exist := pls.IsFileExisted(path); exist {
		//存在则删除文件
		if ok := pls.IsFileExisted(path); ok {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}
	}

	JsonData, err := c.CFG2Json()
	err = pls.WriteFast(path, JsonData)
	if err != nil {
		return err
	}

	return nil
}

//设置转Json字符串
func (c *CFG) CFG2Json() (string, error) {
	JsonData, err := json.Marshal(c) //第二个参数要地址传递
	if err != nil {
		return "", err
	}

	//json写入文件
	var str bytes.Buffer
	_ = json.Indent(&str, JsonData, "", "    ")
	return str.String(), nil
}

//检查设置，更新覆盖部分设置 TODO
func (c *CFG) checkConfig(cfg CFG) (CFG, error) {
	//cfg.AppVersion = defaultCFG.AppVersion

	return cfg, nil
}

//设置CFG参数
func (c *CFG) SetCFG(cfg CFG) {
	c.AppVersion = cfg.AppVersion
	c.Init = cfg.Init
	c.FFmpegRegExp = cfg.FFmpegRegExp
	c.X264RegExp = cfg.X264RegExp
	c.X265RegExp = cfg.X265RegExp
	c.FFmpeg = cfg.FFmpeg
	c.FFprobe = cfg.FFprobe
	c.X264 = cfg.X264
	c.X265 = cfg.X265
	c.VapourSynth = cfg.VapourSynth
	c.Pssuspend = cfg.Pssuspend
	c.FFmpegParam = cfg.FFmpegParam
	c.FFprobeParam = cfg.FFprobeParam
	c.X264Param = cfg.X264Param
	c.X265Param = cfg.X265Param
}
