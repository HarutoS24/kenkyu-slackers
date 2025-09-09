<script setup lang="ts">
</script>

<template>
  <el-container>
    <el-header>
      <div class="header">
        プレスリリース レビュワー
      </div>
    </el-header>
    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template>

<style scoped>
  :deep(.el-header) {
    padding: 0;
  }

  .header {
    background-color: #335;
    align-content: center;
    user-select: none;

    color: #bbc;
    font-weight: 600;

    margin: 0;
    padding: 0 10vw;
    width: 80vw;
    height: 100%;
  }
</style>
