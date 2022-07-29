package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/article"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ArticleSearch struct{
    article.Article
    request.PageInfo
}
