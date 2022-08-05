<template>
  <div class="container" style="min-height: 100%; padding-bottom: 100px;">
    <el-table :data="tableData" style="width: 100%">
      <el-table-column
        :prop="propItem.prop"
        :label="propItem.label"
        v-for="propItem in propList"
        :key="propItem.prop"
      >
        <template v-slot="{ row }">
          <span v-if="!propItem.component">{{ row[propItem.prop] }}</span>
          <component
            v-else
            v-bind:is="propItem.component"
            :rowinfo="row"
          ></component>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
  export default {
    props: [],
    components: {},
    data() {
      return {
        propItem: "",
        propList: [
          {
            label: "姓名",
            prop: "name",
          },
          {
            label: "日期",
            prop: "date",
          },
          {
            label: "地址",
            prop: "address",
          },
          {
            label: "动态组件",
            prop: "lc-component",
            component: "el-switch",
          },
        ],
        row: "",
        tableData: [
          {
            date: "2016-05-02",
            name: "王小虎",
            address: "上海市普陀区金沙江路 1518 弄",
          },
          {
            date: "2016-05-04",
            name: "王小虎",
            address: "上海市普陀区金沙江路 1517 弄",
          },
          {
            date: "2016-05-01",
            name: "王小虎",
            address: "上海市普陀区金沙江路 1519 弄",
          },
          {
            date: "2016-05-03",
            name: "王小虎",
            address: "上海市普陀区金沙江路 1516 弄",
          },
        ],
        showValue: "开启预览模式后,点击我显示预设逻辑",
        showText: "这里的值声明于预设JS代码",
      }
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
        alert("来自预设逻辑代码的问候")
      },
    },
    fillter: {},
  }
</script>