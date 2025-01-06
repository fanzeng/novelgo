<script setup>
import GridPoint from '@/components/GridPoint.vue'
</script>

<template>
<div class="min-h-screen bg-white text-black dark:bg-gray-900 dark:text-white">
  <h1 class="text-2xl">NovelGo</h1>
  <p>
    Novel Go Games
  </p>
  <div id="board">
    <div>
      <GridPoint /><GridPoint /><GridPoint />
    </div>
    <div>
      <GridPoint /><GridPoint /><GridPoint />
    </div>
    <div>
      <GridPoint /><GridPoint /><GridPoint />
    </div>
  </div>
</div>
</template>

<style scoped></style>
