package tool

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/gen2brain/go-unarr"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

//打开文件和读内容 利用io/ioutil
func ReadAll(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	//对内容的操作
	//ReadFile返回的是[]byte字节切片，要用string()方法转变成字符串
	//去除内容结尾的换行符
	str := strings.TrimRight(string(content), "\n")
	return str, nil
}

//文件写入 先清空再写入 利用ioutil
func WriteFast(filePath string, content string) error {
	dir, _ := path.Split(filePath)
	exist, err := IsFileExisted(dir)
	if err != nil {
		return err
	} else if exist == false {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(filePath, []byte(content), 0666)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//判断文件/文件夹是否存在
func IsFileExisted(path string) (bool, error) {
	//返回 true, nil = 存在
	//返回 false, nil = 不存在
	//返回 _, !nil = 位置错误，无法判断
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//利用HTTP Get请求获得数据json
func GetHttpData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	//body, err := resp.Js
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	_ = resp.Body.Close()

	return string(data), nil
}

//下载文件 (下载地址，存放位置)
func DownloadFile(url string, location string) error {
	//利用HTTP下载文件并读取内容给data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		errorInfo := "http failed, check if file exists, HTTP Status Code:" + strconv.Itoa(resp.StatusCode)
		return errors.New(errorInfo)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_ = resp.Body.Close()

	//确保下载位置存在
	_, fileName := path.Split(url)
	ok, err := IsFileExisted(location)
	if err != nil {
		return err
	} else if ok == false {
		err := os.Mkdir(location, os.ModePerm)
		if err != nil {
			return err
		}
	}
	//文件写入 先清空再写入 利用ioutil
	err = ioutil.WriteFile(location+"/"+fileName, data, 0666)
	if err != nil {
		return err
	} else {
		return nil
	}
}

//判断是不是non-ASCII
func IsNonASCII(str string) bool {
	re := regexp.MustCompile("[[:^ascii:]]")
	return re.MatchString(str)
	//var count int
	//for _, v := range str {
	//	if unicode.Is(unicode.Han, v) {
	//		count++
	//		break
	//	}
	//}
	//return count > 0
}

//解压zip 7z rar tar
func Decompress(from string, to string) error {
	a, err := unarr.NewArchive(from)
	if err != nil {
		return err
	}
	defer a.Close()

	_, err = a.Extract(to)
	if err != nil {
		return err
	}

	return nil
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//规格化路径
func FormatPath(s string) string {
	if strings.HasPrefix(s, ".") {
		s = strings.Replace(s, ".", getCurrentDirectory(), 1)
	}
	s = strings.Replace(s, "\\", "/", -1)

	return strings.TrimRight(s, "\\")
}

//复制文件夹
func CopyDir(from string, to string) error {
	from = FormatPath(from)
	to = FormatPath(to)

	//确保目标路径存在，否则复制报错exit status 4
	exist, err := IsFileExisted(to)
	if err != nil {
		return err
	} else if exist == false {
		err := os.Mkdir(to, os.ModePerm)
		if err != nil {
			return err
		}
	}
	var out string
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("xcopy", from, to, "/I", "/E", "/Y", "/R")
	} else {
		cmd = exec.Command("cp", "-R", from, to)
	}

	//if runtime.GOOS == "windows" {
	//	out, err = Cmd("xcopy /I /E /Y " + strconv.Quote(from) + " " + strconv.Quote(to))
	//} else {
	//	out, err = Cmd("cp -R " + from + " " + to)
	//}
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Println(out, err)
	}
	return err
}

//执行一次command指令
func Cmd(command string) (string, error) {
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		c := exec.Command("cmd.exe", "/c", command)
		out, err = c.CombinedOutput()
	} else {
		c := exec.Command("/bin/bash", "-c", command)
		out, err = c.CombinedOutput()
	}
	//cmd.Args = a
	return string(out), err
}

