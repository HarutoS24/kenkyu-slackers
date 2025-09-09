<script setup lang="ts">
  import type { ReviewCustomizeOption } from "@/pages/AppRoot/types";
  import OptionSelector from "./OptionSelector.vue"
  import { onMounted, ref } from "vue";
  import MarkdownEditor from "@/pages/AppRoot/MarkdownEditor.vue";
  import { getAxios } from "@/util/axios";

  const getOptions = async (optionName: string): Promise<ReviewCustomizeOption> => {
    const endpointMap = new Map([
      ["industry", "/industry_ids"]
    ])

    if (endpointMap.has(optionName)) {
      const res = await getAxios().get(endpointMap.get(optionName)!);
      if (res.status === 200) {
        const rawData = Object.entries(res.data);
        const data: ReviewCustomizeOption = Object.fromEntries(
          rawData.map(e => [e[0], { value: e[0], label: e[1] as string }])
        );
        return data;
      }
      else {
        throw new Error(`The API response was invalid (status ${res.status}).`);
      }
    }
    else {
      throw new Error(`Specified option name ${optionName} is not valid.`);
    }
  }

  const industryOptions = ref<ReviewCustomizeOption>({});
  const fugaOptions = ref<ReviewCustomizeOption>({});
  const industryValue = ref([]);
  const fugaValue = ref([]);
  const markdownContent = ref("");
  onMounted(async () => {
    industryOptions.value = await getOptions("industry");
    fugaOptions.value = await getOptions("industry");
  })

  const resultSuggestion = ref("");
  const resultAdvice = ref("");
</script>

<template>
  <div class="container">
    <el-form>
      <el-form-item label="業種">
        <option-selector :values="industryValue" :options="industryOptions" type="select" />
      </el-form-item>
      <el-form-item label="選択肢2">
        <option-selector :values="fugaValue" :options="fugaOptions" type="checkbox" />
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
