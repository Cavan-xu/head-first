package singleton

/*
	单例模式
*/

import (
	"sync"
)

var (
	_controller *Controller
	once        *sync.Once
)

type Controller struct{}

func init() {
	_controller = &Controller{}
}

func GetInstance() *Controller {
	if _controller == nil {
		once.Do(func() {
			_controller = &Controller{}
		})
	}
	return _controller
}
