<script setup lang="ts">
import axios, { AxiosError } from "axios";
import { ref } from "vue";

  const hoge = () => {
    window.alert("Hoge!");
  }

  const method = ref("");
  const result = ref("");
  const post = async () => {
    await axios.post(`https://httpbun.com/mix/s=${method.value}`).then(res => {
      result.value = `${res.status}`;
    }).catch((res: AxiosError) => {
      result.value = `${res.status}`;
    });
  }
</script>

<template>
  <el-button @click="hoge">Button!</el-button>
  <el-input v-model="method"/>
  <el-button @click="post">Post</el-button>
  <el-input v-model="result" disabled/>
</template>

<style scoped></style>
