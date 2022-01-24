package handler

// 通过语雀的 API 获取到的数据都抽象为 YuqueDataHandler 接口
// 是否有必要抽象出来这层接口？
type YuqueDataHandler interface {
	// 获取 API 响应的数据
	Get(h *HandlerObject, name string) error
	// 处理 API 响应的数据
	Handle(h *HandlerObject) error
	// 参考一下 go 实现的 tcping 项目，要不要再几个方法？~
	// Result() *Result // 返回某个通用的结果
	// SetTarget(target *Target) // 设置一下要处理的 API 目标，比如设定 url 路径之类的
}
