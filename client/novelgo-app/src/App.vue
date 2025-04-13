<script setup>
import { ref, onMounted, computed } from 'vue';
import GridPoint from '@/components/GridPoint.vue';
import GameSettings from '@/components/GameSettings.vue';

let game = {
  Id: '',
  Name: 'Untitled Game',
  Settings: {
    BoardWidth: 5,
    BoardHeight: 5
  },
  Gameplay: {
    PlayerMoves: []
  }
};

const board = ref({
  Id: '',
  width: 0,
  height: 0,
  gridPoints: [],
});

const width = computed(() => board.value.width);
const gridClass = computed(() => `grid grid-cols-${width.value} gap-0`);

const showSettings = ref(false);

const createNewGame = async (settings) => {
  game.Settings.BoardWidth = settings.BoardWidth;
  game.Settings.BoardHeight = settings.BoardHeight;
  showSettings.value = false;
  await createNewBoard();
};

const createNewBoard = async () => {
  try {
    const response = await fetch('http://localhost:58303/games', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(game)
    });

    if (response.ok) {
      game = await response.json();
      board.value = {
        Id: game.Id,
        width: game.Settings.BoardWidth,
        height: game.Settings.BoardHeight,
        gridPoints: game.Gameplay.BoardGridPoints,
      };
    } else {
      console.error('Failed to create a new board');
    }
  } catch (error) {
    console.error('Error:', error);
  }
};

const updateState = async (index) => {
  const w = game.Settings.BoardWidth;
  game.Gameplay.PlayerMoves.push({ 'Row': Math.floor(index / w), 'Col': index % w });
  try {
    const response = await fetch(`http://127.0.0.1:58303/games/${game.Id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(game)
    });

    if (response.ok) {
      console.log('State updated successfully');
      game = await response.json();
      board.value.gridPoints = [...game.Gameplay.BoardGridPoints];
    } else {
      console.error('Failed to update state');
    }
  } catch (error) {
    console.error('Error:', error);
  }
};

const handleClick = (index) => {
  if (board.value.gridPoints[index] <= 1) {
    updateState(index);
  } else {
    alert('invalid');
  }
};

onMounted(() => {
  // createNewBoard();
});
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
      <button v-if="!showSettings" @click="showSettings = true"
        class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >New Game</button>
      <GameSettings v-if="showSettings" @create-game="createNewGame" />
    </div>
    <div v-if="board.Id" class="m-4 text-center flex justify-center">
      <div>{{board.Id}}</div>
    </div>
    <div v-if="board.Id" class="m-4 text-center flex justify-center">
      <div id="board" class="m-4 justify-center relative">
        <div class="grid-lines"></div>
        <div :class="gridClass">
          <div
            v-for="(item, index) in board.gridPoints"
            :key="index"
            class="w-12 h-12 flex items-center justify-center"
          >
            <GridPoint :state="item" @click="handleClick(index)" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.grid-lines {
  position: absolute;
  top: 24px;
  left: 24px;
  width: calc(100% - 48px);
  height: calc(100% - 48px);
  background-image:
    linear-gradient(to right, #8b5a2b 1px, transparent 1px),
    linear-gradient(to bottom, #8b5a2b 1px, transparent 1px);
  background-size: 48px 48px;
  background-position: 0 0, 0 0;
  border-right: 1px solid #8b5a2b;
  border-bottom: 1px solid #8b5a2b
}
</style>
