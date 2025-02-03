<script setup>
// import { defineProps } from 'vue';
const props = defineProps({
  state: {
    type: Number,
    required: true
  }
});
</script>

<template>
  <div class="font-mono text-center text-gray-400 font-thin">
    <template v-if="state <= 1">+</template>
    <template v-else-if="state === 2">●</template>
    <template v-else-if="state === 3">○</template>
  </div>
</template>

<style scoped>
  .font-thin {
    font-weight: 100;
  }
</style>
