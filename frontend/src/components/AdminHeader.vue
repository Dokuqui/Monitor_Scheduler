<template>
  <header class="admin-header">
    <h1>Admin Dashboard</h1>
    <nav class="admin-nav">
      <ul>
        <li><router-link to="/admin/dashboard">Dashboard</router-link></li>
        <li><router-link to="/admin/users">Manage Users</router-link></li>
        <li><router-link to="/admin/reports">Reports</router-link></li>
        <li>
          <router-link to="/logout" @click.prevent="logout">Logout</router-link>
        </li>
      </ul>
    </nav>
    <div class="user-profile">
      <span>Welcome, Admin</span>
      <!-- You can replace "Admin" with dynamic user info -->
    </div>
  </header>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'AdminHeader',
  setup() {
    const router = useRouter()

    const logout = () => {
      // Implement logout logic (e.g., clear token and redirect)
      localStorage.removeItem('token')
      router.push('/login') // Redirect to login page after logout
    }

    return {
      logout,
    }
  },
})
</script>

<style scoped>
.admin-header {
  margin-bottom: 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #f8f9fa; /* Light background for the header */
  padding: 1rem;
  border-bottom: 1px solid #ccc; /* Subtle border for separation */
}

.admin-nav ul {
  list-style-type: none;
  padding: 0;
  display: flex;
  gap: 1rem; /* Space between nav items */
}

.admin-nav li {
  display: inline;
}

.admin-nav a {
  text-decoration: none;
  color: #007bff; /* Primary link color */
}

.admin-nav a:hover {
  text-decoration: underline; /* Underline on hover */
}

.user-profile {
  font-weight: bold;
  color: #333; /* Darker color for contrast */
}
</style>
