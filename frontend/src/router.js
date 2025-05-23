import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from './components/Dashboard.vue';
import JobsPage from './components/JobsPage.vue';
import AuthForms from './components/AuthForms.vue';

const routes = [
  { path: '/', name: 'auth', component: AuthForms },
  { path: '/dashboard', name: 'dashboard', component: Dashboard },
  { path: '/jobs', name: 'jobs', component: JobsPage },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
