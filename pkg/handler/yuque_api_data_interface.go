package handler

// 通过语雀的 API 获取到的数据都抽象为 YuqueData 接口
// 是否有必要抽象出来这层接口？
type YuqueData interface {
	// 获取 API 响应的数据
	Get(h *HandlerObject, name string) error
	// 处理 API 响应的数据
	Handle(h *HandlerObject) error
}
