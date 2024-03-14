<template>
  <el-card class="dashboard-container">
    <template #header>
      <div class="header">
        <h1>Welcome Ladder Admin V2!</h1>
      </div>
    </template>

    <el-card class="announcement-card">
      <template #header>
        <div class="announcement-header">
          <h3>近期公告</h3>
        </div>
      </template>
      <div class="demo-collapse">
        <el-collapse
          v-for="(item,index) in announcementList"
          :key="index"
          v-model="collapseActive"
          :accordion="accordionOpt"
        >
          <el-collapse-item
            :title="item.title"
            :name="index"
          >
            <div v-html="item.content" />
          </el-collapse-item>
        </el-collapse>
      </div>
    </el-card>
    <div class="card-content-container">

      <el-card
        class="usage-card"
        style="width: 30%;"
      >
        <el-descriptions
          title="当前流量包信息"
          :column="2"
          direction="vertical"
        >
          <el-descriptions-item
            label="生效日期"
            width="50%"
          >{{ dashboardInfo.using.start }}</el-descriptions-item>
          <el-descriptions-item label="到期时间">{{ dashboardInfo.using.end }}</el-descriptions-item>
        </el-descriptions>
        <el-descriptions
          title="用量"
          :column="2"
          direction="vertical"
        >
          <el-descriptions-item
            label="总量"
            width="50%"
          >{{ dashboardInfo.using.standardData }}</el-descriptions-item>
          <el-descriptions-item label="剩余">{{ dashboardInfo.using.remain }}</el-descriptions-item>
          <el-descriptions-item> <el-progress
            :percentage="dashboardInfo.using.percentage"
            :stroke-width="8"
            striped
            striped-flow
            :duration="20"
          /></el-descriptions-item>
        </el-descriptions>
      </el-card>

      <el-card
        class="qrcode-card"
        style="width: 30%;"
      >
        <el-descriptions
          title="节点二维码"
          :column="1"
          direction="vertical"
        >
          <el-descriptions-item
            label="点击查看大图"
          >  <el-image
            style="width: 100px; height: 100px"
            :src="dashboardInfo.base.QRCode"
            :zoom-rate="1.2"
            :max-scale="7"
            :min-scale="0.2"
            :preview-src-list="QRCodeSrcList"
            :initial-index="4"
            fit="cover"
            close-on-press-escape
            :hide-on-click-modal="true"
          /></el-descriptions-item>
          <el-descriptions-item
            label="订阅链接"
            class="subscribeURL"
          >{{ dashboardInfo.base.subUrl }}
            <el-button
              id="copySubURLButton"
              type="primary"
              :icon="CopyDocument"
              circle
              size="small"
              @click="copySubURLFunc(dashboardInfo.base.subUrl)"
            />
          </el-descriptions-item>

        </el-descriptions>

      </el-card>
    </div>

  </el-card>
</template>

<script setup>
import {
  CopyDocument,
} from '@element-plus/icons-vue'
import { onMounted, ref } from 'vue'
import { getAnnouncementList } from '@/api/announcement'
import { getDashboardInfo } from '@/api/userExt'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'Dashboard',
})

const collapseActive = ref(0) // 控制默认展开第一个

const QRCodeSrcList = ref([])

const dashboardInfo = ref({
  base: {},
  using: {},
})
const announcementList = ref([])
onMounted(async() => {
  const res = await getAnnouncementList({ page: 1, pageSize: 5 })
  console.log(res)
  announcementList.value = res.data.list

  const dashboardInfoQuery = await getDashboardInfo(false)
  if (dashboardInfoQuery.code === 0) {
    dashboardInfo.value = dashboardInfoQuery.data
    console.log(dashboardInfo.value)
    QRCodeSrcList.value.push(dashboardInfo.value.base.QRCode)
    console.log(QRCodeSrcList.value)
  } else {
    console.log(dashboardInfoQuery)
    ElMessage({
      type: 'error',
      msg: '请求失败：' + dashboardInfoQuery.msg
    })
  }
})
const accordionOpt = ref(true)

const copySubURLFunc = (subURL) => {
  if (navigator.clipboard !== undefined) {
    // 浏览器支持
    navigator.clipboard.writeText(subURL).then(function() {
      /* clipboard successfully set */
      console.log('navigator.clipboard 复制成功')
    }, function() {
      /* clipboard write failed */
      console.log('navigator.clipboard 复制失败')
    })
  } else {
    // 浏览器不支持，使用老方法
    const oInput = document.createElement('input')
    oInput.value = subURL
    document.body.appendChild(oInput)
    oInput.select() // 选择对象;
    console.log('使用老方法复制' + oInput.value)
    document.execCommand('Copy') // 执行浏览器复制命令
    oInput.remove()
  }
}
</script>

<style scoped>

.announcement-card {
    margin-bottom: 10px;
}

.card-content-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: flex-start;
}

.card-content-container :deep(.el-card) {
    margin-right: 10px;
}

.dashboard-container :deep(.el-card__header) {
    padding-top: 0;
    padding-bottom: 0;
  border-bottom: 0px;
}

.announcement-card :deep(.el-card__header) {
  margin-top: 30px;
}

</style>
