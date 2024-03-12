<template>
  <el-form
    ref="vForm"
    :model="formData"
    label-position="left"
    label-width="80px"
    size="default"
    @submit.prevent
  >
    <el-form-item>
      <el-button
        type="default"
        @click="back"
      >返回</el-button>
    </el-form-item>

    <el-form-item
      v-if="isUpdate"
      label="ID"
      prop="id"
      class="required"
    >
      <el-input
        v-model="formData.ID"
        type="text"
        disabled
      />
    </el-form-item>
    <el-form-item
      label="标题"
      prop="title"
      class="required"
    >
      <el-input
        v-model="formData.title"
        type="text"
        clearable
      />
    </el-form-item>
    <el-form-item
      label="内容"
      prop="content"
    >
      <div style="z-index:1001; border: 1px solid #ccc">
        <Toolbar
          style="border-bottom: 1px solid #ccc"
          :editor="editorRef"
          :default-config="toolbarConfig"
          mode="default"
        />
        <Editor
          v-model="formData.content"
          style="height: 500px; overflow-y: hidden;"
          :default-config="editorConfig"
          mode="default"
          @onCreated="handleCreated"
        />
      </div>
    </el-form-item>
    <el-form-item>
      <el-button
        type="primary"
        icon="select"
        @click="onSubmit"
      >{{ isUpdate ? '更新' : '新建' }}</el-button>
    </el-form-item>
  </el-form>
</template>

<script setup>
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { useRoute, useRouter } from 'vue-router'
import { onBeforeUnmount, ref, shallowRef, onMounted, onUpdated, onActivated } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { findAnnouncement, updateAnnouncement, createAnnouncement } from '@/api/announcement'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'AnnouncementEdit'
})

const route = useRoute()
const router = useRouter()
const formData = ref({})
let announcementID = 0
const isUpdate = ref(false)

// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef()

onMounted(() => {
  announcementID = route.query.id
  if (announcementID > 0) {
    getData()
  }
})

onUpdated(() => {
  console.log('更新了')
  if (route.query.id !== announcementID) {
    announcementID = route.query.id
    // 更新了id，触发数据更新
    getData()
  }
})

const toolbarConfig = {}
const editorConfig = { placeholder: '请输入内容...' }

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()
})

const handleCreated = (editor) => {
  editorRef.value = editor // 记录 editor 实例，重要！
}

const onSubmit = async() => {
  let res = null
  if (isUpdate.value) {
    res = await updateAnnouncement(formData.value)
  } else {
    res = await createAnnouncement(formData.value)
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    router.push({ name: 'announcement' })
  } else {
    console.log(res)
    ElMessage({
      type: 'error',
      message: '创建/更新失败'
    })
  }
}

const getData = async() => {
  const query = await findAnnouncement({ ID: announcementID })
  console.log(query)

  if (query.code === 0) {
    formData.value = query.data.announcement
    isUpdate.value = true
  } else {
    ElMessage({
      type: 'error',
      message: '查询数据失败，请刷新'
    })
    throw new Error(query.msg)
  }
}

const back = () => {
  router.back()
}

</script>
