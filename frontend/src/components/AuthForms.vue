<template>
  <div class="auth-container">
    <h2>Register</h2>
    <form @submit.prevent="register">
      <input v-model="registerForm.name" placeholder="Name" required />
      <input v-model="registerForm.email" type="email" placeholder="Email" required />
      <input v-model="registerForm.password" type="password" placeholder="Password" required />
      <button type="submit">Register</button>
      <div v-if="registerError" class="error">{{ registerError }}</div>
      <div v-if="registerSuccess" class="success">Registration successful! Please log in.</div>
    </form>

    <h2>Login</h2>
    <form @submit.prevent="login">
      <input v-model="loginForm.email" type="email" placeholder="Email" required />
      <input v-model="loginForm.password" type="password" placeholder="Password" required />
      <button type="submit">Login</button>
      <div v-if="loginError" class="error">{{ loginError }}</div>
      <div v-if="token" class="success">Login successful! Token: {{ token }}</div>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const registerForm = ref({ name: '', email: '', password: '' })
const loginForm = ref({ email: '', password: '' })
const registerError = ref('')
const registerSuccess = ref(false)
const loginError = ref('')
const token = ref('')

async function register() {
  registerError.value = ''
  registerSuccess.value = false
  try {
    const res = await fetch('/api/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(registerForm.value),
    })
    if (!res.ok) {
      const data = await res.json().catch(() => ({}))
      registerError.value = data.error || res.statusText
      return
    }
    registerSuccess.value = true
    registerForm.value = { name: '', email: '', password: '' }
  } catch (e) {
    registerError.value = 'Registration failed.'
  }
}

async function login() {
  loginError.value = ''
  token.value = ''
  try {
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(loginForm.value),
    })
    if (!res.ok) {
      const data = await res.json().catch(() => ({}))
      loginError.value = data.error || res.statusText
      return
    }
    const data = await res.json()
    token.value = data.token
    localStorage.setItem('token', data.token)
    loginForm.value = { email: '', password: '' }
    // Redirect to dashboard
    window.location.href = '/dashboard'
  } catch (e) {
    loginError.value = 'Login failed.'
  }
}
</script>

<style scoped>
.auth-container {
  max-width: 400px;
  margin: 2rem auto;
  padding: 2rem;
  border: 1px solid #ddd;
  border-radius: 8px;
  background: #fafafa;
}
input {
  display: block;
  width: 100%;
  margin-bottom: 1rem;
  padding: 0.5rem;
  font-size: 1rem;
}
button {
  padding: 0.5rem 1.5rem;
  font-size: 1rem;
  margin-bottom: 1rem;
}
.error {
  color: #c00;
  margin-bottom: 1rem;
}
.success {
  color: #080;
  margin-bottom: 1rem;
}
</style>
