<script setup lang="ts">
  import ModalWindow from "@/components/ModalWindow.vue";
  import { getIndustryIds } from "@/pages/AppRoot/api-call";
  import OptionSelector from "@/pages/AppRoot/OptionSelector.vue";
  import type { ReviewCustomizeOption } from "@/pages/AppRoot/types";
  import { onMounted, ref, watch } from "vue";

  const getOptions = async (optionName: string): Promise<ReviewCustomizeOption> => {
    return {
      value1: { value: "value1", label: "label1" },
      value2: { value: "value2", label: "label2" },
      value3: { value: "value3", label: "label3" },
    }
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

  const showState = ref(false);

  const industryOptions = ref<ReviewCustomizeOption>({});
  const aspectOptions = ref<ReviewCustomizeOption>({});
  onMounted(async () => {
    industryOptions.value = await getOptions("industry");
    aspectOptions.value = await getOptions("aspect");
  })

  const industryValues = defineModel<string[]>("industry", { required: true });
  const aspectValues = defineModel<string[]>("aspect", { required: true });

  const emit = defineEmits(["setIndustryLabel", "setAspectLabel"]);
  watch(industryValues, newValues => {
    console.log("a");
    emit("setIndustryLabel", newValues.map(e => industryOptions.value[e]?.label ?? `unknown value (${e})`));
  });
  watch(aspectValues, newValues => {
    console.log("b");
    emit("setAspectLabel", newValues.map(e => aspectOptions.value[e]?.label ?? `unknown value (${e})`));
  });
</script>

<template>
  <el-button @click="showState = true">設定</el-button>

  <Teleport to="body">
    <modal-window :show="showState" @close="showState = false">
      <template #body>
        <el-form>
          <el-form-item label="業種">
            <option-selector v-model="industryValues" :options="industryOptions" type="select" />
          </el-form-item>
          <el-form-item label="選択肢2">
            <option-selector v-model="aspectValues" :options="aspectOptions" type="checkbox" />
          </el-form-item>
        </el-form>
      </template>
    </modal-window>
  </Teleport>
</template>

<style scoped>
</style>
