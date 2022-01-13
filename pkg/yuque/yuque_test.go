package yuque_test

import (
	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/DesistDaydream/yuque-export/pkg/yuque"
)

// 怎么写测试用例？~~~o(╯□╰)o
func DocsTest() {
	var h *handler.HandlerObject
	docsList := yuque.NewDocsList()
	err := docsList.Get(h)
	if err != nil {
		panic(err)
	}
}
