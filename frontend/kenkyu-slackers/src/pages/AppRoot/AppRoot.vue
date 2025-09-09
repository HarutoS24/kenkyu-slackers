<script setup lang="ts">
  import type { ReviewCustomizeOption } from "@/pages/AppRoot/types";
  import OptionSelector from "./OptionSelector.vue"
  import { onMounted, ref } from "vue";
  import MarkdownEditor from "@/pages/AppRoot/MarkdownEditor.vue";

  const getOptions = (optionType: string): ReviewCustomizeOption => {
    // TODO: 本番実装置き換え
    return {
      1: {
        value: `${optionType}-value1`,
        label: `${optionType}-label1`,
      },
      2: {
        value: `${optionType}-value2`,
        label: `${optionType}-label2`,
      },
      3: {
        value: `${optionType}-value3`,
        label: `${optionType}-label3`,
      },
    }
  }

  const hogeOptions = ref<ReviewCustomizeOption>({});
  const fugaOptions = ref<ReviewCustomizeOption>({});
  const hogeValues = ref([]);
  const fugaValues = ref([]);
  const markdownContent = ref("");
  onMounted(() => {
    hogeOptions.value = getOptions("hoge");
    fugaOptions.value = getOptions("fuga");
  })

  const resultSuggestion = ref("");
  const resultAdvice = ref("");
</script>

<template>
  <div class="container">
    <el-form>
      <el-form-item label="選択肢1">
        <option-selector :values="hogeValues" :options="hogeOptions" type="select" />
      </el-form-item>
      <el-form-item label="選択肢2">
        <option-selector :values="fugaValues" :options="fugaOptions" type="checkbox" />
      </el-form-item>
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item label="本文（マークダウン）" label-position="top">
            <markdown-editor :value="markdownContent" />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="変更案" label-position="top">
            <el-input
              v-model="resultSuggestion"
              type="textarea"
              disabled
            />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item>
        <el-button type="primary">
          送信
        </el-button>
      </el-form-item>
    </el-form>
    <el-form>
      <el-form-item label="アドバイス" label-position="top">
        <el-input
          v-model="resultAdvice"
          type="textarea"
          disabled
        />
      </el-form-item>
    </el-form>
  </div>
</template>

<style scoped>
  .container {
    padding: 0 8vw;
  }
</style>
