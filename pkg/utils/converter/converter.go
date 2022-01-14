package converter

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/sirupsen/logrus"
)

func ConvertHTML2Markdown(html string) (string, error) {
	converter := md.NewConverter("", true, nil)
	md, err := converter.ConvertString(html)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	return md, nil
}
