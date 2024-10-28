<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import '@/assets/styles/global.css'

export default defineComponent({
  name: 'App',
})
</script>

<style>
#app {
  display: flex;
  justify-content: center;
}
</style>
