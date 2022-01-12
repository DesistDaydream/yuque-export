package handler

import "time"

var (
	YuqueBaseAPI  = "https://www.yuque.com/api/v2"
	YuqueReposAPI = "/repos"
	YuqueUserAPI  = "/user/"
)

type ExportsData struct {
	Data Export `json:"data"`
}

type Export struct {
	State string `json:"state"`
	URL   string `json:"url"`
}

// 用户数据
type UserData struct {
	Data User `json:"data"`
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

func (u *UserData) Get() {

}

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

func (r *ReposList) Get() {

}

// TOCs 节点列表。
type TocsList struct {
	Data []TOC `json:"data"`
}

// TOCData 语雀的节点就是指一篇 笔记、表格、思维图 等等，语雀将所有内容抽象为 TOC(节点)
type TOC struct {
	// 节点类型
	Type string `json:"type"`
	// 节点名称
	Title string `json:"title"`
	// 节点唯一标识符
	UUID string `json:"uuid"`
	// 该节点的链接或 slug
	URL string `json:"url"`
	// 上一个节点的 UUID
	PrevUUID string `json:"prev_uuid"`
	// 下一个节点的 UUID
	SiblingUUID string `json:"sibling_uuid"`
	// 第一个子节点的 UUID
	ChildUUID string `json:"child_uuid"`
	// 父节点的 UUID
	ParentUUID string `json:"parent_uuid"`
	// 文档类型节点的标识符
	DocID int `json:"doc_id"`
	// 节点层级
	Level int `json:"level"`
	// 节点 ID
	ID int `json:"id"`
	// 连接是否在新窗口打开，0 当前页面打开、1 在新窗口打开
	OpenWindow int `json:"open_window"`
	// 节点是否可见。0 不可见、1 可见
	Visible int `json:"visible"`
	Depth   int `json:"depth"`
	// 节点 URL 的 slug
	Slug string `json:"slug"`
}

func (t *TocsList) Get() {

}
