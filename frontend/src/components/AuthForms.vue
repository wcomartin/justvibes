<template>
  <div class="auth-container">
    <div class="auth-card">
      <h2 v-if="!showRegister">Login</h2>
      <h2 v-else>Register</h2>
      <form v-if="!showRegister" @submit.prevent="login">
        <div class="form-group">
          <label for="login-email">Email</label>
          <input id="login-email" v-model="loginForm.email" type="email" placeholder="Email" required />
        </div>
        <div class="form-group">
          <label for="login-password">Password</label>
          <input id="login-password" v-model="loginForm.password" type="password" placeholder="Password" required />
        </div>
        <button type="submit" class="primary-btn">Login</button>
        <div v-if="loginError" class="error">{{ loginError }}</div>
        <div v-if="token" class="success">Login successful! Token: {{ token }}</div>
        <div class="switch-link">
          <span>Don't have an account?</span>
          <button type="button" class="link-btn" @click="showRegister = true">Register</button>
        </div>
      </form>
      <form v-else @submit.prevent="register">
        <div class="form-group">
          <label for="register-name">Name</label>
          <input id="register-name" v-model="registerForm.name" placeholder="Name" required />
        </div>
        <div class="form-group">
          <label for="register-email">Email</label>
          <input id="register-email" v-model="registerForm.email" type="email" placeholder="Email" required />
        </div>
        <div class="form-group">
          <label for="register-password">Password</label>
          <input id="register-password" v-model="registerForm.password" type="password" placeholder="Password" required />
        </div>
        <button type="submit" class="primary-btn">Register</button>
        <div v-if="registerError" class="error">{{ registerError }}</div>
        <div v-if="registerSuccess" class="success">Registration successful! Please log in.</div>
        <div class="switch-link">
          <span>Already have an account?</span>
          <button type="button" class="link-btn" @click="showRegister = false">Back to Login</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const showRegister = ref(false)
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
    const res = await fetch('/api/auth/users', {
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
    const res = await fetch('/api/auth/login', {
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
  min-height: 100vh;
  height: 100vh;
  max-height: 100vh;
  overflow: hidden;
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 2rem;
}
.auth-card {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 16px 0 rgba(60, 72, 88, 0.10), 0 1.5px 4px 0 rgba(60, 72, 88, 0.08);
  padding: 2.5rem 2.2rem 2rem 2.2rem;
  width: 360px;
  max-width: 95vw;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  transition: box-shadow 0.2s;
}
.auth-card:hover {
  box-shadow: 0 4px 32px 0 rgba(60, 72, 88, 0.16), 0 2px 8px 0 rgba(60, 72, 88, 0.10);
}
h2 {
  margin-bottom: 1.5rem;
  color: #263238;
  text-align: center;
  font-weight: 500;
  letter-spacing: 0.01em;
}
.form-group {
  margin-bottom: 1.5rem;
  display: flex;
  flex-direction: column;
}
label {
  font-size: 1.02rem;
  color: #607d8b;
  margin-bottom: 0.4rem;
  font-weight: 500;
  letter-spacing: 0.01em;
}
input {
  padding: 0.85rem 1rem 0.85rem 1rem;
  border: none;
  border-radius: 6px 6px 0 0;
  font-size: 1.08rem;
  background: #f5f7fa;
  box-shadow: 0 1.5px 0 0 #bdbdbd;
  color: #263238;
  transition: box-shadow 0.2s, background 0.2s, color 0.2s;
}
input:focus {
  outline: none;
  background: #e3f2fd;
  box-shadow: 0 2px 0 0 #1976d2;
  color: #111;
}
.primary-btn {
  background: #1976d2;
  color: #fff;
  border: none;
  border-radius: 24px;
  padding: 0.85rem 2.2rem;
  font-size: 1.08rem;
  font-weight: 600;
  cursor: pointer;
  margin-top: 0.5rem;
  box-shadow: 0 1.5px 4px 0 rgba(25, 118, 210, 0.10);
  letter-spacing: 0.03em;
  transition: background 0.2s, box-shadow 0.2s;
}
.primary-btn:hover {
  background: #1565c0;
  box-shadow: 0 4px 12px 0 rgba(25, 118, 210, 0.18);
}
.switch-link {
  margin-top: 1.5rem;
  text-align: center;
  color: #607d8b;
  font-size: 0.98rem;
}
.link-btn {
  background: none;
  border: none;
  color: #1976d2;
  font-size: 1rem;
  cursor: pointer;
  text-decoration: underline;
  margin-left: 0.3rem;
  padding: 0;
  font-weight: 500;
  letter-spacing: 0.01em;
  transition: color 0.2s;
}
.link-btn:hover {
  color: #1565c0;
}
.error {
  color: #d32f2f;
  margin-top: 0.7rem;
  font-size: 1.01rem;
  text-align: center;
  letter-spacing: 0.01em;
}
.success {
  color: #388e3c;
  margin-top: 0.7rem;
  font-size: 1.01rem;
  text-align: center;
  letter-spacing: 0.01em;
}
</style>
