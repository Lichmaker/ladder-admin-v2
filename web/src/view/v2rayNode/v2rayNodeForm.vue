<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="uniqueId字段:" prop="uniqueId">
          <el-input v-model="formData.uniqueId" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="configRaw字段:" prop="configRaw">
          <el-input v-model="formData.configRaw" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="name字段:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="secret字段:" prop="secret">
          <el-input v-model="formData.secret" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="secretIv字段:" prop="secretIv">
          <el-input v-model="formData.secretIv" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="vmessPort字段:" prop="vmessPort">
          <el-input v-model.number="formData.vmessPort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="rpcPort字段:" prop="rpcPort">
          <el-input v-model.number="formData.rpcPort" :clearable="true" placeholder="请输入" />
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
  createV2rayNode,
  updateV2rayNode,
  findV2rayNode
} from '@/api/v2rayNode'

defineOptions({
    name: 'V2rayNodeForm'
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
            uniqueId: '',
            configRaw: '',
            name: '',
            secret: '',
            secretIv: '',
            vmessPort: 0,
            rpcPort: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findV2rayNode({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.v2rayNode
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
               res = await createV2rayNode(formData.value)
               break
             case 'update':
               res = await updateV2rayNode(formData.value)
               break
             default:
               res = await createV2rayNode(formData.value)
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
