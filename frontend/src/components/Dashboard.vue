<template>
  <div class="dashboard-layout">
    <aside class="sidebar">
      <div class="logo">JustVibes</div>
      <router-link class="sidebar-link" to="/jobs">Jobs</router-link>
      <div class="sidebar-spacer"></div>
      <button class="sidebar-logout-btn" @click="logout">Logout</button>
    </aside>
    <main class="dashboard-main">
      <h1 class="dashboard-title">Welcome to JustVibes Dashboard</h1>
      <div class="dashboard-stats">
        <div class="stat-card">
          <div class="stat-value">42</div>
          <div class="stat-label">Active Jobs</div>
        </div>
        <div class="stat-card">
          <div class="stat-value">15</div>
          <div class="stat-label">Work Orders Today</div>
        </div>
        <div class="stat-card">
          <div class="stat-value">7</div>
          <div class="stat-label">Completed This Week</div>
        </div>
      </div>
      <div class="dashboard-charts">
        <div class="chart-container">
            <h2>Jobs Over Time</h2>
            <div class="canvas-wrapper"><canvas id="lineChart"></canvas></div>
        </div>
        <div class="chart-container">
            <h2>Work Orders by Type</h2>
            <div class="canvas-wrapper"><canvas id="barChart"></canvas></div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { onMounted } from 'vue';
import Chart from 'chart.js/auto';

export default {
  name: "Dashboard",
  methods: {
    logout() {
      localStorage.removeItem('token');
      // Always redirect to root (login is on the root page)
      if (this.$router) {
        this.$router.push('/').catch(() => {});
      } else {
        window.location.href = '/';
      }
    }
  },
  mounted() {
    // Line Chart: Jobs Over Time
    const lineCtx = document.getElementById('lineChart').getContext('2d');
    new Chart(lineCtx, {
      type: 'line',
      data: {
        labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        datasets: [{
          label: 'Jobs',
          data: [5, 8, 6, 10, 7, 3, 4],
          borderColor: '#3b82f6',
          backgroundColor: 'rgba(59,130,246,0.1)',
          tension: 0.4,
          fill: true
        }]
      },
      options: {
        responsive: true,
        plugins: { legend: { display: false } }
      }
    });
    // Bar Chart: Work Orders by Type
    const barCtx = document.getElementById('barChart').getContext('2d');
    new Chart(barCtx, {
      type: 'bar',
      data: {
        labels: ['Install', 'Repair', 'Inspection', 'Upgrade'],
        datasets: [{
          label: 'Work Orders',
          data: [6, 4, 3, 2],
          backgroundColor: ['#3b82f6', '#10b981', '#f59e42', '#6366f1'],
        }]
      },
      options: {
        responsive: true,
        plugins: { legend: { display: false } }
      }
    });
  }
};
</script>

<style scoped>
html, body, #app {
  height: 100%;
  margin: 0;
  padding: 0;
}
.dashboard-layout {
  display: flex;
  max-height: 100vh;
  width: 100vw;
  background: #f5f6fa;
  overflow: hidden;
}
.sidebar {
  width: 280px;
  background: #f3f4f6; /* Tailwind's gray-100 */
  color: #222;
  display: flex;
  flex-direction: column;
  align-items: stretch;
  padding: 0 0 24px 0;
  box-shadow: 2px 0 8px rgba(30,41,59,0.04);
  min-height: 100vh;
  flex-shrink: 0;
}
.logo {
  font-size: 1.6rem;
  font-weight: bold;
  padding: 32px 0 24px 32px;
  letter-spacing: 1px;
  color: #3b82f6;
}
.sidebar-link {
  display: block;
  color: #2563eb;
  font-weight: 600;
  padding: 12px 32px 12px 32px;
  text-decoration: none;
  border-radius: 8px;
  margin: 8px 0 0 0;
  transition: background 0.2s, color 0.2s;
}
.sidebar-link:hover, .sidebar-link.router-link-exact-active {
  background: #e0e7ff;
  color: #1e293b;
}
.sidebar-spacer {
  flex: 1 1 auto;
}
.sidebar-logout-btn {
  margin: 0 24px 24px 24px;
  padding: 12px 0;
  border: none;
  border-radius: 999px;
  background: #3b82f6;
  color: #fff;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s;
}
.sidebar-logout-btn:hover {
  background: #ef4444;
}
.dashboard-main {
  flex: 1 1 auto;
  padding: 32px;
  overflow-y: auto;
  min-width: 0;
  min-height: 0;
  max-height: 100vh;
}
.dashboard-title {
  font-size: 2.2rem;
  font-weight: 700;
  margin-bottom: 32px;
  color: #222;
}
.dashboard-stats {
  display: flex;
  gap: 24px;
  margin-bottom: 40px;
}
.stat-card {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 8px rgba(30,41,59,0.06);
  padding: 24px 32px;
  min-width: 160px;
  text-align: center;
}
.stat-value {
  font-size: 2rem;
  font-weight: 600;
  color: #3b82f6;
}
.stat-label {
  font-size: 1rem;
  color: #666;
  margin-top: 8px;
}
.dashboard-charts {
  display: grid;
  grid-template-columns: repeat(2, minmax(260px, 1fr));
  gap: 24px;
  justify-items: stretch;
  align-items: start;
}
.canvas-wrapper {
  height: 300px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.dashboard-charts .chart-container {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 8px rgba(30,41,59,0.06);
  padding: 8px;
  display: block;
}
.dashboard-charts .chart-container h2 {
  font-size: 1.3rem;
  font-weight: 700;
  color: #2563eb;
  margin: 0 0 16px 16px;
  text-align: left;
  letter-spacing: 0.5px;
  text-shadow: 0 2px 8px rgba(30,41,59,0.06);
  padding-left: 8px;
}
</style>
