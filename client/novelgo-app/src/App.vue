<script setup>
import { ref, onMounted, computed } from 'vue';
import GridPoint from '@/components/GridPoint.vue';

const board = ref({
  width: 5,
  height: 2,
  gridPoints: [1, 2, 1, 1, 2, 1, 2, 1, 1, 1]
});

const width = computed(() => board.value.width);
const gridClass = computed(() => `grid grid-cols-${width.value} gap-0`);
let game = {
  Id: 'new-game',
  Name: 'New Game',
  Settings: {
    BoardWidth: 5,
    BoardHeight: 2
  },
  Gameplay: {
    PlayerMoves: []
  }
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
  game.Gameplay.PlayerMoves.push({ 'Row': Math.floor(index / w), 'Col': index % w })
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
  createNewBoard();
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
      <div id="board" class="m-0 justify-center">
        <div :class="gridClass">
          <div
            v-for="(item, index) in board.gridPoints"
            :key="index"
            class="w-12 h-12 flex items-center justify-center"
          >
            <GridPoint :state="item" @click="handleClick(index)"/>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
