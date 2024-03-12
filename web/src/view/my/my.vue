<template>
  <div class="card-content-container">
    <el-card
      class="usage-card"
    >
      <el-descriptions
        title="账号有效期"
        :column="2"
        direction="vertical"
      >
        <el-descriptions-item
          label="生效日期"
          width="50%"
        >{{ dashboardInfo.base.validStart }}</el-descriptions-item>
        <el-descriptions-item label="到期时间">{{ dashboardInfo.base.validEnd }}</el-descriptions-item>
        <el-descriptions-item label="当前是否可用（流量已耗完或超出有效期都将停用）"> <el-switch
          v-model="dashboardInfo.base.enable"
          size="large"
          active-text="可用"
          inactive-text="停用"
          disabled
        /></el-descriptions-item>
      </el-descriptions>

    </el-card>

    <el-card
      class="usage-card"
    >
      <el-descriptions
        title="当前流量包信息"
        :column="2"
        direction="vertical"
      >
        <el-descriptions-item
          label="开始日期"
          width="50%"
        >{{ dashboardInfo.using.start }}</el-descriptions-item>
        <el-descriptions-item label="结束日期">{{ dashboardInfo.using.end }}</el-descriptions-item>
      </el-descriptions>
      <el-descriptions
        title=""
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
    >
      <el-descriptions
        title="订阅二维码"
        :column="1"
        direction="vertical"
      >
        <el-descriptions-item
          label=""
        >    <el-image
          style="width: 200px; height: 200px"
          :src="dashboardInfo.base.QRCode"
          fit="contain"
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
</template>

<script setup>
import {
  CopyDocument,
} from '@element-plus/icons-vue'
import { ref, onMounted } from 'vue'
import { getDashboardInfo } from '@/api/userExt'

const dashboardInfo = ref({
  base: {},
  using: {},
})

onMounted(async() => {
  const apiData = await getDashboardInfo(true)
  if (apiData.code === 0) {
    dashboardInfo.value = apiData.data
  }
})

defineOptions({
  name: 'My',
})

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
.usage-card  {
    margin-bottom: 10px;
}
</style>
