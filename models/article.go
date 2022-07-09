package models

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	DB.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func GetArticle(id int) (article Article, err error) {
	err = DB.Where("id = ?", id).Preload("Tag").First(&article).Debug().Error
	return
}

//GetArticleTotal 查询文章数
func GetArticleTotal() (count int64, err error) {
	err = DB.Model(Article{}).Count(&count).Error
	return
}

func GetArticles(pageNum int, pageSize int) (articles []Article, err error) {
	DB.Preload("Tag").Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

//CreateArticle 添加新文章
func CreateArticle(article *Article) (err error) {
	return DB.Create(article).Debug().Error
}
