<script setup lang="ts">
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
        industry_id: singleIndustryId.value,
        important_aspects: importantAspects.value,
      });

      advice.value = result.Advice;
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
  アドバイス
  {{ advice }}
  修正後
  {{ improvedContent }}
</template>

<style scoped>
</style>
