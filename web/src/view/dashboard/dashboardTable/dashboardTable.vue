<template>
       <Markdown :content="content" />
       <!-- <myvue :content="content"></myvue> -->
</template>

<script>
// import Markdown from '@/components/markdownvue/markdownvue.vue'
// export default {
//   name: 'DashboardTable',
//   components: {
//     Markdown
//   },
//   data: () => ({
//     content: '# 测试aaaa35?'
//   }),
//   created () {}
// }

import Markdown from '@/components/mdpreview/mdpreview.vue'
export default {
  name: 'DashboardTable',
  components: {
    Markdown
  },
  data: () => ({
    content: '测试aaaa35?   ------   `zhegene`'
  }),
  created () {}
}
</script>