<script setup>
const props = defineProps({
  state: {
    type: Number,
    required: true
  }
});
</script>

<template>
  <div class="board">
    <div v-if="state <= 1" class="stone empty"></div>
    <div v-else-if="state === 2" class="stone black"></div>
    <div v-else-if="state === 3" class="stone white"></div>
  </div>
</template>

<style scoped>
.board {
  width: 48px;
  height: 48px;
  background-color: #dcb35c;
}

.stone {
  width: 42px;
  height: 42px;
  top: 3px;
  left: 3px;
  position: relative;
}

.black {
  background-color: black;
  border-radius: 50%;
}

.white {
  background-color: white;
  border-radius: 50%;
}

.empty {
  background-color: transparent;
}

</style>
