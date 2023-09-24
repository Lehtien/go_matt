<template>
  <div>
    <h1>ましゅー</h1>
    <hr>
    <div class="images">
      <img v-for="image in image_data" :key="image" alt="image" :src="image" @click=select($event)>
    </div>
    <div class="cover" :class={selected}></div>
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
const selected = ref(false)
const select = (el) => {
  if(Array.from(el.target.classList).includes("selected")) {
    el.target.classList = ""
    selected.value = false
  } else {
    el.target.classList = "selected"
    selected.value = true
  }
}
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
@media screen and (max-width: 1024px) {
  .images {
    grid-template-columns: repeat(2, 1fr);
  }
}
img {
  position: relative;
  width: 100%;
  border: 5px solid black;
}
img:hover {
  transform: translate(0, -0.5em);
  transition: transform 0.25s;
}
img.selected {
  position: absolute;
  z-index: 10;
  width:auto;
  max-height: 90%;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  transition: transform 0.25s;
}

.cover.selected {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  z-index: 5;
  background-color: #d3d3d3;
  opacity: 0.5;
}

</style>
