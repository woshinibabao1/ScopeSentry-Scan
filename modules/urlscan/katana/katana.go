// katana-------------------------------------
// @file      : katana.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2024/10/11 21:33
// -------------------------------------------

package katana

import (
	"errors"
	"fmt"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/global"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/interfaces"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/types"
	"github.com/Autumn-27/ScopeSentry-Scan/pkg/logger"
	"github.com/Autumn-27/ScopeSentry-Scan/pkg/utils"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Plugin struct {
	Name           string
	Module         string
	Parameter      string
	PluginId       string
	Result         chan interface{}
	Custom         interface{}
	KatanaFileName string
	OsType         string
	KatanaDir      string
}

func NewPlugin() *Plugin {
	osType := runtime.GOOS
	var path string
	var dir string
	switch osType {
	case "windows":
		path = "katana.exe"
		dir = "win"
	case "linux":
		path = "katana"
		dir = "linux"
	default:
		dir = "darwin"
		path = "katana"
	}
	return &Plugin{
		Name:           "katana",
		Module:         "URLScan",
		PluginId:       "9669d0dcc52a5ca6dbbe580ffc99c364",
		KatanaFileName: path,
		KatanaDir:      dir,
		OsType:         osType,
	}
}
func (p *Plugin) SetCustom(cu interface{}) {
	p.Custom = cu
}

func (p *Plugin) GetCustom() interface{} {
	return p.Custom
}
func (p *Plugin) SetPluginId(id string) {
	p.PluginId = id
}

func (p *Plugin) GetPluginId() string {
	return p.PluginId
}

func (p *Plugin) SetResult(ch chan interface{}) {
	p.Result = ch
}

func (p *Plugin) SetName(name string) {
	p.Name = name
}

func (p *Plugin) GetName() string {
	return p.Name
}

func (p *Plugin) SetModule(module string) {
	p.Module = module
}

func (p *Plugin) GetModule() string {
	return p.Module
}

func (p *Plugin) Install() error {
	katanaPath := filepath.Join(global.ExtDir, "katana")
	if err := os.MkdirAll(katanaPath, os.ModePerm); err != nil {
		logger.SlogError(fmt.Sprintf("Failed to create katana folder:", err))
		return err
	}
	targetPath := filepath.Join(katanaPath, "target")
	if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
		logger.SlogError(fmt.Sprintf("Failed to create targetPath folder:", err))
		return err
	}
	resultPath := filepath.Join(katanaPath, "result")
	if err := os.MkdirAll(resultPath, os.ModePerm); err != nil {
		logger.SlogError(fmt.Sprintf("Failed to create resultPath folder:", err))
		return err
	}
	KatanaExecPath := filepath.Join(katanaPath, p.KatanaFileName)
	if _, err := os.Stat(KatanaExecPath); os.IsNotExist(err) {
		_, err := utils.Tools.HttpGetDownloadFile(fmt.Sprintf("%v/%v/%v", "https://raw.githubusercontent.com/Autumn-27/ScopeSentry-Scan/refs/heads/1.5-restructure/tools", p.KatanaDir, p.KatanaFileName), KatanaExecPath)
		if err != nil {
			_, err = utils.Tools.HttpGetDownloadFile(fmt.Sprintf("%v/%v/%v", "https://gitee.com/constL/ScopeSentry-Scan/raw/main/tools", p.KatanaDir, p.KatanaFileName), KatanaExecPath)
			if err != nil {
				return err
			}
		}
		if p.OsType == "linux" {
			err = os.Chmod(KatanaExecPath, 0755)
			if err != nil {
				logger.SlogError(fmt.Sprintf("Chmod katana Tool Fail: %s", err))
				return err
			}
		}
	}
	return nil
}

func (p *Plugin) Check() error {
	return nil
}

func (p *Plugin) SetParameter(args string) {
	p.Parameter = args
}

func (p *Plugin) GetParameter() string {
	return p.Parameter
}

