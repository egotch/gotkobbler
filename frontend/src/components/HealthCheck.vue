<template>
  <div>
    <h2>Backend Health Check</h2>
    <!-- Loading State -->
    <div v-if="loading">
      <p>Checking Backend...</p>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="error">
      <p>❌ Error: {{ error }}</p>
    </div>

    <!-- Success State -->
    <div v-else-if="healthData" class="success">
      <p>✅ Status: {{ healthData.status }}</p>
      <p>💾 Database: {{ healthData.database }}</p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

console.log('HealthCheck component is being set up.');

// Reactive state - like variables that update the UI when changed
const loading = ref(true);
const error = ref(null);
const healthData = ref(null);

console.log('Initial state - loading:', loading.value, 'error:', error.value, 'healthData:',
  healthData.value);

// Function to fetch health data from the backend
const checkHealth = async () => {
  console.log('checkHealth function called.');
  try {
    console.log('Starting health check...');
    loading.value = true;
    error.value = null;

    const response = await fetch('/api/health');
    console.log('Response received:', response);
    console.log('Response status:', response.status);
    console.log('Response ok?:', response.ok);
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
  }

    const data = await response.json();
    console.log(' Parsed data:', data);
    healthData.value = data;
    console.log('healthData.value set to:', healthData.value);
  } catch (err) {
    console.error('Error during health check:', err);
    error.value = err.message;
  } finally {
    console.log('Finally block executed. Setting loading to false.');
    loading.value = false;
  }

};
  // Run when the component is mounted
  onMounted(() => {
    console.log('Component mounted. Initiating health check.');
    checkHealth();
  });

</script>

<style scoped>
.health-check {
  padding: 20px;
  border: 2px solid #42b983;
  border-radius: 8px;
  max-width: 400px;
  margin: 20px;
}

.error {
  color: #ff4444;
}

.success {
  color: #42b983;
}
</style>
