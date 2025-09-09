<script setup lang="ts">
  import type { ReviewCustomizeOption } from '@/pages/AppRoot/types';
  import { computed } from 'vue';

  defineProps<{
    options: ReviewCustomizeOption,
    type: "select" | "checkbox",
  }>()

  const values = defineModel<string[]>({ required: true });
  const singleValue = computed({
    get: () => values.value[0],
    set: (newValue) => {
      values.value = [newValue];
    },
  });
</script>

<template>
  <el-select v-model="singleValue" v-if="type === 'select'">
    <el-option
      v-for="[key, {value, label}] of Object.entries(options)"
      :key
      :value
      :label
    />
  </el-select>
  <el-checkbox-group v-model="values" v-if="type === 'checkbox'">
    <el-checkbox
      v-for="[key, {value, label}] of Object.entries(options)"
      :key
      :value
      :label
    />
  </el-checkbox-group>
</template>

<style scoped></style>
