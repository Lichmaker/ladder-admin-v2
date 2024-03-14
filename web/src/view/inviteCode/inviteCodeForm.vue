<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="邀请码:" prop="code">
          <el-input v-model="formData.code" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建人uuid:" prop="creatorUuid">
          <el-input v-model="formData.creatorUuid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="使用者uuid:" prop="newUuid">
          <el-input v-model="formData.newUuid" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="freeData字段:" prop="freeData">
          <el-input v-model.number="formData.freeData" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="freeDays字段:" prop="freeDays">
          <el-input v-model.number="formData.freeDays" :clearable="true" placeholder="请输入" />
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
  createInviteCode,
  updateInviteCode,
  findInviteCode
} from '@/api/inviteCode'

defineOptions({
    name: 'InviteCodeForm'
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
            code: '',
            creatorUuid: '',
            newUuid: '',
            freeData: 0,
            freeDays: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findInviteCode({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reinviteCode
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
               res = await createInviteCode(formData.value)
               break
             case 'update':
               res = await updateInviteCode(formData.value)
               break
             default:
               res = await createInviteCode(formData.value)
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
