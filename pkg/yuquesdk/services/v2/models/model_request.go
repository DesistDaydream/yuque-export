package models

//CreateGroup struct for create group
type CreateGroup struct {
	Name        string `request:"name"`
	Login       string `request:"login"`
	Description string `request:"description"`
}

//CreateRepo struct for create repo
type CreateRepo struct {
	Name        string `request:"name"`
	Slug        string `request:"slug"`
	Description string `request:"description"`
	// 0 私密, 1 内网公开, 2 全网公开
	Public int `request:"public"`
	// ‘Book’ 文库, ‘Design’ 画板, 请注意大小写
	Type string `request:"type"`
}

//UpdateRepo struct for update repo
type UpdateRepo struct {
	Name        string `request:"name"`
	Slug        string `request:"slug"`
	Description string `request:"description"`
	// 0 私密, 1 内网公开, 2 全网公开
	Public int    `request:"public"`
	Toc    string `request:"toc"`
}

//GroupAddUser struct for update repo
type GroupAddUser struct {
	Role int `request:"role"`
}

//DocGet struct for get doc
type DocGet struct {
	Raw int `request:"raw"`
}

//DocCreate struct for create doc
type DocCreate struct {
	Title  string `request:"title"`
	Slug   string `request:"slug"`
	Public int    `request:"public"`
	// markdown or lake, default is markdown
	Format string `request:"format"`
	//max 5Mb
	Body string `request:"body"`
}
