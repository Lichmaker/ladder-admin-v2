package announcementModel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/announcementModel"
    announcementModelReq "github.com/flipped-aurora/gin-vue-admin/server/model/announcementModel/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AnnouncementApi struct {
}

var announcementService = service.ServiceGroupApp.AnnouncementModelServiceGroup.AnnouncementService


// CreateAnnouncement 创建announcement表
// @Tags Announcement
// @Summary 创建announcement表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body announcementModel.Announcement true "创建announcement表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /announcement/createAnnouncement [post]
func (announcementApi *AnnouncementApi) CreateAnnouncement(c *gin.Context) {
	var announcement announcementModel.Announcement
	err := c.ShouldBindJSON(&announcement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := announcementService.CreateAnnouncement(&announcement); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAnnouncement 删除announcement表
// @Tags Announcement
// @Summary 删除announcement表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body announcementModel.Announcement true "删除announcement表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /announcement/deleteAnnouncement [delete]
func (announcementApi *AnnouncementApi) DeleteAnnouncement(c *gin.Context) {
	id := c.Query("ID")
	if err := announcementService.DeleteAnnouncement(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAnnouncementByIds 批量删除announcement表
// @Tags Announcement
// @Summary 批量删除announcement表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除announcement表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /announcement/deleteAnnouncementByIds [delete]
func (announcementApi *AnnouncementApi) DeleteAnnouncementByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := announcementService.DeleteAnnouncementByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAnnouncement 更新announcement表
// @Tags Announcement
// @Summary 更新announcement表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body announcementModel.Announcement true "更新announcement表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /announcement/updateAnnouncement [put]
func (announcementApi *AnnouncementApi) UpdateAnnouncement(c *gin.Context) {
	var announcement announcementModel.Announcement
	err := c.ShouldBindJSON(&announcement)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := announcementService.UpdateAnnouncement(announcement); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAnnouncement 用id查询announcement表
// @Tags Announcement
// @Summary 用id查询announcement表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query announcementModel.Announcement true "用id查询announcement表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /announcement/findAnnouncement [get]
func (announcementApi *AnnouncementApi) FindAnnouncement(c *gin.Context) {
	id := c.Query("ID")
	if announcement, err := announcementService.GetAnnouncement(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"announcement": announcement}, c)
	}
}

// GetAnnouncementList 分页获取announcement表列表
// @Tags Announcement
// @Summary 分页获取announcement表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query announcementModelReq.AnnouncementSearch true "分页获取announcement表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /announcement/getAnnouncementList [get]
func (announcementApi *AnnouncementApi) GetAnnouncementList(c *gin.Context) {
	var pageInfo announcementModelReq.AnnouncementSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := announcementService.GetAnnouncementInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
