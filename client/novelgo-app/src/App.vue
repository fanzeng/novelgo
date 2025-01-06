<script setup>
import { computed } from 'vue';
import GridPoint from '@/components/GridPoint.vue';

const width = 5;
const height = 2;
const gridClass = computed(() => `grid grid-cols-${width} gap-0`);
const totalItems = width * height;
</script>

<template>
  <div class="min-h-screen bg-white text-black dark:bg-gray-900 dark:text-white">
    <div class="text-center justify-center grid-cols">
      <h1 class="text-2xl m-0 p-4">NovelGo</h1>
      <p class="m-2">
        Novel Go Games
      </p>
    </div>
    <div class="m-4 text-center flex justify-center">
      <div id="board" class="m-0 justify-center">
        <div :class="gridClass">
          <div
            v-for="(item, index) in totalItems"
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
