<template>
  <div class="sidebar">
    <h2>Admin Dashboard</h2>
    <ul>
      <li><router-link to="/admin/dashboard">Home</router-link></li>
      <li><router-link to="/admin/users">Users</router-link></li>
      <li><router-link to="/admin/settings">Settings</router-link></li>
      <!-- Add more links as needed -->
    </ul>
  </div>
</template>

<script lang="ts">
export default {
  name: 'AdminSidebar',
}
</script>

<style scoped>
.sidebar {
  width: 250px;
  background-color: #f9f1f7; /* Light pink background */
  padding: 1rem;
  height: 100vh; /* Full height */
  position: fixed;
  top: 0;
  left: 0;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.1);
}

.sidebar h2 {
  color: #4a2c61; /* Dark purple */
}

.sidebar ul {
  list-style: none;
  padding: 0;
}

.sidebar li {
  margin: 1rem 0;
}

.sidebar a {
  text-decoration: none;
  color: #4a2c61; /* Dark purple */
}

.sidebar a:hover {
  text-decoration: underline;
}
</style>
