<template>
  <div class="signup-container">
    <h1 class="signup-title">Create Your Account</h1>
    <form @submit.prevent="signup" class="signup-form">
      <div class="input-group">
        <input
          v-model="username"
          placeholder=" "
          class="input-field"
          required
        />
        <label class="input-label">Username</label>
      </div>
      <div class="input-group">
        <input
          v-model="lastname"
          placeholder=" "
          class="input-field"
          required
        />
        <label class="input-label">Last Name</label>
      </div>
      <div class="input-group">
        <input
          v-model="firstname"
          placeholder=" "
          class="input-field"
          required
        />
        <label class="input-label">First Name</label>
      </div>
      <div class="input-group">
        <input
          v-model="password"
          :type="showPassword ? 'text' : 'password'"
          placeholder=" "
          class="input-field"
          required
        />
        <label class="input-label">Password</label>
        <button
          type="button"
          @click="togglePasswordVisibility"
          class="toggle-password"
        >
          {{ showPassword ? 'Hide' : 'Show' }}
        </button>
      </div>
      <div class="input-group">
        <select v-model="role" class="input-field" required>
          <option value="" disabled>Select Role</option>
          <option value="user">User</option>
          <option value="manager">Manager</option>
          <option value="admin">Admin</option>
        </select>
        <label class="input-label">Role</label>
      </div>
      <button type="submit" class="submit-button">Sign Up</button>
    </form>
    <p v-if="error" class="error-message">{{ error }}</p>
    <router-link to="/login" class="login-link">
      Already have an account? Login
    </router-link>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'SignupForm',
  setup() {
    const username = ref('')
    const lastname = ref('')
    const firstname = ref('')
    const password = ref('')
    const role = ref('user')
    const error = ref<string | null>(null)
    const showPassword = ref(false)

    const router = useRouter()

    const togglePasswordVisibility = () => {
      showPassword.value = !showPassword.value
    }

    const signup = async () => {
      try {
        await axios.post('http://localhost:8080/signup', {
          username: username.value,
          lastname: lastname.value,
          firstname: firstname.value,
          password: password.value,
          role: role.value,
        })
        router.push('/login')
      } catch (err) {
        if (axios.isAxiosError(err) && err.response) {
          error.value = err.response.data.message || 'Signup failed'
        } else {
          error.value = 'Network Error'
        }
        console.error('Signup error:', error.value)
      }
    }

    return {
      username,
      lastname,
      firstname,
      password,
      role,
      error,
      showPassword,
      togglePasswordVisibility,
      signup,
    }
  },
})
</script>

<style scoped>
.signup-container {
  padding: 2rem;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
  text-align: center;
  overflow: hidden;
  transition:
    transform 0.2s ease,
    box-shadow 0.2s ease;
  width: 100%;
}

.signup-title {
  font-size: 2rem;
  color: #4a2c61;
  font-family: 'Roboto', sans-serif;
}

.signup-form {
  display: flex;
  flex-direction: column;
}

.input-group {
  position: relative;
  margin-bottom: 1.5rem;
}

.input-field {
  padding: 0.8rem;
  border: 1px solid #d1c5e0;
  border-radius: 6px;
  transition: border 0.3s ease;
  background-color: #f9f1f7;
  width: 90%;
}

.input-field:focus {
  border-color: #c69dd8;
  outline: none;
}

.input-label {
  position: absolute;
  top: 0.8rem;
  left: 1rem;
  color: #aaa;
  pointer-events: none;
  transition: 0.2s;
}

.input-field:focus + .input-label,
.input-field:not(:placeholder-shown) + .input-label {
  top: -1rem;
  left: 1rem;
  font-size: 0.8rem;
  color: #4a2c61;
}

.toggle-password {
  position: absolute;
  right: 1rem;
  top: 0.5rem;
  background: #c69dd8;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 0.9rem;
  padding: 0.4rem 0.6rem;
  transition: background-color 0.3s ease;
}

.toggle-password:hover {
  background-color: #b358a1;
}

.submit-button {
  padding: 0.8rem;
  margin-top: 1rem;
  margin-bottom: 1rem;
  background-color: #c69dd8;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition:
    background-color 0.3s ease,
    transform 0.3s ease;
}

.submit-button:hover {
  background-color: #b358a1;
  transform: scale(1.02);
}

.error-message {
  color: red;
  margin-top: 0.5rem;
}

.login-link {
  margin-top: 1rem;
  color: #4a2c61;
  text-decoration: none;
}

.login-link:hover {
  text-decoration: underline;
}
</style>
