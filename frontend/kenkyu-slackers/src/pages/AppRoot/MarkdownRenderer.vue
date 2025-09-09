<script setup lang="ts">
  import VueMarkdownRender from "vue-markdown-render";

  defineProps({
    source: {
      type: String,
      required: true,
    }
  })
</script>

<template>
  <div class="md-renderer-container">
    <vue-markdown-render :source />
  </div>
</template>

<style scoped>
  .md-renderer-container {
    width: 100%;
    height: 50vh;
    overflow-y: scroll;
  }

  .md-renderer-container :deep(img) {
    width: 100%;
  }
</style>
