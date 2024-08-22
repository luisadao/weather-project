<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { useToast } from 'vue-toastification'

const email = ref('')
const password = ref('')
const error = ref('')
const toast = useToast()

const router = useRouter()

const login = async () => {
  try {
    const response = await axios.post('/login', {
      email: email.value,
      password: password.value
    }, {
      withCredentials: true 
    })
      toast.success("Logged in successfully")
        router.push('/')
  } catch (err) {
    error.value = 'Invalid email or password.'
    toast.error("There was an error logging in")
  }
}
</script>

<template>
  <div class="auth-container">
    <h1>Login</h1>
    <form @submit.prevent="login">
      <div>
        <label for="email">Email</label>
        <input v-model="email" type="email" id="email" required />
      </div>
      <div>
        <label for="password">Password</label>
        <input v-model="password" type="password" id="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
    <p>
      Don't have an account? <router-link to="/register">Register here</router-link>
    </p>
    <p v-if="error">{{ error }}</p>
  </div>
</template>



<style scoped>
.auth-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 2rem;
}
form div {
  margin-bottom: 1rem;
}
</style>
