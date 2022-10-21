<template>
  <div class="container" style="min-height: 100%; padding-bottom: 100px">
    <el-result icon="info" title="信息提示" subtitle="请根据提示进行操作">
      <template #extra>
        <el-button type="primary" size="medium">返回</el-button>
      </template>
    </el-result>
  </div>
</template>

<script>
export default {
  props: [],
  components: {},
  data() {
    return {
      showValue: "开启预览模式后,点击我显示预设逻辑",
      showText: "这里的值声明于预设JS代码",
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
