package datapackagecode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/datapackagecode"
    datapackagecodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/datapackagecode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type LaDataPackageCodeApi struct {
}

var dataPackageCodeService = service.ServiceGroupApp.DatapackagecodeServiceGroup.LaDataPackageCodeService


// CreateLaDataPackageCode 创建dataPackageCode表
// @Tags LaDataPackageCode
// @Summary 创建dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datapackagecode.LaDataPackageCode true "创建dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataPackageCode/createLaDataPackageCode [post]
func (dataPackageCodeApi *LaDataPackageCodeApi) CreateLaDataPackageCode(c *gin.Context) {
	var dataPackageCode datapackagecode.LaDataPackageCode
	err := c.ShouldBindJSON(&dataPackageCode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dataPackageCodeService.CreateLaDataPackageCode(&dataPackageCode); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLaDataPackageCode 删除dataPackageCode表
// @Tags LaDataPackageCode
// @Summary 删除dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datapackagecode.LaDataPackageCode true "删除dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataPackageCode/deleteLaDataPackageCode [delete]
func (dataPackageCodeApi *LaDataPackageCodeApi) DeleteLaDataPackageCode(c *gin.Context) {
	id := c.Query("ID")
	if err := dataPackageCodeService.DeleteLaDataPackageCode(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLaDataPackageCodeByIds 批量删除dataPackageCode表
// @Tags LaDataPackageCode
// @Summary 批量删除dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dataPackageCode/deleteLaDataPackageCodeByIds [delete]
func (dataPackageCodeApi *LaDataPackageCodeApi) DeleteLaDataPackageCodeByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := dataPackageCodeService.DeleteLaDataPackageCodeByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLaDataPackageCode 更新dataPackageCode表
// @Tags LaDataPackageCode
// @Summary 更新dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datapackagecode.LaDataPackageCode true "更新dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataPackageCode/updateLaDataPackageCode [put]
func (dataPackageCodeApi *LaDataPackageCodeApi) UpdateLaDataPackageCode(c *gin.Context) {
	var dataPackageCode datapackagecode.LaDataPackageCode
	err := c.ShouldBindJSON(&dataPackageCode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dataPackageCodeService.UpdateLaDataPackageCode(dataPackageCode); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLaDataPackageCode 用id查询dataPackageCode表
// @Tags LaDataPackageCode
// @Summary 用id查询dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query datapackagecode.LaDataPackageCode true "用id查询dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataPackageCode/findLaDataPackageCode [get]
func (dataPackageCodeApi *LaDataPackageCodeApi) FindLaDataPackageCode(c *gin.Context) {
	id := c.Query("ID")
	if redataPackageCode, err := dataPackageCodeService.GetLaDataPackageCode(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redataPackageCode": redataPackageCode}, c)
	}
}

// GetLaDataPackageCodeList 分页获取dataPackageCode表列表
// @Tags LaDataPackageCode
// @Summary 分页获取dataPackageCode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query datapackagecodeReq.LaDataPackageCodeSearch true "分页获取dataPackageCode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataPackageCode/getLaDataPackageCodeList [get]
func (dataPackageCodeApi *LaDataPackageCodeApi) GetLaDataPackageCodeList(c *gin.Context) {
	var pageInfo datapackagecodeReq.LaDataPackageCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dataPackageCodeService.GetLaDataPackageCodeInfoList(pageInfo); err != nil {
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
