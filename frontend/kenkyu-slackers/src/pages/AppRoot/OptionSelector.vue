<script setup lang="ts">
  import type { ReviewCustomizeOption } from '@/pages/AppRoot/types';
  import { ref } from 'vue';

  const { values } = defineProps<{
    options: ReviewCustomizeOption,
    values: string[],
    type: "select" | "checkbox",
  }>()
  const valuesRef = ref(values);

</script>

<template>
  <el-select v-model="valuesRef[0]" v-if="type === 'select'">
    <el-option
      v-for="[key, {value, label}] of Object.entries(options)"
      :key
      :value
      :label
    />
  </el-select>
  <el-checkbox-group v-model="valuesRef" v-if="type === 'checkbox'">
    <el-checkbox
      v-for="[key, {value, label}] of Object.entries(options)"
      :key
      :value
      :label
    />
  </el-checkbox-group>
</template>

<style scoped></style>
