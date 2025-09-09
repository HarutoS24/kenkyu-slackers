<script setup lang="ts">
  import { basicSetup, EditorView } from 'codemirror';
  import { markdown } from "@codemirror/lang-markdown";
  import { onMounted, useTemplateRef } from 'vue';

  let editorView: EditorView | undefined;
  const mdEditorRef = useTemplateRef("markdown-editor");
  onMounted(() => {
    const mdEditorDOM = mdEditorRef.value;
    if (mdEditorDOM !== null) {
      editorView = new EditorView({
        extensions: [basicSetup, markdown()],
        parent: mdEditorDOM
      });
    }
  });

  const getValue = (): string => {
    if (editorView !== undefined) {
      return editorView.state.doc.toString();
    }
    else {
      return "";
    }
  }
  defineExpose({ getValue });
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
