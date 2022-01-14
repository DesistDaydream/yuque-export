package yuque

import "time"

// 用户信息
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

// TOCs 列表。
type TocsList struct {
	Data []TOC `json:"data"`
}

// 语雀的 Toc(节点) 就是指一篇 笔记、表格、思维图 等等，语雀将所有内容抽象为 TOC(节点)
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
	// 节点深度
	Depth int `json:"depth"`
	// 节点 URL 的 slug
	Slug string `json:"slug"`
}

// 文档列表
type DocsList struct {
	Data []Doc `json:"data"`
}

type Doc struct {
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
	LastEditor        User        `json:"last_editor"`
	Book              interface{} `json:"book"`
	Serializer        string      `json:"_serializer"`
}

// 文档详情数据
type DocDetailData struct {
	Abilities Abilities `json:"abilities"`
	Data      DocDetail `json:"data"`
}

type Abilities struct {
	Update  bool `json:"update"`
	Destroy bool `json:"destroy"`
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

type DocDetail struct {
	ID                int         `json:"id"`
	Slug              string      `json:"slug"`
	Title             string      `json:"title"`
	BookID            int         `json:"book_id"`
	Book              Repo        `json:"book"`
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

// 其他
type ExportsData struct {
	Data Export `json:"data"`
}

type Export struct {
	State string `json:"state"`
	URL   string `json:"url"`
}
