package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/article"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    articleReq "github.com/flipped-aurora/gin-vue-admin/server/model/article/request"
)

type ArticleService struct {
}

// CreateArticle 创建Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService) CreateArticle(article article.Article) (err error) {
	err = global.GVA_DB.Create(&article).Error
	return err
}

// DeleteArticle 删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)DeleteArticle(article article.Article) (err error) {
	err = global.GVA_DB.Delete(&article).Error
	return err
}

// DeleteArticleByIds 批量删除Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)DeleteArticleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]article.Article{},"id in ?",ids.Ids).Error
	return err
}

// UpdateArticle 更新Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)UpdateArticle(article article.Article) (err error) {
	err = global.GVA_DB.Save(&article).Error
	return err
}

// GetArticle 根据id获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)GetArticle(id uint) (article article.Article, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&article).Error
	return
}

// GetArticleInfoList 分页获取Article记录
// Author [piexlmax](https://github.com/piexlmax)
func (articleService *ArticleService)GetArticleInfoList(info articleReq.ArticleSearch) (list []article.Article, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&article.Article{})
    var articles []article.Article
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&articles).Error
	return  articles, total, err
}
