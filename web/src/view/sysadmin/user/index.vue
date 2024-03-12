<template>
  <div>
    <!-- <warning-bar title="注：右上角头像下拉可切换角色" /> -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addUser"
        >新增用户</el-button>
      </div>
      <el-table
        :data="tableData"
        row-key="ID"
      >
        <!-- <el-table-column
          align="left"
          label="头像"
          min-width="75"
        >
          <template #default="scope">
            <CustomPic
              style="margin-top:8px"
              :pic-src="scope.row.headerImg"
            />
          </template>
        </el-table-column> -->
        <el-table-column
          align="left"
          label="ID"
          min-width="50"
          prop="ID"
        />
        <el-table-column
          align="left"
          label="用户名"
          min-width="200"
          prop="userName"
        />
        <!-- <el-table-column
          align="left"
          label="昵称"
          min-width="200"
          prop="nickName"
        /> -->
        <el-table-column
          align="left"
          label="邮箱"
          min-width="250"
          prop="email"
        >
          <template #default="{ row }">
            <div>
              <el-button
                :icon="CopyDocument"
                type="primary"
                circle
                size="small"
                @click="copyToClipboard(row.email)"
              />
              {{ row.email }}
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="激活状态"
          min-width="100"
        >
          <template #default="scope">
            <div :class="[ scope.row.userExt.enable == 1 ?'green-light':'red-light']" /> {{ scope.row.userExt.enable == 1 ?'激活':'未激活' }}
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="UUID"
          min-width="350"
          prop="uuid"
        >
          <template #default="{ row }">
            <div>
              <el-button
                :icon="CopyDocument"
                type="primary"
                circle
                size="small"
                @click="copyToClipboard(row.uuid)"
              />
              {{ row.uuid }}
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="用量"
        >
          <template #default="scope">
            {{ scope.row.userExt.standardData }} MB
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="有效期"
          min-width="190"
        >
          <template #default="scope">
            {{ formatColumnValidDate(scope.row.userExt) }}
          </template>
        </el-table-column>

        <el-table-column
          label="查询数据"
          min-width="150"
          fixed="right"
        ><template #default="scope">
          <el-button
            type="primary"
            link
            icon="Histogram"
            @click="toTargetWithUUID('usage', scope.row.uuid)"
          >用量</el-button>
          <el-button
            type="primary"
            link
            icon="Histogram"
            @click="toTargetWithEmail('data-stat', scope.row.email)"
          >历史</el-button>
        </template></el-table-column>
        <el-table-column
          label="操作"
          min-width="250"
          fixed="right"
        >
          <template #default="scope">

            <el-button
              type="primary"
              link
              icon="edit"
              @click="openEdit(scope.row)"
            >编辑</el-button>
            <el-button
              type="primary"
              link
              icon="magic-stick"
              @click="resetPasswordFunc(scope.row)"
            >重置密码</el-button>
            <el-popover
              v-model:visible="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>确定要删除此用户吗</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button
                  type="primary"
                  link
                  @click="scope.row.visible = false"
                >取消</el-button>
                <el-button
                  type="primary"
                  @click="deleteUserFunc(scope.row)"
                >确定</el-button>
              </div>
              <template #reference>
                <el-button
                  type="primary"
                  link
                  icon="delete"
                >删除</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>

      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="addUserDialog"
      title="用户"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:50vh;overflow:auto;padding:0 1px;">
        <el-form
          ref="userForm"
          :rules="rules"
          :model="userInfo"
          label-width="110px"
        >
          <el-form-item
            label="用户名"
            prop="userName"
          >
            <el-input
              v-model="userInfo.userName"
              :disabled="userInfoUpdateDisable"
            />
          </el-form-item>
          <el-form-item
            v-if="dialogFlag === 'add'"
            label="密码"
            prop="password"
          >
            <el-input v-model="userInfo.password" />
          </el-form-item>
          <el-form-item
            label="昵称"
            prop="nickName"
          >
            <el-input v-model="userInfo.nickName" />
          </el-form-item>
          <el-form-item
            label="手机号"
            prop="phone"
          >
            <el-input v-model="userInfo.phone" />
          </el-form-item>
          <el-form-item
            label="邮箱"
            prop="email"
          >
            <el-input v-model="userInfo.email" />
          </el-form-item>
          <el-form-item
            label="基础用量(MB)"
            prop="standardDataMB"
          >
            <el-input-number v-model="userInfo.standardDataMB" />
          </el-form-item>

          <el-form-item
            label="账号有效期"
            prop="validDateTime"
          >
            <el-date-picker
              v-model="userInfo.validDateTime"
              type="daterange"
              unlink-panels
              range-separator="To"
              start-placeholder="Start date"
              end-placeholder="End date"
              :shortcuts="shortcuts"
              size="small"
              @change="validDateTimeChanged"
            />
          </el-form-item>

          <el-form-item
            label="用户角色"
            prop="authorityId"
          >
            <el-cascader
              v-model="userInfo.authorityIds"
              style="width:100%"
              :options="authOptions"
              :show-all-levels="false"
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
              :disabled="userInfoUpdateDisable"
            />
          </el-form-item>
          <el-form-item
            v-if="userInfoUpdateDisable"
            label="启用"
            prop="enable"
          >
            <el-switch
              v-model="userInfo.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="0"
            />
          </el-form-item>
          <!-- <el-form-item
            label="头像"
          >
            <div
              style="display:inline-block"
              @click="openHeaderChange"
            >
              <img
                v-if="userInfo.headerImg"
                alt="头像"
                class="header-img-box"
                :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg"
              >
              <div
                v-else
                class="header-img-box"
              >从媒体库选择</div>
            </div>
          </el-form-item> -->

        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddUserDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterAddUserDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg
      ref="chooseImg"
      :target="userInfo"
      :target-key="`headerImg`"
    />
  </div>
