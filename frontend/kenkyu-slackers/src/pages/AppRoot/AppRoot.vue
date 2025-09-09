<script setup lang="ts">
  import type { ReviewCustomizeOption } from "@/pages/AppRoot/types";
  import { computed, onMounted, ref } from "vue";
  import MarkdownEditor from "@/pages/AppRoot/MarkdownEditor.vue";
  import { getFeedbackFromGPT, getIndustryIds } from "@/pages/AppRoot/api-call";
  import MarkdownRenderer from "@/pages/AppRoot/MarkdownRenderer.vue";
  import OptionModal from "@/pages/AppRoot/OptionModal.vue";

  const getOptions = async (optionName: string): Promise<ReviewCustomizeOption> => {
    if (optionName === "industry") {
      const rawData = Object.entries(await getIndustryIds());
      const data: ReviewCustomizeOption = Object.fromEntries(
        rawData.map(e => [e[0], { value: e[0], label: e[1] as string }])
      );
      return data;
    }
    else {
      throw new Error(`Specified option name ${optionName} is not valid.`);
    }
  }

  const markdownContent = ref("");

  const industryOptions = ref<ReviewCustomizeOption>({});
  const fugaOptions = ref<ReviewCustomizeOption>({});
  const industryValues = ref<string[]>([]);
  const fugaValues = ref<string[]>([]);
  onMounted(async () => {
    industryOptions.value = await getOptions("industry");
    fugaOptions.value = await getOptions("industry");
  })

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

  const resultSuggestion = ref("");
  const resultAdvice = ref("");

  const onSubmit = async () => {
    console.log(markdownContent.value);
    const data = await getFeedbackFromGPT();
    resultSuggestion.value = data.improved_press;
    resultAdvice.value = data.Advice;
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
