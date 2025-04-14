<script setup>
import { ref, onMounted, computed } from 'vue';
import GridPoint from '@/components/GridPoint.vue';
import GameSettings from '@/components/GameSettings.vue';

const apiUrl = import.meta.env.VITE_API_URL;

let game = {
  Id: '',
  Name: 'Untitled Game',
  Settings: {
    BoardWidth: 5,
    BoardHeight: 5,
    CyclicLogic: true,
  },
  Gameplay: {
    PlayerMoves: []
  }
};

const board = ref({
  id: '',
  width: 0,
  height: 0,
  cyclicLogic: true,
  gridPoints: [],
});

const width = computed(() => board.value.width);
const gridClass = computed(() => `grid grid-cols-${width.value} gap-0`);

const showSettings = ref(false);
const cursorPosition = ref({ x: 0, y: 0 });
let stoneColor = 'empty';
let moveNumber = ref(0);

const createNewGame = async (settings) => {
  showSettings.value = false;
  moveNumber.value = 0;
  await createNewBoard(settings);
};

const createNewBoard = async (settings) => {
  const newGame = {
    Settings: settings,
    Gameplay: {}
  }
  console.log(apiUrl)
  try {
    const response = await fetch(`${apiUrl}/games`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(newGame)
    });

    if (response.ok) {
      game = await response.json();
      board.value = {
        id: game.Id,
        width: game.Settings.BoardWidth,
        height: game.Settings.BoardHeight,
        gridPoints: game.Gameplay.BoardGridPoints,
        cyclicLogic: game.Settings.CyclicLogic,
      };
      stoneColor = moveNumber.value % 2 == 0 ? 'black' : 'white';
    } else {
      console.error('Failed to create a new board');
    }
  } catch (error) {
    console.error('Error:', error);
  }
};

const updateState = async (index) => {
  const w = game.Settings.BoardWidth;
  if (!game.Gameplay.PlayerMoves) game.Gameplay.PlayerMoves = [];
  game.Gameplay.PlayerMoves.push({ 'Row': Math.floor(index / w), 'Col': index % w });
  try {
    const response = await fetch(`${apiUrl}/games/${game.Id}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(game)
    });

    if (response.ok) {
      game = await response.json();
      board.value.gridPoints = [...game.Gameplay.BoardGridPoints];
      moveNumber.value++;
      stoneColor = moveNumber.value % 2 == 0 ? 'black' : 'white';
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
    // alert('invalid');
  }
};

const handleMouseMove = (event) => {
  cursorPosition.value = { x: event.clientX, y: event.clientY };
};

onMounted(() => {
  window.addEventListener('mousemove', handleMouseMove);
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
    <div v-if="board.id" class="m-4 text-center flex justify-center">
      <div class="text-sm text-gray-400">{{board.id}}</div>
    </div>
    <div v-if="board.id" class="m-4 text-center flex justify-center">
      <div>Cyclic: {{board.cyclicLogic}}</div>
    </div>
    <div v-if="board.id" class="m-4 text-center flex justify-center">
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
    <div v-if="stoneColor" :class="['cursor-stone', stoneColor]" :style="{ top: cursorPosition.y - 24 + 'px', left: cursorPosition.x - 24 + 'px' }"></div>
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

.cursor-stone {
  position: absolute;
  width: 42px;
  height: 42px;
  border-radius: 50%;
  pointer-events: none;
  z-index: 1000;
  opacity: 0.5;
}

.cursor-stone.black {
  background: radial-gradient(circle at 30% 30%, #333, #000);
  box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.5), 0 5px 10px rgba(0, 0, 0, 0.3);
}

.cursor-stone.white {
  background: radial-gradient(circle at 30% 30%, #fff, #ccc);
  box-shadow: inset 0 0 5px rgba(0, 0, 0, 0.5), 0 5px 10px rgba(0, 0, 0, 0.3);
}

</style>
