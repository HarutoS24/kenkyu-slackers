<script setup lang="ts">
  import MarkdownRenderer from '@/components/MarkdownRenderer.vue';
import { getFeedbackFromGPT } from '@/pages/ResultPage/api-call';
  import { useReviewContentStore } from '@/stores/review-content';
  import { storeToRefs } from 'pinia';
  import { ref } from 'vue';

  const markdownContentStore = useReviewContentStore();
  const { markdownContent, singleIndustryId, importantAspects } = storeToRefs(markdownContentStore);

  const improvedContent = ref("");
  const advice = ref("");

  const asyncSetup = async () => {
    try {
      const result = await getFeedbackFromGPT({
        text: markdownContent.value,
        release_type_id: singleIndustryId.value,
        important_aspects: importantAspects.value,
      });

      advice.value = result.advice;
      improvedContent.value = result.improved_press;

      return true;
    }
    catch (e) {
      // TODO: 例外処理
      throw e;
    }
  }
  await asyncSetup();
</script>

<template>
  <el-form>
    <el-form-item label="アドバイス" label-position="top">
      {{ advice }}
    </el-form-item>
    <el-row :gutter="20">
      <el-col :span="12">
        <el-form-item label="オリジナル" label-position="top">
          <markdown-renderer :source="markdownContent" style="height: 50vh" />
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-form-item label="修正例" label-position="top">
          <markdown-renderer :source="improvedContent" style="height: 50vh" />
        </el-form-item>
      </el-col>
    </el-row>
  </el-form>
</template>

<style scoped>
</style>
