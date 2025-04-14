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
  box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.5), 0 5px 10px rgba(0, 0, 0, 0.3);
  background: radial-gradient(circle at 30% 30%, #333, #000);
}

.white {
  background-color: #ddd;
  border-radius: 50%;
  box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.5), 0 5px 10px rgba(0, 0, 0, 0.3);
  background: radial-gradient(circle at 30% 30%, #ddd, #ccc);
}

.stone.black::before {
  content: '';
  position: absolute;
  top: 20%;
  left: 20%;
  width: 30%;
  height: 24%;
  background: radial-gradient(circle at 30% 30%, rgba(150, 150, 150, 0.6), transparent);
  border-radius: 50%;
  pointer-events: none;
}

.stone.white::before {
  content: '';
  position: absolute;
  top: 20%;
  left: 20%;
  width: 30%;
  height: 24%;
  background: radial-gradient(circle at 30% 30%, rgba(255, 255, 255, 0.6), transparent);
  border-radius: 50%;
  pointer-events: none;
}

.empty {
  background-color: transparent;
}

</style>
