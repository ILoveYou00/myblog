package models

type Tag struct {
	Model
	Name       string `json:"name" form:"name"`
	CreatedBy  string `json:"created_by" form:"created_by"`
	ModifiedBy string `json:"modified_by" form:"modified_by"`
	State      int    `json:"state" form:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int64) {
	DB.Model(&Tag{}).Where(maps).Count(&count)
	return
}
