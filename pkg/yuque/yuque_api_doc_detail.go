package yuque

import (
	"fmt"
	"time"

	"github.com/DesistDaydream/yuque-export/pkg/handler"
	"github.com/sirupsen/logrus"
)

type DocDetail struct {
	Abilities Abilities `json:"abilities"`
	Data      DocData   `json:"data"`
}
type Abilities struct {
	Update  bool `json:"update"`
	Destroy bool `json:"destroy"`
}

type Book struct {
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
	// 这个 User 与 user api 下的 User 一样
	User       User   `json:"user"`
	Serializer string `json:"_serializer"`
}
type Creator struct {
	ID               int       `json:"id"`
	Type             string    `json:"type"`
	Login            string    `json:"login"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	AvatarURL        string    `json:"avatar_url"`
	BooksCount       int       `json:"books_count"`
	PublicBooksCount int       `json:"public_books_count"`
	FollowersCount   int       `json:"followers_count"`
	FollowingCount   int       `json:"following_count"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Serializer       string    `json:"_serializer"`
}
type DocData struct {
	ID                int         `json:"id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	BookID            int         `json:"book_id"`
	Book              Book        `json:"book"`
	UserID            int         `json:"user_id"`
	Creator           Creator     `json:"creator"`
	Format            string      `json:"format"`
	Body              string      `json:"body"`
	BodyDraft         string      `json:"body_draft"`
	BodyHTML          string      `json:"body_html"`
	BodyLake          string      `json:"body_lake"`
	BodyDraftLake     string      `json:"body_draft_lake"`
	Public            int         `json:"public"`
	Status            int         `json:"status"`
	ViewStatus        int         `json:"view_status"`
	ReadStatus        int         `json:"read_status"`
	LikesCount        int         `json:"likes_count"`
	CommentsCount     int         `json:"comments_count"`
	ContentUpdatedAt  time.Time   `json:"content_updated_at"`
	DeletedAt         interface{} `json:"deleted_at"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	PublishedAt       time.Time   `json:"published_at"`
	FirstPublishedAt  time.Time   `json:"first_published_at"`
	WordCount         int         `json:"word_count"`
	Cover             string      `json:"cover"`
	Description       string      `json:"description"`
	CustomDescription interface{} `json:"custom_description"`
	Hits              int         `json:"hits"`
	Serializer        string      `json:"_serializer"`
}

func NewDocDetail() *DocDetail {
	return &DocDetail{}
}

func (dd *DocDetail) Get(h *handler.HandlerObject) error {
	url := YuqueBaseAPI + "/repos/" + fmt.Sprint(h.Namespace) + "/docs" + h.DocsSlug[0]
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debugf("检查 URL，获取%v仓库的文档列表", h.Opts.RepoName)

	err := h.HttpHandler("GET", url, dd)
	if err != nil {
		return err
	}

	return nil
}

func (dd *DocDetail) Handle(h *handler.HandlerObject) error {
	panic("not implemented") // TODO: Implement
}
