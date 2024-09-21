// subdomain-------------------------------------
// @file      : module.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2024/9/10 19:35
// -------------------------------------------

package targethandler

import (
	"fmt"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/handle"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/interfaces"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/options"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/plugins"
	"github.com/Autumn-27/ScopeSentry-Scan/internal/pool"
	"github.com/Autumn-27/ScopeSentry-Scan/pkg/logger"
	"github.com/Autumn-27/ScopeSentry-Scan/pkg/utils"
	"sync"
)

type Runner struct {
	Option     *options.TaskOptions
	NextModule interfaces.ModuleRunner
	Input      chan interface{}
	Name       string
}

func NewRunner(op *options.TaskOptions, nextModule interfaces.ModuleRunner) *Runner {
	return &Runner{
		Option:     op,
		NextModule: nextModule,
	}
}

func (r *Runner) SetInput(ch chan interface{}) {
	r.Input = ch
}

func (r *Runner) GetName() string {
	return "TargetHandler"
}

func (r *Runner) ModuleRun() error {
	handle.TaskHandle.ProgressStart(r.GetName(), r.Option.Target, r.Option.ID, len(r.Option.TargetParser))
	var allPluginWg sync.WaitGroup
	var resultWg sync.WaitGroup
	// 创建一个共享的 result 通道
	resultChan := make(chan interface{})
	go func() {
		err := r.NextModule.ModuleRun()
		if err != nil {
			logger.SlogError(fmt.Sprintf("Next module run error: %v", err))
		}
	}()
	// 结果处理 goroutine，异步读取插件的结果
	resultWg.Add(1)
	go func() {
		defer resultWg.Done()
		for {
			select {
			case result, ok := <-resultChan:
				if !ok {
					fmt.Println("关闭tartget输入")
					r.NextModule.CloseInput()
					return
				}
				// 处理每个插件的结果
				logger.SlogInfoLocal(fmt.Sprintf("%v module result: %v", r.GetName(), result))
				r.NextModule.GetInput() <- result
			}
		}
	}()

	for {
		select {
		case data, ok := <-r.Input:
			if !ok {
				// 等待所有插件运行完毕
				allPluginWg.Wait()
				close(resultChan)
				handle.TaskHandle.ProgressEnd(r.GetName(), r.Option.Target, r.Option.ID, len(r.Option.TargetParser))
				r.Option.ModuleRunWg.Done()
				// 等待结果处理完毕
				resultWg.Wait()
				return nil
			}
			allPluginWg.Add(1)
			go func(data interface{}) {
				defer allPluginWg.Done()
				// 处理输入数据
				for _, pluginName := range r.Option.TargetParser {
					var plgWg sync.WaitGroup
					logger.SlogInfoLocal(fmt.Sprintf("%v plugin start execute: %v", pluginName, data))
					plg, flag := plugins.GlobalPluginManager.GetPlugin(r.GetName(), pluginName)
					if flag {
						plgWg.Add(1)
						args, argsFlag := utils.Tools.GetParameter(r.Option.Parameters, r.GetName(), plg.GetName())
						if argsFlag {
							plg.SetParameter(args)
						} else {
							plg.SetParameter("")
						}
						plg.SetResult(resultChan)
						pluginFunc := func(data interface{}) func() {
							return func() {
								defer plgWg.Done()
								err := plg.Execute(data)
								if err != nil {
								}
							}
						}(data)
						err := pool.PoolManage.SubmitTask(r.GetName(), pluginFunc)
						if err != nil {
							plgWg.Done()
							logger.SlogError(fmt.Sprintf("task pool error: %v", err))
						}
						plgWg.Wait()
					} else {
						logger.SlogError(fmt.Sprintf("plugin %v not found", pluginName))
					}
					logger.SlogInfoLocal(fmt.Sprintf("%v plugin end execute: %v", pluginName, data))
				}
			}(data)

		}
	}
}

func (r *Runner) GetInput() chan interface{} {
	return r.Input
}

func (r *Runner) CloseInput() {
	close(r.Input)
}