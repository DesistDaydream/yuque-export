package handler

import (
	"time"

	"github.com/sirupsen/logrus"
)

// 知识库列表
type ReposList struct {
	Data []Repo `json:"data"`
}

type Repo struct {
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

func NewReposList() *ReposList {
	return &ReposList{}
}

// 从语雀的 API 中获取知识库列表
func (r *ReposList) Get(handler *HandlerObject) error {
	url := YuqueBaseAPI + "/users/" + handler.Opts.UserName + "/repos"
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debug("检查 URL，获取知识库列表")

	err := HttpHandler("GET", url, handler.Opts.Token, r)
	if err != nil {
		return err
	}

	return nil
}
