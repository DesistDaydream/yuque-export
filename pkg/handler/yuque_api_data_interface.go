package handler

// 通过语雀的 API 获取到的数据都抽象为 YuqueData 接口
// 是否有必要抽象出来这层接口？
type YuqueData interface {
	Get(h *HandlerObject) error
	Handle(h *HandlerObject) error
}