</template>

<script setup>

import {
//   getUserList,
//   setUserAuthorities,
  //   register,
  deleteUser
} from '@/api/user'

import {
  getUserListV2,
  createUserV2,
  updateUserV2,
} from '@/api/userExt'

import { getAuthorityList } from '@/api/authority'
// import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
// import WarningBar from '@/components/warningBar/warningBar.vue'
import {
  // setUserInfo,
  resetPassword
} from '@/api/user.js'
import Clipboard from 'clipboard'

import { nextTick, ref, watch, inject } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useRouter } from 'vue-router'

import {
  CopyDocument,
} from '@element-plus/icons-vue'

defineOptions({
  name: 'User',
})
const dayjs = inject('$dayjs')
// const path = ref(import.meta.env.VITE_BASE_API + '/')
// 初始化相关
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
          AuthorityData.forEach(item => {
            if (item.children && item.children.length) {
              const option = {
                authorityId: item.authorityId,
                authorityName: item.authorityName,
                children: []
              }
              setAuthorityOptions(item.children, option.children)
              optionsData.push(option)
            } else {
              const option = {
                authorityId: item.authorityId,
                authorityName: item.authorityName
              }
              optionsData.push(option)
            }
          })
}

const router = useRouter()
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const userInfoUpdateDisable = ref(true)
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getUserListV2({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async() => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    '是否将此用户密码重置为123456?',
    '警告',
    {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning',
    }
  ).then(async() => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    user.authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
  })
}

// const chooseImg = ref(null)
// const openHeaderChange = () => {
//   chooseImg.value.open()
// }

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}

const deleteUserFunc = async(row) => {
  const res = await deleteUser({ id: row.ID })
  if (res.code === 0) {
    ElMessage.success('删除成功')
    row.visible = false
    await getTableData()
  }
}

// 弹窗相关
const userInfo = ref({
  userName: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: [],
  enable: 1,
  validDateTime: [],
  validStart: '',
  validEnd: '',
  userExt: {},
})

const rules = ref({
  userName: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 5, message: '最低5位字符', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入用户密码', trigger: 'blur' },
    { min: 6, message: '最低6位字符', trigger: 'blur' }
  ],
  nickName: [
    { required: true, message: '请输入用户昵称', trigger: 'blur' }
  ],
  standardDataMB: [
    { required: true, message: '请输入用量（MB）', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/, message: '请输入合法手机号', trigger: 'blur' },
  ],
  email: [
    { required: true, pattern: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/, message: '请输入正确的邮箱', trigger: 'blur' },
  ],
  authorityId: [
    { required: true, message: '请选择用户角色', trigger: 'blur' }
  ],
  validDateTime: [
    { required: true, message: '请选择账号有效期', trigger: 'blur' }
  ]
})
const userForm = ref(null)

