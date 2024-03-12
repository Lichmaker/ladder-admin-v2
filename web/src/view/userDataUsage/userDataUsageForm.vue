<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="userUuid字段:" prop="userUuid">
          <el-input v-model="formData.userUuid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="startDate字段:" prop="startDate">
          <el-date-picker v-model="formData.startDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="endDate字段:" prop="endDate">
          <el-date-picker v-model="formData.endDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="standardData字段:" prop="standardData">
          <el-input v-model.number="formData.standardData" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="usage字段:" prop="usage">
          <el-input v-model.number="formData.usage" :clearable="true" placeholder="请输入" />
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
  createUserDataUsage,
  updateUserDataUsage,
  findUserDataUsage
} from '@/api/userDataUsage'

defineOptions({
    name: 'UserDataUsageForm'
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
            userUuid: '',
            startDate: new Date(),
            endDate: new Date(),
            standardData: 0,
            usage: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserDataUsage({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reuserDataUsage
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
