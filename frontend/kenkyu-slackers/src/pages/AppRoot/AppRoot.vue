<script setup lang="ts">
  import { computed, ref } from "vue";
  import MarkdownEditor from "@/pages/AppRoot/MarkdownEditor.vue";
  import MarkdownRenderer from "@/pages/AppRoot/MarkdownRenderer.vue";
  import OptionModal from "@/pages/AppRoot/OptionModal.vue";
  import router from "@/router";
  import { useReviewContentStore } from "@/stores/review-content";
  import { storeToRefs } from "pinia";

  const markdownContentStore = useReviewContentStore();
  const { markdownContent, industryIds, importantAspects } = storeToRefs(markdownContentStore);

  const industryLabels = ref<string[]>([]);
  const setIndustryLabels = (v: string[]) => industryLabels.value = v;
  const industryLabelString = computed(() => {
    return industryLabels.value[0];
  });

  const aspectLabels = ref<string[]>([]);
  const setAspectLabels = (v: string[]) => aspectLabels.value = v;
  const aspectLabelString = computed(() => {
    if (aspectLabels.value.length === 0) {
      return "未選択"
    }
    else {
      return aspectLabels.value.join(", ");
    }
  });

  const onSubmit = async () => {
    router.push("/result");
  }
</script>

<template>
  <div class="container">
    <el-form>
      <el-form-item class="option-modal">
        <option-modal
          v-model:industry="industryIds"
          v-model:aspect="importantAspects"
          @set-industry-label="setIndustryLabels"
          @set-aspect-label="setAspectLabels"
        />
        <div class="line"></div>
        <div class="label">業種</div>
        <div class="value">{{ industryLabelString }}</div>
        <div class="line"></div>
        <div class="label">メディアフック</div>
        <div class="value">{{ aspectLabelString }}</div>
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
