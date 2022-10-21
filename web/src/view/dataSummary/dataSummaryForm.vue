<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="上行流量:" prop="uplinkByte">
          <el-input v-model.number="formData.uplinkByte" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="下行流量:" prop="downlinkByte">
          <el-input v-model.number="formData.downlinkByte" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="日期:" prop="date">
          <el-date-picker v-model="formData.date" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DataSummary'
}
</script>

<script setup>
import {
  createDataSummary,
  updateDataSummary,
  findDataSummary
} from '@/api/dataSummary'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            uplinkByte: 0,
            downlinkByte: 0,
            date: new Date(),
            username: '',
        })
// 验证规则
const rule = reactive({
               uplinkByte : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               downlinkByte : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               date : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               username : [{
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
      const res = await findDataSummary({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redataSummary
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
               res = await createDataSummary(formData.value)
               break
             case 'update':
               res = await updateDataSummary(formData.value)
               break
             default:
               res = await createDataSummary(formData.value)
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
