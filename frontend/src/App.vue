<template>
  <div id="app">
    <h1>Gotkobbler - Recipe Catalog</h1>
    <HealthCheck />
  </div>
</template>

<script setup>
import HealthCheck from './components/HealthCheck.vue'
</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
