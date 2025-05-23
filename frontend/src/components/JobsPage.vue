<template>
  <div class="jobs-layout">
    <aside class="sidebar">
      <div class="logo">JustVibes</div>
      <router-link class="sidebar-link" to="/dashboard">Dashboard</router-link>
      <div class="sidebar-spacer"></div>
      <button class="sidebar-logout-btn" @click="logout">Logout</button>
    </aside>
    <main class="jobs-main">
      <h1 class="jobs-title">Job Management</h1>
      <div class="jobs-summary">
        <div class="summary-card">
          <div class="summary-value">42</div>
          <div class="summary-label">Active Jobs</div>
        </div>
        <div class="summary-card">
          <div class="summary-value">7</div>
          <div class="summary-label">Completed This Week</div>
        </div>
        <div class="summary-card">
          <div class="summary-value">15</div>
          <div class="summary-label">Work Orders Today</div>
        </div>
      </div>
      <div class="jobs-table-container">
        <table class="jobs-table">
          <thead>
            <tr>
              <th>Job ID</th>
              <th>Title</th>
              <th>Status</th>
              <th>Assigned To</th>
              <th>Start Date</th>
              <th>Due Date</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="job in jobs" :key="job.id">
              <td>{{ job.id }}</td>
              <td>{{ job.title }}</td>
              <td>
                <span :class="['status-badge', job.status.toLowerCase()]">{{ job.status }}</span>
              </td>
              <td>{{ job.assignedTo }}</td>
              <td>{{ job.startDate }}</td>
              <td>{{ job.dueDate }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  name: "JobsPage",
  data() {
    return {
      jobs: [
        { id: 101, title: "Install HVAC", status: "Active", assignedTo: "Alice", startDate: "2025-05-20", dueDate: "2025-05-25" },
        { id: 102, title: "Repair Elevator", status: "Active", assignedTo: "Bob", startDate: "2025-05-21", dueDate: "2025-05-26" },
        { id: 103, title: "Inspect Fire Alarm", status: "Completed", assignedTo: "Charlie", startDate: "2025-05-18", dueDate: "2025-05-19" },
        { id: 104, title: "Upgrade Lighting", status: "Active", assignedTo: "Dana", startDate: "2025-05-22", dueDate: "2025-05-28" },
        { id: 105, title: "Repair Plumbing", status: "Active", assignedTo: "Eve", startDate: "2025-05-23", dueDate: "2025-05-27" },
      ]
    };
  },
  methods: {
    logout() {
      localStorage.removeItem('token');
      if (this.$router) {
        this.$router.push('/').catch(() => {});
      } else {
        window.location.href = '/';
      }
    }
  }
};
</script>

<style scoped>
.jobs-layout {
  display: flex;
  max-height: 100vh;
  width: 100vw;
  background: #f5f6fa;
  overflow: hidden;
}
.sidebar {
  width: 280px;
  background: #f3f4f6;
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
.jobs-main {
  flex: 1 1 auto;
  padding: 32px;
  overflow-y: auto;
  min-width: 0;
  min-height: 0;
  max-height: 100vh;
}
.jobs-title {
  font-size: 2rem;
  font-weight: 700;
  color: #222;
  margin-bottom: 32px;
}
.jobs-summary {
  display: flex;
  gap: 24px;
  margin-bottom: 32px;
}
.summary-card {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 8px rgba(30,41,59,0.06);
  padding: 24px 32px;
  min-width: 160px;
  text-align: center;
}
.summary-value {
  font-size: 2rem;
  font-weight: 600;
  color: #3b82f6;
}
.summary-label {
  font-size: 1rem;
  color: #666;
  margin-top: 8px;
}
.jobs-table-container {
  background: #fff;
  border-radius: 16px;
  box-shadow: 0 2px 8px rgba(30,41,59,0.06);
  padding: 24px;
  overflow-x: auto;
}
.jobs-table {
  width: 100%;
  border-collapse: collapse;
}
.jobs-table th, .jobs-table td {
  padding: 12px 16px;
  text-align: left;
}
.jobs-table th {
  background: #f3f4f6;
  color: #222;
  font-weight: 600;
}
.jobs-table tr:nth-child(even) {
  background: #f9fafb;
}
.status-badge {
  display: inline-block;
  padding: 4px 12px;
  border-radius: 999px;
  font-size: 0.95em;
  font-weight: 500;
  color: #fff;
}
.status-badge.active {
  background: #3b82f6;
}
.status-badge.completed {
  background: #10b981;
}
</style>
