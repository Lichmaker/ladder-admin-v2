// 自动生成模板Article
package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Article 结构体
type Article struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:;"`
}


// TableName Article 表名
func (Article) TableName() string {
  return "article"
}

