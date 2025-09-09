<script setup lang="ts">
  import { basicSetup, EditorView } from 'codemirror';
  import { markdown } from "@codemirror/lang-markdown";
  import { onMounted, useTemplateRef } from 'vue';

  // コンポーネント外から読む（リアクティブ）
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const value = defineModel({
    type: String
  });

  const mdEditorRef = useTemplateRef("markdown-editor");
  onMounted(() => {
    const mdEditorDOM = mdEditorRef.value;
    if (mdEditorDOM !== null) {
      new EditorView({
        extensions: [
          basicSetup,
          markdown(),
          EditorView.updateListener.of((update) => {
            if (update.docChanged) {
              value.value = update.state.doc.toString();
            }
          }),
        ],
        parent: mdEditorDOM
      });
    }
  });
</script>

<template>
  <div class="md-container">
    <div ref="markdown-editor"></div>
  </div>
</template>

<style scoped>
  .md-container {
    width: 100%;
  }

  :deep(.cm-editor), :deep(.cm-scroller) {
    height: 50vh;
  }

  :deep(.cm-gutter), :deep(.cm-content) {
    min-height: 50vh;
  }
</style>