const enterAddUserDialog = async() => {
  userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await createUserV2(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await updateUserV2(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userInfoUpdateDisable.value = true
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  userInfoUpdateDisable.value = false
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

// const tempAuth = {}
// const changeAuthority = async(row, flag, removeAuth) => {
//   if (flag) {
//     if (!removeAuth) {
//       tempAuth[row.ID] = [...row.authorityIds]
//     }
//     return
//   }
//   await nextTick()
//   const res = await setUserAuthorities({
//     ID: row.ID,
//     authorityIds: row.authorityIds
//   })
//   if (res.code === 0) {
//     ElMessage({ type: 'success', message: '角色设置成功' })
//   } else {
//     if (!removeAuth) {
//       row.authorityIds = [...tempAuth[row.ID]]
//       delete tempAuth[row.ID]
//     } else {
//       row.authorityIds = [removeAuth, ...row.authorityIds]
//     }
//   }
// }

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  addUserDialog.value = true

  nextTick(() => {
    userInfo.value = JSON.parse(JSON.stringify(row))
    userInfo.value.standardDataMB = row.userExt.standardData
    userInfo.value.validDateTime = []
    if (row.userExt.validStart) {
      userInfo.value.validDateTime[0] = dayjs(row.userExt.validStart).tz()
      userInfo.value.validStart = dayjs(row.userExt.validStart).tz().format('YYYY-MM-DD')
    }
    if (row.userExt.validEnd) {
      userInfo.value.validDateTime[1] = dayjs(row.userExt.validEnd).tz()
      userInfo.value.validEnd = dayjs(row.userExt.validEnd).tz().format('YYYY-MM-DD')
    }
    userInfo.value.enable = row.userExt.enable
  })
}

// const switchEnable = async(row) => {
//   userInfo.value = JSON.parse(JSON.stringify(row))
//   await nextTick()
//   const req = {
//     ...userInfo.value
//   }
//   const res = await setUserInfo(req)
//   if (res.code === 0) {
//     ElMessage({ type: 'success', message: `${req.enable === 2 ? '禁用' : '启用'}成功` })
//     await getTableData()
//     userInfo.value.headerImg = ''
//     userInfo.value.authorityIds = []
//   }
// }

const shortcuts = [
  {
    text: '一周',
    value: () => {
      const start = new Date()
      const end = new Date()
      end.setTime(end.getTime() + 3600 * 1000 * 24 * 7)
      return [start, end]
    },
  },
  {
    text: '一月',
    value: () => {
      const start = new Date()
      const end = new Date()
      end.setTime(end.getTime() + 3600 * 1000 * 24 * 30)
      return [start, end]
    },
  },
  {
    text: '三月',
    value: () => {
      const start = new Date()
      const end = new Date()
      end.setTime(end.getTime() + 3600 * 1000 * 24 * 90)
      return [start, end]
    },
  },
]

const validDateTimeChanged = (data) => {
  userInfo.value.validStart = dayjs.tz(data[0]).format('YYYY-MM-DD')
  userInfo.value.validEnd = dayjs.tz(data[1]).format('YYYY-MM-DD')
}

const formatColumnValidDate = (userExtData) => {
  if (!userExtData.validStart || !userExtData.validEnd) {
    return '无限制'
  }
  const start = dayjs(userExtData.validStart).tz().format('YYYY-MM-DD')
  const end = dayjs(userExtData.validEnd).tz().format('YYYY-MM-DD')
  return start + ' ~ ' + end
}

const toTargetWithUUID = (targetView, uuid) => {
  router.push({ name: targetView, query: { uuid: uuid }})
}

const toTargetWithEmail = (targetView, email) => {
  router.push({ name: targetView, query: { email: email }})
}

const copyToClipboard = (textData) => {
  // 使用 clipboard.js 复制文本
  const clipboard = new Clipboard('.el-button', {
    text: () => textData,
  })

  clipboard.on('success', () => {
    console.log('复制成功')
    clipboard.destroy()
  })

  clipboard.on('error', () => {
    console.error('复制失败')
    clipboard.destroy()
  })

  // 触发复制动作
  clipboard.onClick()
}

</script>

<style scoped>
.red-light {
  width: 10px;
  height: 10px;
  background-color: red;
  border-radius: 50%;
  display: inline-block;
}

.green-light {
  width: 10px;
  height: 10px;
  background-color: green;
  border-radius: 50%;
  display: inline-block;
}
</style>
