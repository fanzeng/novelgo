<script setup>
import GridPoint from '@/components/GridPoint.vue';
const width = 4;
const height = 4;
</script>

<template>
<div class="min-h-screen bg-white text-black dark:bg-gray-900 dark:text-white">
  <div class="text-center justify-center">
    <h1 class="text-2xl m-0 p-4">NovelGo</h1>
    <p class="m-2">
      Novel Go Games
    </p>
  </div>
  <div class="text-center justify-center">
    <div id="board" class="m-4">
      <div v-for="rowIndex in height" :key="rowIndex" class="grid-row">
        <GridPoint v-for="colIndex in width" :key="colIndex" />
      </div>
    </div>
  </div>
</div>
</template>

<style scoped></style>
