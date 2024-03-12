<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >

        <!-- <el-form-item
          label="UUID查询"  TODO：做成远程搜索
          prop="uuid"
        >
          <el-select
            v-model="searchInfo.uuid"
            placeholder="Select"
            style="width: 240px"
          >
            <el-option
              v-for="item in userUUIDOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
              :disabled="item.disabled"
            />
          </el-select>
        </el-form-item> -->
        <el-form-item
          label="UUID查询"
          prop="uuid"
        > <el-input
          v-model="searchInfo.uuid"
          clearable
          placeholder="请输入uuid"
          style="width:330px"
        />
        </el-form-item>
        <!-- <el-form-item
          label="创建日期"
          prop="createdAt"
        >
          <template #label>
            <span>
              创建日期
              <el-tooltip
                content="搜索范围是开始日期（包含）至结束日期（不包含）"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="开始日期"
            :disabled-date="
              (time) =>
                searchInfo.endCreatedAt
                  ? time.getTime() > searchInfo.endCreatedAt.getTime()
                  : false
            "
          />
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="结束日期"
            :disabled-date="
              (time) =>
                searchInfo.startCreatedAt
                  ? time.getTime() < searchInfo.startCreatedAt.getTime()
                  : false
            "
          />
        </el-form-item> -->
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >查询</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >新增</el-button>
        <el-popover
          v-model:visible="deleteVisible"
          :disabled="!multipleSelection.length"
          placement="top"
          width="160"
        >
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px">
            <el-button
              type="primary"
              link
              @click="deleteVisible = false"
            >取消</el-button>
            <el-button
              type="primary"
              @click="onDelete"
            >确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <!-- <el-table-column
          align="left"
          label="日期"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column> -->
        <el-table-column
          align="left"
          label="昵称"
          prop="nickname"
          width="120"
        >
          <template #default="scope">{{ scope.row.nickName }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="用户uuid"
          prop="userUuid"
          width="300"
        />
        <el-table-column
          align="left"
          label="开始日期"
          width="100"
        >
          <template #default="scope">{{
            formatOnlyDate(scope.row.startDate)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="结束日期"
          width="100"
        >
          <template #default="scope">{{
            formatOnlyDate(scope.row.endDate)
          }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="标准用量"
          prop="standardData"
          width="120"
        ><template #default="scope">{{ scope.row.standardData }} MiB</template></el-table-column>
        <el-table-column
          align="left"
          label="已用量"
          prop="usage"
          width="120"
        ><template #default="scope">{{
          formatFileSize(scope.row.usage)
        }}</template></el-table-column>
        <el-table-column
          align="left"
          label="已用比例"
          prop="usagePercent"
          width="200"
        >
          <template #default="scope">
            <el-progress
              :percentage="scope.row.progressPercentage"
              :stroke-width="8"
              striped
              striped-flow
              :duration="20"
            /></template>
        </el-table-column>
        <el-table-column
          align="left"
          label="操作"
          min-width="120"
        >
          <template #default="scope">
            <!-- <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              查看详情
            </el-button> -->
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateUserDataUsageFunc(scope.row)"
            >变更</el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
            >删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type === 'create' ? '添加' : '修改'"
      destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          label-width="110px"
        >
          <el-form-item
            label="用户uuid:"
            prop="userUuid"
          >
            <el-input
              v-model="formData.userUuid"
              :clearable="true"
              placeholder="请输入userUuid字段"
              :disabled="type === 'update'"
            />
          </el-form-item>
          <el-form-item
            label="开始日期:"
            prop="startDate"
          >
            <el-date-picker
              v-model="formData.startDate"
              type="date"
              style="width: 100%"
              placeholder="选择日期"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item
            label="结束日期:"
            prop="endDate"
          >
            <el-date-picker
              v-model="formData.endDate"
              type="date"
              style="width: 100%"
              placeholder="选择日期"
              :clearable="true"
            />
          </el-form-item>
          <el-form-item
            label="标准用量(MB):"
            prop="standardData"
          >
            <el-input
              v-model.number="formData.standardData"
              :clearable="true"
              placeholder="请输入standardData字段"
            />
          </el-form-item>
          <el-form-item
            label="已用量(byte):"
            prop="usage"
          >
            <el-input
              v-model.number="formData.usage"
              :clearable="true"
              placeholder="请输入usage字段"
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- <el-dialog
      v-model="detailShow"
      style="width: 800px"
      lock-scroll
      :before-close="closeDetailShow"
      title="查看详情"
      destroy-on-close
    >
      <el-scrollbar height="550px">
        <el-descriptions
          column="1"
          border
        >
          <el-descriptions-item label="userUuid字段">
            {{ formData.userUuid }}
          </el-descriptions-item>
          <el-descriptions-item label="startDate字段">
            {{ formatDate(formData.startDate) }}
          </el-descriptions-item>
          <el-descriptions-item label="endDate字段">
            {{ formatDate(formData.endDate) }}
          </el-descriptions-item>
          <el-descriptions-item label="standardData字段">
            {{ formData.standardData }}
          </el-descriptions-item>
          <el-descriptions-item label="usage字段">
            {{ formData.usage }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog> -->
  </div>
</template>

<script setup>
import {
  createUserDataUsage,
  deleteUserDataUsage,
  deleteUserDataUsageByIds,
  updateUserDataUsage,
  findUserDataUsage,
  getUserDataUsageList,
} from '@/api/userDataUsage'
import { batchGetUserInfoByUUID } from '@/api/user'

// 全量引入格式化工具 请按需保留
import {
//   getDictFunc,
//   formatDate,
  formatOnlyDate,
  //   formatBoolean,
  //   filterDict,
  //   ReturnArrImg,
  //   onDownloadFile,
  formatFileSize,
} from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, onMounted } from 'vue'
import NP from 'number-precision'
import { useRoute } from 'vue-router'

