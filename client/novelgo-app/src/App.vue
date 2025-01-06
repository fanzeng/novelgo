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
  <div class="m-4 text-center flex justify-center">
    <div id="board" class="m-0 justify-center">
      <div class="grid grid-cols-4 gap-0 w-64 h-64">
        <div
          v-for="(item, index) in Array(16).fill('')"
          :key="index"
          class="w-12 h-12 flex items-center justify-center"
        >
          <GridPoint :state="index % 4" />
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<style scoped></style>
