// pool-------------------------------------
// @file      : interface.go
// @author    : Autumn
// @contact   : rainy-autumn@outlook.com
// @time      : 2024/9/10 19:26
// -------------------------------------------

package plugins

type Plugin interface {
	GetName() string
	SetName(name string)
	GetModule() string
	SetModule(name string)
	SetResult(ch chan interface{})
	SetParameter(args string)
	GetParameter() string
	Execute(input interface{}) error
	Install() error
	Check() error
}