func (p *Plugin) Execute(input interface{}) (interface{}, error) {
	data, ok := input.(types.AssetHttp)
	if !ok {
		logger.SlogError(fmt.Sprintf("%v error: %v input is not a string\n", p.Name, input))
		return nil, errors.New("input is not a string")
	}
	parameter := p.GetParameter()
	threads := "10"
	timeout := "5"
	maxDepth := "5"
	if parameter != "" {
		args, err := utils.Tools.ParseArgs(parameter, "t", "timeout", "depth")
		if err != nil {
		} else {
			for key, value := range args {
				switch key {
				case "t":
					threads = value
				case "timeout":
					timeout = value
				case "depth":
					maxDepth = value
				default:
					continue
				}
			}
		}
	}

	cmd := filepath.Join(filepath.Join(global.ExtDir, "katana"), p.KatanaFileName)
	resultFile := filepath.Join(filepath.Join(filepath.Join(global.ExtDir, "katana"), "result"), utils.Tools.GenerateRandomString(16))
	//defer utils.Tools.DeleteFile(resultFile)
	args := []string{
		"-u", data.URL,
		"-depth", maxDepth,
		"-fs", "rdn", "-js-crawl", "-jsonl",
		"-ef", "png,apng,bmp,gif,ico,cur,jpg,jpeg,jfif,pjp,pjpeg,svg,tif,tiff,webp,xbm,3gp,aac,flac,mpg,mpeg,mp3,mp4,m4a,m4v,m4p,oga,ogg,ogv,mov,wav,webm,eot,woff,woff2,ttf,otf,css",
		"-kf", "all", "-timeout", timeout,
		"-c", threads,
		"-p", "10",
		"-o", resultFile,
	}
	fmt.Printf("%v %v", cmd, args)
	err := utils.Tools.ExecuteCommandWithTimeout(cmd, args, 1*time.Hour)
	if err != nil {
		logger.SlogError(fmt.Sprintf("%v ExecuteCommandWithTimeout error: %v", p.GetName(), err))
	}
	return nil, nil
}

//func (p *Plugin) Execute(input interface{}) (interface{}, error) {
//	data, ok := input.(types.AssetHttp)
//	if !ok {
//		logger.SlogError(fmt.Sprintf("%v error: %v input is not a string\n", p.Name, input))
//		return nil, errors.New("input is not a string")
//	}
//
//	parameter := p.GetParameter()
//	threads := 10
//	timeout := 3
//	maxDepth := 5
//	if parameter != "" {
//		args, err := utils.Tools.ParseArgs(parameter, "t", "timeout", "depth")
//		if err != nil {
//		} else {
//			for key, value := range args {
//				switch key {
//				case "t":
//					threads, _ = strconv.Atoi(value)
//				case "timeout":
//					timeout, _ = strconv.Atoi(value)
//				case "depth":
//					maxDepth, _ = strconv.Atoi(value)
//				default:
//					continue
//				}
//			}
//		}
//	}
//	var urllist []string
//	var mu sync.Mutex
//	options := &katanaTypes.Options{
//		MaxDepth:          maxDepth,    // Maximum depth to crawl
//		FieldScope:        "rdn",       // Crawling Scope Field
//		BodyReadSize:      math.MaxInt, // Maximum response size to read
//		ScrapeJSResponses: true,
//		ExtensionFilter:   []string{"png", "apng", "bmp", "gif", "ico", "cur", "jpg", "jpeg", "jfif", "pjp", "pjpeg", "svg", "tif", "tiff", "webp", "xbm", "3gp", "aac", "flac", "mpg", "mpeg", "mp3", "mp4", "m4a", "m4v", "m4p", "oga", "ogg", "ogv", "mov", "wav", "webm", "eot", "woff", "woff2", "ttf", "otf", "css"},
//		KnownFiles:        "robotstxt,sitemapxml",
//		Timeout:           timeout,       // Timeout is the time to wait for request in seconds
//		Concurrency:       threads,       // Concurrency is the number of concurrent crawling goroutines
//		Parallelism:       10,            // Parallelism is the number of urls processing goroutines
//		Delay:             0,             // Delay is the delay between each crawl requests in seconds
//		RateLimit:         150,           // Maximum requests to send per second
//		Strategy:          "depth-first", // Visit strategy (depth-first, breadth-first)
//		OnResult: func(result output.Result) { // Callback function to execute for result
//			var r types.UrlResult
//			r.Input = data.URL
//			r.Source = result.Request.Source
//			r.Output = result.Request.URL
//			r.OutputType = result.Request.Attribute
//			r.Status = result.Response.StatusCode
//			r.Length = len(result.Response.Body)
//			r.Body = result.Response.Body
//			mu.Lock()
//			urllist = append(urllist, result.Request.URL)
//			mu.Unlock()
//			p.Result <- r
//		},
//	}
//	crawlerOptions, err := katanaTypes.NewCrawlerOptions(options)
//	if err != nil {
//		logger.SlogErrorLocal(fmt.Sprintf("katana error %v", err.Error()))
//	}
//	defer crawlerOptions.Close()
//	crawler, err := standard.New(crawlerOptions)
//	if err != nil {
//		logger.SlogErrorLocal(fmt.Sprintf("katana standard.New error %v", err.Error()))
//	}
//	defer crawler.Close()
//	err = crawler.Crawl(data.URL)
//	if err != nil {
//		logger.SlogErrorLocal(fmt.Sprintf("katana crawler.Crawl error %v: %v", input, err.Error()))
//	}
//	return urllist, nil
//}

func (p *Plugin) Clone() interfaces.Plugin {
	return &Plugin{
		Name:           p.Name,
		Module:         p.Module,
		PluginId:       p.PluginId,
		Custom:         p.Custom,
		KatanaFileName: p.KatanaFileName,
		KatanaDir:      p.KatanaDir,
		OsType:         p.OsType,
	}
}