package get

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// ReqBodyForGetExportURL is
type ReqBodyForExportToc struct {
	Type         string `json:"type"`
	Force        int    `json:"force"`
	Title        string `json:"title"`
	TocNodeUUID  string `json:"toc_node_uuid"`
	TocNodeURL   string `json:"toc_node_url"`
	WithChildren bool   `json:"with_children"`
}

// YuqueUserData is
type YuqueUserData struct {
	RepoID int
	Opts   YuqueUserOpts
}

// NewYuqueUserData 实例化语雀用户数据
func NewYuqueUserData(opts YuqueUserOpts) *YuqueUserData {
	url := "https://www.yuque.com/api/v2/users/" + opts.UserName + "/repos"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Auth-Token", opts.Token)

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	var (
		repoData RepoData // 知识库信息
	)

	err = json.Unmarshal(respBody, &repoData)
	if err != nil {
		return nil
	}

	var repoID int
	for _, data := range repoData.Data {
		if data.Name == opts.RepoName {
			repoID = data.ID
		}
	}

	return &YuqueUserData{
		RepoID: repoID,
		Opts:   opts,
	}
}

// GetURLForExportToc 获取待导出 TOC 的 URL
func (yud *YuqueUserData) GetURLForExportToc(tocdata TOCData) (string, error) {
	url := "https://www.yuque.com/api/books/11199981/export"
	method := "POST"

	// 根据节点信息，配置当前待导出节点的请求体信息
	reqBodyForExportToc := ReqBodyForExportToc{
		Type:         "lakebook",
		Force:        0,
		Title:        tocdata.Title,
		TocNodeUUID:  tocdata.UUID,
		TocNodeURL:   tocdata.URL,
		WithChildren: true,
	}

	// 解析请求体
	reqBodyByte, err := json.Marshal(reqBodyForExportToc)
	if err != nil {
		return "", err
	}

	// 实例化 HTTP 请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBodyByte))
	if err != nil {
		return "", err
	}

	req.Header.Add("authority", "www.yuque.com")
	req.Header.Add("accept", "application/json")
	// req.Header.Add("x-csrf-token", "3t1WleUAVuVSM5P-82xwT3Bl")
	// req.Header.Add("x-requested-with", "XMLHttpRequest")
	// req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("content-type", "application/json")
	// req.Header.Add("origin", "https://www.yuque.com")
	// req.Header.Add("sec-fetch-site", "same-origin")
	// req.Header.Add("sec-fetch-mode", "cors")
	// req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", yud.Opts.Referer)
	req.Header.Add("cookie", yud.Opts.Cookie)

	// 建立连接
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var exportDatas ExportDatas

	err = json.Unmarshal(respBody, &exportDatas)
	if err != nil {
		return "", err
	}

	if resp.StatusCode == 200 {
		logrus.WithFields(logrus.Fields{"toc": tocdata.Title, "status": resp.Status, "url": exportDatas.Data.URL}).Infof("获取待导出 TOC 的 URL 成功！")
	} else {
		return "", fmt.Errorf(resp.Status)
	}

	return exportDatas.Data.URL, nil
}

// YuqueUserOpts 通过命令行标志传递的认证选项
type YuqueUserOpts struct {
	UserName string
	RepoName string
	Cookie   string
	Referer  string
	Token    string
}

// AddFlag 用来为语雀用户数据设置一些值
func (opt *YuqueUserOpts) AddFlag() {
	pflag.StringVar(&opt.UserName, "yuque-user-name", "DesistDaydream", "用户名称")
	pflag.StringVar(&opt.RepoName, "yuque-repo-name", "学习知识库", "待导出知识库名称")
	pflag.StringVar(&opt.Token, "yuque-user-token", "", "用户 Token,在 https://www.yuque.com/settings/tokens/ 创建")
	pflag.StringVar(&opt.Cookie, "yuque-user-cookie", "", "用户 Cookie,通过浏览器的 F12 查看")
	pflag.StringVar(&opt.Referer, "yuque-referer", "https://www.yuque.com/desistdaydream/learning", "当前知识库的 URL。")
}

type ExportDatas struct {
	Data ExportData `json:"data"`
}

type ExportData struct {
	State string `json:"state"`
	URL   string `json:"url"`
}

type RepoData struct {
	Data []Data `json:"data"`
}
type User struct {
	ID             int       `json:"id"`
	Type           string    `json:"type"`
	Login          string    `json:"login"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	AvatarURL      string    `json:"avatar_url"`
	FollowersCount int       `json:"followers_count"`
	FollowingCount int       `json:"following_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Serializer     string    `json:"_serializer"`
}
type Data struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	Slug             string    `json:"slug"`
	Name             string    `json:"name"`
	UserID           int       `json:"user_id"`
	Description      string    `json:"description"`
	CreatorID        int       `json:"creator_id"`
	Public           int       `json:"public"`
	ItemsCount       int       `json:"items_count"`
	LikesCount       int       `json:"likes_count"`
	WatchesCount     int       `json:"watches_count"`
	ContentUpdatedAt time.Time `json:"content_updated_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`
	Namespace        string    `json:"namespace"`
	User             User      `json:"user"`
	Serializer       string    `json:"_serializer"`
}
