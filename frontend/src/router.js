import { createRouter, createWebHistory } from 'vue-router'
import AuthForms from './components/AuthForms.vue'
import Dashboard from './components/Dashboard.vue'

const routes = [
  { path: '/', component: AuthForms },
  { path: '/dashboard', component: Dashboard },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
