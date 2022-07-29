<template>
  <components :is="htmla" class="markdown-body"></components>
</template>

<script>

// markdown组件参考： https://www.cnblogs.com/youxam/p/vue-markdown-render.html

import MarkdownIt from 'markdown-it'
import hljs from 'markdown-it-highlightjs'
import latex from 'markdown-it-katex'
export default {
  name: 'Markdown',
  props: {
    content: String
  },
  data: () => ({
    md: null
  }),
  computed: {
    // 使用 computed 才能在动态绑定时动态更新
    htmla: function () {
              console.log("heyhey")
      let res = this.md.render(this.content)
      console.log(this.content)
      console.log(res)
      // 使用正则表达式将站内链接替换为 router-link 标签, 注意排除锚点跳转的
      res = res.replace(
        /<a href="(?!http:\/\/|https:\/\/)([^#]*?)">(.*?)<\/a>/g,
        '<router-link to="$1">$2</router-link>'
      )
      // 使用正则表达式将站外链接在新窗口中打开
      res = res.replace(/<a href="(.*?)">(.*?)<\/a>/g, '<a href="$1" target="_blank">$2</a>')
      return {
        template: '<div>' + res + '</div>'
      }
    }
  },
  created () {
    this.md = new MarkdownIt()
    this.md.use(hljs).use(latex)
  }
}
</script>
