# 概述
由于语雀的导出限制，当知识库内文章过多时，将会导出失败，所以写个小工具，逐一导出

语雀将 笔记、表格、思维图 等等，语雀将所有内容抽象为 TOC(节点)

# 运行
- tongke 从[此处](https://www.yuque.com/settings/tokens)获取
- cookie 从浏览器的 F12 中获取
- referer 为当前知识库的 URL

```bash
go run *.go --log-level=debug --yuque-user-token="XXX" --yuque-user-cookie="YYY" --yuque-referer="ZZZ" --export=true
```

# TODO
- 请求速度过快会被限流~导致导出部分文档失败~
- 导出后将文档发送到云盘 