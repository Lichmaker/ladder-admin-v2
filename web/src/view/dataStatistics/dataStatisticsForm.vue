<template>
  <div>
    <div class="gva-form-box">
      <el-form
        ref="elFormRef"
        :model="formData"
        label-position="right"
        :rules="rule"
        label-width="80px"
      >
        <el-form-item
          label="userUuid字段:"
          prop="userUuid"
        >
          <el-input
            v-model="formData.userUuid"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="month字段:"
          prop="month"
        >
          <el-input
            v-model="formData.month"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item
          label="value字段:"
          prop="value"
        >
          <el-input
            v-model.number="formData.value"
            :clearable="true"
            placeholder="请输入"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="save"
          >保存</el-button>
          <el-button
            type="primary"
            @click="back"
          >返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createDataStatistics,
  updateDataStatistics,
  findDataStatistics
} from '@/api/dataStatistics'

defineOptions({
  name: 'DataStatisticsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  userUuid: '',
  month: '',
  value: 0,
})
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findDataStatistics({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.redataStatistics
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createDataStatistics(formData.value)
        break
      case 'update':
        res = await updateDataStatistics(formData.value)
        break
      default:
        res = await createDataStatistics(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
