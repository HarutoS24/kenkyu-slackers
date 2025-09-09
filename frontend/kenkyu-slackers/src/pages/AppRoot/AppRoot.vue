<script setup lang="ts">
  import { computed, ref } from "vue";
  import MarkdownEditor from "@/pages/AppRoot/MarkdownEditor.vue";
  import MarkdownRenderer from "@/pages/AppRoot/MarkdownRenderer.vue";
  import OptionModal from "@/pages/AppRoot/OptionModal.vue";
  import router from "@/router";

  const markdownContent = ref("");
  const industryValues = ref<string[]>([]);
  const fugaValues = ref<string[]>([]);

  const industryLabels = ref<string[]>([]);
  const setIndustryLabels = (v: string[]) => industryLabels.value = v;
  const industryLabelString = computed(() => {
    if (industryLabels.value.length === 0) {
      return "未選択"
    }
    else {
      return industryLabels.value[0];
    }
  });

  const fugaLabels = ref<string[]>([]);
  const setFugaLabels = (v: string[]) => fugaLabels.value = v;
  const fugaLabelString = computed(() => {
    if (fugaLabels.value.length === 0) {
      return "未選択"
    }
    else {
      return fugaLabels.value.join(", ");
    }
  });

  const onSubmit = async () => {
    router.push("/submit");
  }
</script>

<template>
  <div class="container">
    <el-form>
      <el-form-item class="option-modal">
        <option-modal
          v-model:industry="industryValues"
          v-model:fuga="fugaValues"
          @set-industry-label="setIndustryLabels"
          @set-fuga-label="setFugaLabels"
        />
        <div class="line"></div>
        <div class="label">業種</div>
        <div class="value">{{ industryLabelString }}</div>
        <div class="line"></div>
        <div class="label">選択肢2</div>
        <div class="value">{{ fugaLabelString }}</div>
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="本文（マークダウン）" label-position="top">
            <markdown-editor v-model="markdownContent" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="プレビュー" label-position="top">
            <markdown-renderer class="md-renderer" :source="markdownContent" />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">
          送信
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<style scoped>
  .container {
    padding: 0 8vw;
    scrollbar-color: #ccc #fff;
  }

  .option-modal :deep(.el-form-item__content) {
    gap: 15px;
  }

  .option-modal div.label {
    font-size: 0.66rem;
    padding: 4px;
    border: 1px solid #ccc;
    line-height: 1;
  }

  .option-modal div.line {
    height: 15px;
    border-left: 1px solid #999;
  }

  .option-modal div.label + div.value {
    margin-left: -8px;
  }
</style>
