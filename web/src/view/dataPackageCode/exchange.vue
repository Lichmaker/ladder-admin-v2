<template>
  <div>
    <el-card>
      <h1>兑换</h1>

      <el-form
        ref="exchangeFormRef"
        :model="exchangeForm"
        :rules="exchangeRules"
        label-width="80px"
        label-position="top"
      >
        <el-form-item>
          <p>兑换后不会马上转成流量包，可放心兑换！<br>
            兑换后请在下方列表中点击“确认使用”，再到菜单“我的”-->“我的用量数据”中进行核实。<br>
            发生兑换失败或使用失败可邮件联系 lich.wu2014@gmail.com （24小时内必回复）<br></br>
          </p>

        </el-form-item>
        <el-form-item
          label="兑换码"
          prop="code"
        >
          <el-input
            v-model="exchangeForm.code"
            style="width: 30vw;"
            placeholder="请输入兑换码"
          />
        </el-form-item>
        <el-form-item
          label="验证码"
          prop="captcha"
        >
          <el-row
            :gutter="10"
            class="row-bg"
            justify="start"
            style="width: 100%;"
          >
            <el-col :span="5">
              <el-input
                v-model="exchangeForm.captcha"
                placeholder="请输入验证码"
              />
            </el-col>

            <el-col :span="6">
              <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                <img
                  v-if="picPath"
                  class="w-full h-full"
                  :src="picPath"
                  alt="请输入验证码"
                  @click="captchaVerify()"
                >
              </div>
              <!-- <el-button
                type="primary"
                @click="handleCaptchaRefresh"
              >刷新验证码</el-button> -->
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            :disabled="exchangeBtnDisable"
            @click="handleExchangeSubmit"
          >立即兑换</el-button>
        </el-form-item>
      </el-form>
    </el-card>
    <el-divider border-style="dashed" />
    <el-card>
      <div class="gva-table-box">
        <el-table
          ref="multipleTable"
          style="width: 100%"
          tooltip-effect="dark"
          :data="tableData"
          row-key="ID"
          table-layout="auto"
        >
          <el-table-column
            label="兑换码"
            prop="uniqueCode"
          />
          <el-table-column
            label="兑换时间"
            prop="usedAt"
          >
            <template #default="scope">
              {{ dayjs(scope.row.usedAt).tz().format('YYYY-MM-DD H:m:s') }}
            </template>
          </el-table-column>
          <el-table-column
            label="使用时间"
            prop="consumedAt"
          >
            <template #default="scope">
              {{ scope.row.consumedAt == null ? '未使用' : dayjs(scope.row.usedAt).tz().format('YYYY-MM-DD  H:m:s') }}
            </template>
          </el-table-column>
          <el-table-column
            label="基础流量(MB)"
            prop="standardData"
          />
          <el-table-column
            label="有效天数"
            prop="days"
          />
          <el-table-column
            align="left"
            label="操作"
            min-width="120"
          > <template #default="scope">

            <el-button
              v-if="scope.row.consumedAt == null"
              type="primary"
              link
              :icon="SuccessFilled"
              class="table-button"
              :disabled="consumeBtnDisable"
              @click="consumePackageByID(scope.row.ID)"
            >确认使用</el-button>

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
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, inject } from 'vue'
import { captcha } from '@/api/user'
import { exchangeDataPackageCode, getMyDataPackageCodeList, consumePackage } from '@/api/dataPackageCode'
import { ElMessage } from 'element-plus'
import { onMounted } from 'vue'
import { SuccessFilled } from '@element-plus/icons-vue'

const dayjs = inject('$dayjs')
// 获取验证码
const captchaVerify = () => {
  captcha({}).then(async(ele) => {
    exchangeRules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur',
    })
    picPath.value = ele.data.picPath
    exchangeForm.captchaId = ele.data.captchaId
    exchangeForm.openCaptcha = ele.data.openCaptcha
  })
}
captchaVerify()

const exchangeFormRef = ref(null)
const exchangeForm = reactive({
  code: '',
  captcha: ''
})
const picPath = ref('')
const exchangeRules = {
  code: [{ required: true, message: '请输入兑换码', trigger: 'blur' }],
  captcha: [{ required: true, message: '请输入验证码', trigger: 'blur' }]
}
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const consumeBtnDisable = ref(false)
const exchangeBtnDisable = ref(false)

const handleExchangeSubmit = () => {
  exchangeBtnDisable.value = true
  exchangeFormRef.value?.validate(async(valid) => {
    if (!valid) {
      exchangeBtnDisable.value = false
      return
    }

    const res = await exchangeDataPackageCode(exchangeForm)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '兑换成功',
      })
      getTableData()
    }
    captchaVerify()
    exchangeBtnDisable.value = false
  })
}

const getTableData = async() => {
  const queryValue = await getMyDataPackageCodeList({
    page: page.value,
    pageSize: pageSize.value,
  })
  if (queryValue.code === 0) {
    tableData.value = queryValue.data.list
    total.value = queryValue.data.total
    page.value = queryValue.data.page
    pageSize.value = queryValue.data.pageSize
  }
}

onMounted(() => {
  getTableData()
})

const consumePackageByID = async(packageID) => {
  consumeBtnDisable.value = true
  const req = reactive({
    packageID: packageID
  })

  const result = await consumePackage(req)
  if (result.code === 0) {
    ElMessage({
      type: 'success',
      message: '使用成功',
    })
    getTableData()
  }
  consumeBtnDisable.value = false
}

</script>