func CmdRealtime(path string, arg string, method func(progress float64))  {
	//command := "./x264 /Users/purp1e/vd/in.mp4 --crf 18 --preset 3 -o /Users/purp1e/vd/outx264.mp4"
	//command := "ffmpeg -i /Users/purp1e/vd/in.mp4 -vcodec libx264 -crf 18 -preset 3 -acodec copy /Users/purp1e/vd/out3.mp4 -y"
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		//cmd = exec.Command("cmd.exe", "/c", command)
		args := strings.Split(arg, " ")
		cmd = exec.Command(path, args...)

	} else {
		cmd = exec.Command("/bin/bash", path + " " + arg)
		//cmd = exec.Command("/bin/bash", "-c", command)
	}

	//隐藏黑框
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()
	stdin, _ := cmd.StdinPipe()
	defer stdin.Close()
	_ = cmd.Start()
	//原方法
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanRunes)

	go func() {
		var line, t string
		var durationOK, fpsOK, processOK bool
		var time float64
		var fps float64
		var totalFrames uint64
		d := regexp.MustCompile("Duration: (([\\d]+):([\\d]+):([\\d]+).([\\d]+)),")
		f := regexp.MustCompile("([\\d]+) fps,")	//TODO fps可能是小数
		r := regexp.MustCompile("(frame= *\\s*(\\d+) *fps=\\s*[\\d.]+ *q= *\\s*[\\d.-]+ (L?)size=\\s*[\\d\\S]+ *time=[\\d:.]+ bitrate= *[\\d.]*\\S?bits/s speed= *\\s*[\\d.]+x)")
		//
		//end := regexp.MustCompile("(frame=\\s*\\d+ fps=\\s*[\\d.]+ q=\\s*[\\d.]+ Lsize=\\s*[\\d\\S]+ time=[\\d:.]+ bitrate=[\\d.]*\\S?bits/s speed=\\s*[\\d.]+x)")
		//Duration: 00:00:06.45
		//Stream #0:0(und): Video: h264 (High 4:4:4 Predictive) (avc1 / 0x31637661), yuv420p,
		//1920x1080 [SAR 1:1 DAR 16:, 350821 kb/s, 60 fps, 60 tbr, 15360 tbn, 120 tbc (default)
		for scanner.Scan() {
			t = scanner.Text()
			line += t
			if t == "\n" {
				//fmt.Print("获得: " + line)
				line = ""
			} else {
				if durationOK == false {
					res := d.FindStringSubmatch(line)
					if len(res) == 6 {
						durationOK = true
						//fmt.Println("Duration:", line)
						//fmt.Println("匹配时长：", res)
						line = ""
						//获取时长s
						hour,_ := strconv.Atoi(res[2])
						minute,_ := strconv.Atoi(res[3])
						second,_ := strconv.Atoi(res[4])
						frac,_ :=strconv.ParseFloat(res[5], 64)
						time = float64(hour*3600 + minute*60 + second) + frac / 100
						fmt.Println("视频时长:", time, "s")
					}
				} else if fpsOK == false {
					res := f.FindStringSubmatch(line)
					if len(res) == 2 {
						fpsOK = true
						//fmt.Println("匹配帧率：", res)
						line = ""
						//获取帧率fps和视频总帧数
						fps,_ =strconv.ParseFloat(res[1], 64)
						totalFrames = uint64(fps * time)
					}
				} else if processOK == false{
					res := r.FindStringSubmatch(line)
					if len(res) == 4 {
						//获得进度
						t, _ := strconv.ParseUint(res[2],10, 64)
						progress := float64(t * 10000 / totalFrames) / 100

						//处理进度数据，调用时指定
						method( progress )
						//count := int(t * 100 / totalFrames)
						//fmt.Printf("\r%v%v %3v%%\n", strings.Repeat("▇", count/2), strings.Repeat("░", 50-count/2), count)
						line = ""
						//
						if res[3] == "L" {
							fmt.Println("编码结束！")
						}
					}
				}
			}
		}
	}()

	_ = cmd.Wait()
}
