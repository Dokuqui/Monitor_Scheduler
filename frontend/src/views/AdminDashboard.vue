<template>
  <div class="dashboard-container">
    <AdminSidebar />
    <div class="main-content">
      <AdminHeader />
      <!-- Update here -->
      <div class="content">
        <h2 class="dashboard-title">Admin Dashboard</h2>
        <p>
          Welcome to the Admin Dashboard! Here you can manage users, view
          analytics, and more.
        </p>

        <!-- Dummy Stats Cards -->
        <div class="stats-cards">
          <div class="card">
            <h3>Total Users</h3>
            <p>{{ totalUsers }}</p>
          </div>
          <div class="card">
            <h3>Total Tasks</h3>
            <p>{{ totalTasks }}</p>
          </div>
          <div class="card">
            <h3>Pending Approvals</h3>
            <p>{{ pendingApprovals }}</p>
          </div>
        </div>

        <!-- Placeholder for More Features -->
        <div class="features">
          <h3>Manage Users</h3>
          <button @click="navigateToUserManagement">
            Go to User Management
          </button>

          <h3>View Analytics</h3>
          <button @click="navigateToAnalytics">Go to Analytics</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import AdminSidebar from '@/components/AdminSidebar.vue'
import AdminHeader from '@/components/AdminHeader.vue'

export default defineComponent({
  name: 'AdminDashboard',
  components: {
    AdminSidebar,
    AdminHeader,
  },
  data() {
    return {
      totalUsers: 120,
      totalTasks: 350,
      pendingApprovals: 15,
    }
  },
  methods: {
    navigateToUserManagement() {
      this.$router.push('/admin/user-management')
    },
    navigateToAnalytics() {
      this.$router.push('/admin/analytics')
    },
  },
})
</script>

<style scoped>
.dashboard-container {
  display: flex;
}

.main-content {
  flex: 1;
  padding: 2rem;
}

.dashboard-title {
  font-size: 2rem;
  color: #4a2c61; /* Dark purple */
  margin-bottom: 1rem;
}

.stats-cards {
  display: flex;
  justify-content: space-around;
  margin-bottom: 2rem;
}

.card {
  background-color: #f9f1f7; /* Light pink background */
  border-radius: 12px;
  padding: 1.5rem;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  width: 30%;
  text-align: center;
}

.card h3 {
  margin: 0;
  color: #4a2c61; /* Dark purple */
}

.features {
  margin-top: 2rem;
}

.features h3 {
  margin: 1rem 0;
}

.features button {
  background-color: #c69dd8; /* Soft rose */
  color: white;
  border: none;
  border-radius: 6px;
  padding: 0.8rem 1.2rem;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.features button:hover {
  background-color: #b358a1; /* Darker soft rose */
}
</style>
