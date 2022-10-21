<template>
  <div class="container" style="min-height: 100%; padding-bottom: 100px">
    <div></div>
    <el-container></el-container>
    <el-row :gutter="20">
      <el-col :span="6"></el-col>
      <el-col :span="6"></el-col>
      <el-col :span="6"></el-col>
      <el-col :span="6"></el-col>
    </el-row>
    <el-container>
      <el-header>系统操作（谨慎）</el-header>
      <el-main>
        <el-popconfirm
          confirm-button-text="确认"
          cancel-button-text="点错噜"
          @confirm="requestRestartV2Ray"
          title="将会发指令到服务器重启v2ray，确认配置文件没问题？"
        >
          <template #reference>
            <el-button round type="danger">重启v2ray</el-button>
          </template>
        </el-popconfirm>

        <el-popconfirm
          confirm-button-text="确认"
          cancel-button-text="点错噜"
          @confirm="requestRestartV2Ray"
          title="将会把配置文件写入到服务器中"
        >
          <template #reference>
            <el-button type="danger" round>配置文件写入</el-button>
          </template>
        </el-popconfirm>

        <el-popconfirm
          confirm-button-text="确认"
          cancel-button-text="点错噜"
          @confirm="requestRestartV2Ray"
          title="将会把用户数据同步到配置文件中"
        >
          <template #reference>
            <el-button type="danger" round>同步用户配置</el-button>
          </template>
        </el-popconfirm>
      </el-main>
    </el-container>

    <el-container>
      <el-header>Debug测试</el-header>
      <el-main>
        <el-popconfirm
          confirm-button-text="确认"
          cancel-button-text="点错噜"
          @confirm="requestPlayground"
          title="这是一个playground，执行之前最好确认要下代码和go server是否已重启"
        >
          <template #reference>
            <el-button round type="danger">Playground</el-button>
          </template>
        </el-popconfirm>
      </el-main>
    </el-container>

    <el-dialog
      title="提示"
      v-model="dialogFormVisible"
      width="30%"
      :before-close="closeDialog"
    >
      <span>{{ dialogAttr.noticeMsg }}</span>
      <span slot="footer" class="dialog-footer">
        <!-- <el-button @click="dialogVisible = false">取 消</el-button> -->
        <br />
        <el-button type="primary" @click="closeDialog">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive } from "vue";

import { restartV2Ray, myPlayground } from "@/api/sysControl";

// 弹窗控制标记
const dialogFormVisible = ref(false);

const closeDialog = () => {
  //   type.value = "create";
  dialogFormVisible.value = false;
};

// 弹窗属性
const dialogAttr = ref({
  noticeMsg: "执行中请稍等...",
});

// 重启v2ray，并弹窗
const requestRestartV2Ray = async () => {
  dialogFormVisible.value = true;

  const resp = await restartV2Ray({
    hey: "hello",
  });
  console.log(resp);
  if (resp.code === 0) {
    dialogAttr.value.noticeMsg = resp.msg;
  }
};

// 一个playground
const requestPlayground = async () => {
  dialogFormVisible.value = true;

  const resp = await myPlayground({
    hey: "hello",
  });
  console.log(resp);
  if (resp.code === 0) {
    dialogAttr.value.noticeMsg = resp.msg;
  }
};
</script>

<script>
export default {
  props: [],
  components: {},
  data() {
    return {
      showValue: "开启预览模式后,点击我显示预设逻辑",
      showText: "这里的值声明于预设JS代码",
      dialogVisible: false,
    };
  },
  watch: {},
  computed: {},
  beforeCreate() {},
  created() {},
  beforeMount() {},
  mounted() {},
  beforeUpdate() {},
  updated() {},
  destroyed() {},
  methods: {
    request() {},
    onButtonClick() {},
    hello() {
      alert("来自预设逻辑代码的问候");
    },
  },
  fillter: {},
};
</script>

<style scoped>
.container {
}
</style>
