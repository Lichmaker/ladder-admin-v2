package datapackagecode

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datapackagecode"
	datapackagecodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/datapackagecode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userDataUsageModel"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/captcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DataPackageCodeApi struct {
}

var dataPackageCodeService = service.ServiceGroupApp.DatapackagecodeServiceGroup.DataPackageCodeService
var userDataUsageService = service.ServiceGroupApp.UserDataUsageModelServiceGroup.UserDataUsageService
var store = captcha.NewDefaultRedisStore()
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

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

	// 如果没有传入唯一编号，则自动生成
	if len(dataPackageCode.UniqueCode) == 0 {
		dataPackageCode.UniqueCode = strings.ToUpper(utils.RandomString(15))
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
	if list, total, err := dataPackageCodeService.GetDataPackageCodeInfoList(pageInfo, "id ASC"); err != nil {
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

func (dataPackageCodeApi *DataPackageCodeApi) BatchCreateDataPackageCode(c *gin.Context) {
	var req datapackagecodeReq.BatchCreateDataPackageCode
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.Total <= 0 || req.StandardData <= 0 || req.Days <= 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	standardData := int(req.StandardData)
	days := int(req.Days)

	batchData := make([]datapackagecode.DataPackageCode, 0)
	for i := 0; i < int(req.Total); i++ {
		tmp := datapackagecode.DataPackageCode{}
		tmp.UniqueCode = req.CodePrefix + strings.ToUpper(utils.RandomString(15))
		tmp.StandardData = &standardData
		tmp.Days = &days
		tmp.Remark = req.Remark
		tmp.Enable = true
		batchData = append(batchData, tmp)
	}

	dbTransaction := global.GVA_DB.Begin()

	if err := dataPackageCodeService.BatchCreateDataPackageCode(dbTransaction, batchData); err != nil {
		dbTransaction.Rollback()
		global.GVA_LOG.Error("批量创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		if err := dbTransaction.Commit().Error; err != nil {
			dbTransaction.Rollback()
			global.GVA_LOG.Error("批量创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	}
}

func (dataPackageCodeApi *DataPackageCodeApi) ExchangeDataPackageCode(c *gin.Context) {
	var req datapackagecodeReq.ExchangeDataPackageCode
	now := time.Now()
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	if cast.ToInt(v) >= 20 || !store.Verify(req.CaptchaId, req.Captcha, true) {
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("验证码错误", c)
		return
	}

	query, err := dataPackageCodeService.GetByCode(req.Code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("不存在的激活码", c)
			return
		} else {
			response.FailWithMessage("数据库异常，请稍后再试", c)
			return
		}
	}
	if !query.Enable {
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("该激活码已被禁用", c)
		return
	}
	if len(query.UsedBy) > 0 {
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("该激活码已被使用", c)
		return
	}

	currentUser, err := userService.GetUserInfo(utils.GetUserUuid(c))
	if err != nil {
		response.FailWithMessage("数据库异常，请稍后再试", c)
		return
	}

	// redis锁定激活码
	lockKey := fmt.Sprintf("data_package_lock_%s", req.Code)
	lockTTL := time.Second * 20
	if !global.GVA_REDIS.SetNX(context.Background(), lockKey, 1, lockTTL).Val() {
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("服务器过热，请稍后再试", c)
		return
	}

	query.UsedAt = &now
	query.UsedBy = currentUser.UUID.String()

	if err := dataPackageCodeService.UpdateDataPackageCode(query); err != nil {
		response.FailWithMessage("数据库异常，请稍后再试", c)
		return
	}

	response.OkWithMessage("兑换成功", c)
}

func (dataPackageCodeApi *DataPackageCodeApi) GetMyDataPackageCodeList(c *gin.Context) {
	var pageInfo datapackagecodeReq.DataPackageCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UsedBy = utils.GetUserUuid(c).String()
	if list, total, err := dataPackageCodeService.GetDataPackageCodeInfoList(pageInfo, "consumed_at ASC, id DESC"); err != nil {
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

func (dataPackageCodeApi *DataPackageCodeApi) ConsumePackage(c *gin.Context) {
	var req datapackagecodeReq.ConsumePackageReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	now := carbon.Now("Asia/Shanghai")
	nowSTD := now.ToStdTime()

	query, err := dataPackageCodeService.GetDataPackageCode(cast.ToString(req.PackageID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if query.ConsumedAt != nil {
		response.FailWithMessage("已使用，无法重复操作", c)
		return
	}

	consumePackLock := fmt.Sprintf("consume_lock_%d", req.PackageID)
	if !global.GVA_REDIS.SetNX(context.Background(), consumePackLock, 1, time.Second*3).Val() {
		response.FailWithMessage("服务器冒烟了，请稍后再操作", c)
		return
	}

	dbTransaction := global.GVA_DB.Begin()

	query.ConsumedAt = &nowSTD
	if err := dataPackageCodeService.UpdateDataPackageCodeWithDB(dbTransaction, query); err != nil {
		dbTransaction.Rollback()
		response.FailWithMessage(err.Error(), c)
		return
	}

	endTime := now.AddDays(*query.Days).EndOfDay().ToStdTime()

	var userDataUsage userDataUsageModel.UserDataUsage
	userDataUsage.UserUuid = utils.GetUserUuid(c).String()
	userDataUsage.StartDate = &nowSTD
	userDataUsage.EndDate = &endTime
	userDataUsage.StandardData = *query.StandardData
	if err := userDataUsageService.CreateUserDataUsageWithDB(dbTransaction, &userDataUsage); err != nil {
		dbTransaction.Rollback()
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	if err := dbTransaction.Commit().Error; err != nil {
		dbTransaction.Rollback()
		global.GVA_LOG.Error("提交失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}
