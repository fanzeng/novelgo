
<script setup>
import { ref } from 'vue';

const settings = ref({
  BoardWidth: 5,
  BoardHeight: 5,
  CyclicLogic: true,
  GameMode: 'ComputerOpponent',
});

const emit = defineEmits(['create-game']);

const submitSettings = () => {
  emit('create-game', settings.value);
};
</script>

<template>
  <div class="game-settings">
    <h2>Game Settings</h2>
    <form @submit.prevent="submitSettings">
      <div>
        <label for="boardWidth">Board Width:</label>
        <input type="number" v-model.number="settings.BoardWidth" id="boardWidth" min="2" max="19" required
          class="bg-white dark:bg-gray-800 text-black dark:text-white border border-gray-300 dark:border-gray-600 rounded m-2 p-2"
        />
      </div>
      <div>
        <label for="boardHeight">Board Height:</label>
        <input type="number" v-model.number="settings.BoardHeight" id="boardHeight" min="2" max="19" required
          class="bg-white dark:bg-gray-800 text-black dark:text-white border border-gray-300 dark:border-gray-600 rounded m-2 p-2"
        />
      </div>
      <div class="game-mode-container">
        <label>Game Mode</label>
        <div class="radio-group">
          <div class="radio-item">
            <input type="radio" v-model="settings.GameMode" id="ComputerOpponent" value="ComputerOpponent"
              class="bg-white dark:bg-gray-800 text-black dark:text-white border border-gray-300 dark:border-gray-600 rounded m-2 p-2"
            />
            <label for="ComputerOpponent">Computer Opponent</label>
          </div>
          <div class="radio-item">
            <input type="radio" v-model="settings.GameMode" id="SelfStudy" value="SelfStudy"
              class="bg-white dark:bg-gray-800 text-black dark:text-white border border-gray-300 dark:border-gray-600 rounded m-2 p-2"
            />
            <label for="SelfStudy">Self Study</label>
          </div>
        </div>
      </div>
      <div>
        <label for="cyclicLogic">Cyclic Logic:</label>
        <input type="checkbox" v-model="settings.CyclicLogic" id="cyclicLogic"
          class="bg-white dark:bg-gray-800 text-black dark:text-white border border-gray-300 dark:border-gray-600 rounded m-2 p-2"
        />
      </div>
      <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">Go</button>
    </form>
  </div>
</template>
<style scoped>
.game-mode-container {
  display: flex;
  flex-direction: column;
  margin: 10px;
}

.radio-group {
  display: flex;
  flex-direction: row;
  gap: 10px;
}

.radio-item {
  display: flex;
  align-items: center;
  gap: 5px;
}
</style>