defineOptions({
  name: 'UserDataUsage',
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  userUuid: '',
  startDate: new Date(),
  endDate: new Date(),
  standardData: 0,
  usage: 0,
})

const route = useRoute()

onMounted(() => {
  if (route.query.uuid) {
    searchInfo.value.uuid = route.query.uuid
  }

  loadAllUsers()
  getTableData()
})

// 验证规则
const rule = reactive({})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (
          !searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt
        ) {
          callback(new Error('请填写开始日期'))
        } else if (
          searchInfo.value.startCreatedAt &&
          searchInfo.value.endCreatedAt &&
          (searchInfo.value.startCreatedAt.getTime() ===
            searchInfo.value.endCreatedAt.getTime() ||
            searchInfo.value.startCreatedAt.getTime() >
              searchInfo.value.endCreatedAt.getTime())
        ) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      },
      trigger: 'change',
    },
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getUserDataUsageList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    const userUUIDArr = [] // 拿到用户uuid，去查数据

    for (var i = 0, len = table.data.list.length; i < len; i++) {
      table.data.list[i].progressPercentage = parseProgressPercentage(table.data.list[i]) // 计算进度条数据
      userUUIDArr.push(table.data.list[i].userUuid)
    }

    const userMapData = await batchGetUserInfoByUUID(userUUIDArr)
    for (var i2 = 0, len2 = table.data.list.length; i2 < len2; i2++) {
      const userInfo = userMapData.get(table.data.list[i2].userUuid)
      if (!userInfo) {
        table.data.list[i2].nickName = '未知用户'
      } else {
        table.data.list[i2].nickName = userInfo.nickName
      }
    }

    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(() => {
    deleteUserDataUsageFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据',
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map((item) => {
      ids.push(item.ID)
    })
  const res = await deleteUserDataUsageByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateUserDataUsageFunc = async(row) => {
  const res = await findUserDataUsage({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reuserDataUsage
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteUserDataUsageFunc = async(row) => {
  const res = await deleteUserDataUsage({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功',
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 查看详情控制标记
// const detailShow = ref(false)

// 打开详情弹窗
// const openDetailShow = () => {
//   detailShow.value = true
// }

// 打开详情
// const getDetails = async(row) => {
//   // 打开弹窗
//   const res = await findUserDataUsage({ ID: row.ID })
//   if (res.code === 0) {
//     formData.value = res.data.reuserDataUsage
//     openDetailShow()
//   }
// }

// 关闭详情弹窗
// const closeDetailShow = () => {
//   detailShow.value = false
//   formData.value = {
//     userUuid: '',
//     startDate: new Date(),
//     endDate: new Date(),
//     standardData: 0,
//     usage: 0,
//   }
// }

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    userUuid: '',
    startDate: new Date(),
    endDate: new Date(),
    standardData: 0,
    usage: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createUserDataUsage(formData.value)
        break
      case 'update':
        res = await updateUserDataUsage(formData.value)
        break
      default:
        res = await createUserDataUsage(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功',
      })
      closeDialog()
      getTableData()
    }
  })
}

const parseProgressPercentage = (data) => {
  const dataTotal = NP.times(data.standardData, 1024)
  const res1 = NP.round(NP.divide(NP.divide(data.usage, 1024), dataTotal), 2)
  return NP.times(res1, 100)
}

const userUUIDOptions = ref([])

const loadAllUsers = () => {

}

</script>

<style></style>
