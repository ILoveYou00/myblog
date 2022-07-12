package models

// ParamsCreateTag 添加标签
type ParamsCreateTag struct {
	Name      string `json:"name" form:"name" binding:"required,max=100"`
	State     int    `json:"State" form:"State" binding:"oneof=1 0"`
	CreatedBy string `json:"created_by" form:"created_by" binding:"required,max=100"`
}

//ParamsUpdateTag 修改标签
type ParamsUpdateTag struct {
	Id         int    `json:"id" form:"id"`
	Name       string `json:"name" form:"name" binding:"required,max=100"`
	ModifiedBy string `json:"modified_by" form:"modified_by" binding:"required,max=100" `
}

type ParamsGetArticle struct {
	Article []Article `json:"article"`
	Total   int64     `json:"total"`
}

//ParamsCreateArticle 新增文章
type ParamsCreateArticle struct {
	TagId     int    `json:"tag_id" form:"tag_id"`
	Title     string `json:"title" form:"title" binding:"required"`
	Desc      string `json:"desc" form:"desc" binding:"required,max=255"`
	Content   string `json:"content" form:"content" binding:"required,max=65535"`
	CreatedBy string `json:"created_by" form:"created_by" binding:"required"`
	State     int    `json:"state" form:"state" binding:"oneof=1 0"`
}

//ParamsUpdateArticle 修改文章
type ParamsUpdateArticle struct {
	Id         int    `json:"id" form:"id"`
	TagId      int    `json:"tag_id" form:"tag_id"`
	Title      string `json:"title" form:"title" binding:"required"`
	Desc       string `json:"desc" form:"desc" binding:"required,max=255"`
	Content    string `json:"content" form:"content" binding:"required,max=65535"`
	ModifiedBy string `json:"modified_by" form:"modified_by" binding:"required"`
	State      int    `json:"state" form:"state" binding:"oneof=1 0"`
}

type ParamsAuth struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
