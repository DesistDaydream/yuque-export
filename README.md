# 概述
> 参考：
> - https://github.com/wujiyu115/yuqueg (后续优化代码时受该项目启发)

由于语雀的导出限制，当知识库内文章过多时，将会导出失败，所以写个小工具，逐一导出

> 注意：该仓库并不是为了迁移语雀，而是为了备份，所以这里并不是逐一导出知识库下面的每一篇文档，只是根据目录层级深度，导出该目录层及其子目录下的所有文档，这些文档统一封装为语雀的 `.lakebook` 类型文件

语雀将 笔记、表格、思维图 等等，语雀将所有内容抽象为 TOC(节点)

# 运行
- token 从[此处](https://www.yuque.com/settings/tokens)获取
- cookie 从浏览器的 F12 中获取
- referer 为当前知识库的 URL

## 导出文档集合
```bash
go run cmd/main.go --log-level=debug --user-token="XXX" --user-cookie="YYY" --referer="ZZZ" --repo-name="学习知识库" --export=true --time-out=120s
```

## 导出每篇文档
```bash
go run cmd/main.go --log-level=debug --user-token="XXX" --method=all --repo-name="学习知识库" --export=true --export-duration=1
```

## 获取文档详情
```bash
go run cmd/main.go --log-level=debug --user-token="XXX" --method=get --repo-name="学习知识库" --export=true --export-duration=0 --concurrency=1
```

# TODO
- 请求速度过快会被限流~导致导出部分文档失败~
- 导出后将文档发送到云盘