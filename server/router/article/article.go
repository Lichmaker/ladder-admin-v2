package article

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
}

// InitArticleRouter 初始化 Article 路由信息
func (s *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	articleRouterWithoutRecord := Router.Group("article")
	var articleApi = v1.ApiGroupApp.ArticleApiGroup.ArticleApi
	{
		articleRouter.POST("createArticle", articleApi.CreateArticle)   // 新建Article
		articleRouter.DELETE("deleteArticle", articleApi.DeleteArticle) // 删除Article
		articleRouter.DELETE("deleteArticleByIds", articleApi.DeleteArticleByIds) // 批量删除Article
		articleRouter.PUT("updateArticle", articleApi.UpdateArticle)    // 更新Article
	}
	{
		articleRouterWithoutRecord.GET("findArticle", articleApi.FindArticle)        // 根据ID获取Article
		articleRouterWithoutRecord.GET("getArticleList", articleApi.GetArticleList)  // 获取Article列表
	}
}
