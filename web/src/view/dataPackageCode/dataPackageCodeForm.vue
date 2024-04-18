<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="唯一编号:" prop="uniqueCode">
          <el-input v-model="formData.uniqueCode" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="基础流量(MB):" prop="standardData">
          <el-input v-model.number="formData.standardData" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="有效天数:" prop="days">
          <el-input v-model.number="formData.days" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否启用:" prop="enable">
          <el-switch v-model="formData.enable" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="被使用时间:" prop="usedAt">
          <el-date-picker v-model="formData.usedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="备注:" prop="remark">
          <el-input v-model="formData.remark" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createDataPackageCode,
  updateDataPackageCode,
  findDataPackageCode
} from '@/api/dataPackageCode'

defineOptions({
    name: 'DataPackageCodeForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            uniqueCode: '',
            standardData: 0,
            days: 0,
            enable: false,
            usedAt: new Date(),
            remark: '',
        })
// 验证规则
const rule = reactive({
               uniqueCode : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               standardData : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               days : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               enable : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDataPackageCode({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redataPackageCode
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createDataPackageCode(formData.value)
               break
             case 'update':
               res = await updateDataPackageCode(formData.value)
               break
             default:
               res = await createDataPackageCode(formData.value)
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
