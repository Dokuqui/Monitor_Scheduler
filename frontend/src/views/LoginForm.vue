<template>
  <div class="login-container">
    <h1 class="login-title">Login to Your Account</h1>
    <form @submit.prevent="login" class="login-form">
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
      <button type="submit" class="submit-button">Login</button>
    </form>
    <p v-if="error" class="error-message">{{ error }}</p>
    <router-link to="/signup" class="signup-link">
      Don't have an account? Sign Up
    </router-link>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'LoginForm',
  setup() {
    const router = useRouter()
    const username = ref('')
    const password = ref('')
    const error = ref<string | null>(null)
    const showPassword = ref(false)

    const togglePasswordVisibility = () => {
      showPassword.value = !showPassword.value
    }

    const login = async () => {
      try {
        const response = await axios.post('http://localhost:8080/login', {
          username: username.value,
          password: password.value,
        })

        const token = response.data.token
        const redirectPath = response.data.redirect

        // Check if token is defined
        if (token) {
          localStorage.setItem('token', token)
          localStorage.setItem('userRole', response.data.role)
          router.push(redirectPath)
        } else {
          error.value = 'Invalid token'
        }
      } catch (err) {
        console.error(err)
        error.value = 'Invalid credentials'
      }
    }

    return {
      username,
      password,
      error,
      showPassword,
      togglePasswordVisibility,
      login,
    }
  },
})
</script>

<style scoped>
.login-container {
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

.login-title {
  font-size: 2rem;
  color: #4a2c61;
  font-family: 'Roboto', sans-serif;
}

.login-form {
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

.signup-link {
  margin-top: 1rem;
  color: #4a2c61;
  text-decoration: none;
}

.signup-link:hover {
  text-decoration: underline;
}
</style>
