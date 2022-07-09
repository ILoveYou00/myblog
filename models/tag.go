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

func ExistTagByName(name string) bool {
	var tag Tag
	result := DB.Select("id").Where("name = ?", name).First(&tag)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func AddTag(tag *Tag) error {
	return DB.Create(tag).Error
}

func EditTag(p *ParamsUpdateTag) error {
	return DB.Model(&Tag{}).Where("id = ?", p.Id).Updates(Tag{Name: p.Name, ModifiedBy: p.ModifiedBy}).Error
}

func DeleteTag(id int) error {
	return DB.Where("id = ?", id).Delete(&Tag{}).Debug().Error
}
