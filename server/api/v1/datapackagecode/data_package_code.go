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

type DataPackageCodeApi struct {
}

var dataPackageCodeService = service.ServiceGroupApp.DatapackagecodeServiceGroup.DataPackageCodeService


// CreateDataPackageCode 创建dataPackageCode表
// @Tags DataPackageCode
// @Summary 创建dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datapackagecode.DataPackageCode true "创建dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataPackageCode/createDataPackageCode [post]
func (dataPackageCodeApi *DataPackageCodeApi) CreateDataPackageCode(c *gin.Context) {
	var dataPackageCode datapackagecode.DataPackageCode
	err := c.ShouldBindJSON(&dataPackageCode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dataPackageCodeService.CreateDataPackageCode(&dataPackageCode); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDataPackageCode 删除dataPackageCode表
// @Tags DataPackageCode
// @Summary 删除dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datapackagecode.DataPackageCode true "删除dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataPackageCode/deleteDataPackageCode [delete]
func (dataPackageCodeApi *DataPackageCodeApi) DeleteDataPackageCode(c *gin.Context) {
	id := c.Query("ID")
	if err := dataPackageCodeService.DeleteDataPackageCode(id); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataPackageCodeByIds 批量删除dataPackageCode表
// @Tags DataPackageCode
// @Summary 批量删除dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dataPackageCode/deleteDataPackageCodeByIds [delete]
func (dataPackageCodeApi *DataPackageCodeApi) DeleteDataPackageCodeByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := dataPackageCodeService.DeleteDataPackageCodeByIds(ids); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDataPackageCode 更新dataPackageCode表
// @Tags DataPackageCode
// @Summary 更新dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datapackagecode.DataPackageCode true "更新dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataPackageCode/updateDataPackageCode [put]
func (dataPackageCodeApi *DataPackageCodeApi) UpdateDataPackageCode(c *gin.Context) {
	var dataPackageCode datapackagecode.DataPackageCode
	err := c.ShouldBindJSON(&dataPackageCode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dataPackageCodeService.UpdateDataPackageCode(dataPackageCode); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDataPackageCode 用id查询dataPackageCode表
// @Tags DataPackageCode
// @Summary 用id查询dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query datapackagecode.DataPackageCode true "用id查询dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataPackageCode/findDataPackageCode [get]
func (dataPackageCodeApi *DataPackageCodeApi) FindDataPackageCode(c *gin.Context) {
	id := c.Query("ID")
	if redataPackageCode, err := dataPackageCodeService.GetDataPackageCode(id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redataPackageCode": redataPackageCode}, c)
	}
}

// GetDataPackageCodeList 分页获取dataPackageCode表列表
// @Tags DataPackageCode
// @Summary 分页获取dataPackageCode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query datapackagecodeReq.DataPackageCodeSearch true "分页获取dataPackageCode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataPackageCode/getDataPackageCodeList [get]
func (dataPackageCodeApi *DataPackageCodeApi) GetDataPackageCodeList(c *gin.Context) {
	var pageInfo datapackagecodeReq.DataPackageCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dataPackageCodeService.GetDataPackageCodeInfoList(pageInfo); err != nil {
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
