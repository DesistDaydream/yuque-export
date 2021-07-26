# 概述
由于语雀的导出限制，当知识库内文章过多时，将会导出失败，所以写个小工具，逐一导出

语雀将 笔记、表格、思维图 等等，语雀将所有内容抽象为 TOC(节点)

# 运行
从[此处](https://www.yuque.com/settings/tokens)获取认证所需 Token

```bash
go run *.go --log-level=debug --token="XXXXXXX"
```