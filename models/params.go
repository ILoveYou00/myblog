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
	ModifiedBy string `json:"modified_by" form:"modified_by" binding:"required,max=100"`
}
