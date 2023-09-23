<template>
  <div>
    <h1>ましゅー</h1>
    <hr>
    <div class="images">
      <img v-for="image in image_data" :key="image" alt="image" :src="image">
    </div>
  </div>
</template>

<script setup>
import axios from 'axios'
import { ref, onMounted } from 'vue'

const image_data = ref([])

onMounted(() => { 
  axios.get('http://localhost/api/')
    .then(response => {
      image_data.value = response.data
    })
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.images {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 1em;
  margin: 0 2em;
}
img {
  width: 100%;
  border: 5px solid black;
}
</style>
