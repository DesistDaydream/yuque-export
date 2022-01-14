package converter

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/sirupsen/logrus"
)

func ConvertHTML2Markdown(html string) (string, error) {
	converter := md.NewConverter("", true, nil)
	md, err := converter.ConvertString(html)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("笔记内容 HTML 转换到 MD 时错误")
		return "", err
	}
	return md, nil
}
