import service from '@/utils/request'
import { unique } from '@/utils/arrayFun'
// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/login [post]
export const login = (data) => {
  return service({
    url: '/base/login',
    method: 'post',
    data: data
  })
}

// @Summary 获取验证码
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/captcha [post]
export const captcha = (data) => {
  return service({
    url: '/base/captcha',
    method: 'post',
    data: data
  })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /base/resige [post]
export const register = (data) => {
  return service({
    url: '/user/admin_register',
    method: 'post',
    data: data
  })
}

// @Summary 修改密码
// @Produce  application/json
// @Param data body {username:"string",password:"string",newPassword:"string"}
// @Router /user/changePassword [post]
export const changePassword = (data) => {
  return service({
    url: '/user/changePassword',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
export const getUserList = (data) => {
  return service({
    url: '/user/getUserList',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SetUserAuth true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
export const setUserAuthority = (data) => {
  return service({
    url: '/user/setUserAuthority',
    method: 'post',
    data: data
  })
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
export const deleteUser = (data) => {
  return service({
    url: '/user/deleteUser',
    method: 'delete',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserInfo [put]
export const setUserInfo = (data) => {
  return service({
    url: '/user/setUserInfo',
    method: 'put',
    data: data
  })
}

// @Tags SysUser
// @Summary 设置用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysUser true "设置用户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setSelfInfo [put]
export const setSelfInfo = (data) => {
  return service({
    url: '/user/setSelfInfo',
    method: 'put',
    data: data
  })
}

// @Tags User
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.setUserAuthorities true "设置用户权限"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthorities [post]
export const setUserAuthorities = (data) => {
  return service({
    url: '/user/setUserAuthorities',
    method: 'post',
    data: data
  })
}

// @Tags User
// @Summary 获取用户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserInfo [get]
export const getUserInfo = () => {
  return service({
    url: '/user/getUserInfo',
    method: 'get'
  })
}

export const resetPassword = (data) => {
  return service({
    url: '/user/resetPassword',
    method: 'post',
    data: data
  })
}

export const batchGetUserInfoByUUID = async(uuidArray) => {
  const res = new Map() // 存储结果

  // 数组去重
  uuidArray = unique(uuidArray)

  if (uuidArray.length === 0) {
    return res
  }

  // 调用接口查询数据，返回对象
  await service({
    url: '/user/batchQuery',
    method: 'post',
    data: {
      'uuid': uuidArray
    },
    donNotShowLoading: true,
  }).then((apiRes) => {
    if (apiRes.code !== 0) {
      console.log('批量查询用户数据失败！', apiRes)
      throw new Error(apiRes.msg)
    }

    if (apiRes.data) {
      for (var i = 0, len = apiRes.data.length; i < len; i++) {
        res.set(apiRes.data[i].uuid, apiRes.data[i])
      }
    }
  })

  return res
}

export const batchGetUserInfoByEmail = async(emailArray) => {
  const res = new Map() // 存储结果

  // 数组去重
  emailArray = unique(emailArray)

  if (emailArray.length === 0) {
    return res
  }

  // 调用接口查询数据，返回对象
  await service({
    url: '/user/batchQuery',
    method: 'post',
    data: {
      'email': emailArray
    },
    donNotShowLoading: true,
  }).then((apiRes) => {
    if (apiRes.code !== 0) {
      console.log('批量查询用户数据失败！', apiRes)
      throw new Error(apiRes.msg)
    }

    if (apiRes.data) {
      for (var i = 0, len = apiRes.data.length; i < len; i++) {
        res.set(apiRes.data[i].email, apiRes.data[i])
      }
    }
  })

  return res
}

export const registerWithInviteCode = (data) => {
//   console.log(data.value)
  return service({
    url: '/base/registerWithInviteCode',
    method: 'post',
    data: data
  })
}

