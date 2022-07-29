<template>
   
       <Markdown :content="content"></Markdown> 
  
    
</template>

<script>
import Markdown from '@/components/markdownvue/markdownvue.vue'
export default {
  name: 'DashboardTable',
  components: {
    Markdown
  },
  data: () => ({
    content: '# 测试aaaa35?'
  }),
  created () {}
}
</script>