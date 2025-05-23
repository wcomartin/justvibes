<template>
  <div class="dashboard-container">
    <h2>Dashboard</h2>
    <div v-if="error" class="error">{{ error }}</div>
    <div v-else-if="loading">Loading...</div>
    <div v-else>
      <pre>{{ dashboardData }}</pre>
    </div>
    <button @click="logout">Logout</button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

const dashboardData = ref(null)
const error = ref('')
const loading = ref(true)
const router = useRouter()

function logout() {
  localStorage.removeItem('token')
  router.push('/')
}

onMounted(async () => {
  loading.value = true
  error.value = ''
  const token = localStorage.getItem('token')
  if (!token) {
    router.push('/')
    return
  }
  try {
    const res = await fetch('/api/dashboard', {
      headers: { 'Authorization': 'Bearer ' + token }
    })
    if (!res.ok) {
      error.value = 'Failed to load dashboard.'
      if (res.status === 401) router.push('/')
      return
    }
    dashboardData.value = await res.json()
  } catch (e) {
    error.value = 'Error loading dashboard.'
  } finally {
    loading.value = false
  }
})
</script>

<style scoped>
.dashboard-container {
  max-width: 600px;
  margin: 2rem auto;
  padding: 2rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  background: #fafafa;
}
pre {
  background: #eee;
  padding: 1rem;
  border-radius: 4px;
}
.error {
  color: #c00;
  margin-bottom: 1rem;
}
button {
  margin-top: 1rem;
}
</style>
