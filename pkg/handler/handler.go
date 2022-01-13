package handler

// 用来处理语雀API的数据
type HandlerObject struct {
	// 待导出的知识库。可以是仓库的ID，也可以是以斜线分割的用户名和仓库slug的组合
	Namespace int
	Opts      YuqueUserOpts
}

// 根据命令行标志实例化一个处理器
func NewHandlerObject(opts YuqueUserOpts) *HandlerObject {
	return &HandlerObject{
		Opts: opts,
	}
}
