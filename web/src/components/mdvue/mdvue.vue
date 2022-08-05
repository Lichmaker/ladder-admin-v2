
<template>
  <v-md-preview :text="text"></v-md-preview>
</template>

<script>
import VMdPreview from "@kangc/v-md-editor/lib/preview";
import "@kangc/v-md-editor/lib/style/preview.css";
// 引入你所使用的主题 此处以 github 主题为例
import githubTheme from "@kangc/v-md-editor/lib/theme/github";
import "@kangc/v-md-editor/lib/theme/style/github.css";

// highlightjs
import hljs from "highlight.js";

VMdPreview.use(githubTheme, {
  Hljs: hljs,
});

export default {
  name: "mdvue",
  props: {
    content: String,
  },
  components: {
    VMdPreview
  },
  data() {
    return {
      text: this.content,
    };
  },
};
</script>