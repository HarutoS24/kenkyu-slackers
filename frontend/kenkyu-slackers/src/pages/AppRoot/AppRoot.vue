<script setup lang="ts">
  import type { ReviewCustomizeOption } from "@/pages/AppRoot/types";
  import OptionSelector from "./OptionSelector.vue"
  import { onMounted, ref } from "vue";
  import MarkdownEditor from "@/pages/AppRoot/MarkdownEditor.vue";
  import { getFeedbackFromGPT, getIndustryIds } from "@/pages/AppRoot/api-call";
  import MarkdownRenderer from "@/pages/AppRoot/MarkdownRenderer.vue";

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
  const industryValue = ref([]);
  const fugaValue = ref([]);
  onMounted(async () => {
    industryOptions.value = await getOptions("industry");
    fugaOptions.value = await getOptions("industry");
  })

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
      <el-form-item label="業種">
        <option-selector v-model="industryValue" :options="industryOptions" type="select" />
      </el-form-item>
      <el-form-item label="選択肢2">
        <option-selector v-model="fugaValue" :options="fugaOptions" type="checkbox" />
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
</style>
