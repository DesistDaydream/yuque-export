package yuque

import (
	"fmt"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

type DocsList struct {
	Data []Data `json:"data"`
}
type LastEditor struct {
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
	ID                int         `json:"id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	Description       string      `json:"description"`
	UserID            int         `json:"user_id"`
	BookID            int         `json:"book_id"`
	Format            string      `json:"format"`
	Public            int         `json:"public"`
	Status            int         `json:"status"`
	ViewStatus        int         `json:"view_status"`
	ReadStatus        int         `json:"read_status"`
	LikesCount        int         `json:"likes_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	DraftVersion      int         `json:"draft_version"`
	LastEditorID      int         `json:"last_editor_id"`
	WordCount         int         `json:"word_count"`
	Cover             string      `json:"cover"`
	CustomDescription interface{} `json:"custom_description"`
	LastEditor        LastEditor  `json:"last_editor"`
	Book              interface{} `json:"book"`
	Serializer        string      `json:"_serializer"`
}

func NewDocsList() *DocsList {
	return &DocsList{}
}

func (d *DocsList) Get(h *handler.HandlerObject) error {
	url := YuqueBaseAPI + "/repos/" + fmt.Sprint(h.Namespace) + "/docs"
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debugf("检查 URL，获取%v仓库的文档列表", h.Opts.RepoName)

	err := h.HttpHandler("GET", url, d)
	if err != nil {
		return err
	}

	return nil
}
func (d *DocsList) Handle(h *handler.HandlerObject) error {
	panic("not implemented") // TODO: Implement
}
