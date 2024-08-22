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
    <h1 class="auth-header">Login</h1>
   
    <form @submit.prevent="login">
      <div>
        <label for="email">Email </label>
        <input v-model="email" type="email" id="email" required />
      </div>
      <div>
        <label for="password">Password </label>
        <input v-model="password" type="password" id="password" required />
      </div>
      <button type="submit">Login</button>
    </form>
    
    <p v-if="error">{{ error }}</p>
  </div>
  <p class="auth-register">
      Don't have an account? <router-link to="/register">Register here</router-link>
    </p>
</template>


<style scoped>
.auth-container {
  max-width: 400px;
  margin: 2rem auto; 
  padding: 2rem;
  background-color: #f9f9f9; 
  border-radius: 8px; 
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); 
}

.auth-header {
  text-align: center; 
  margin-bottom: 1.5rem; 
  color: #333; 
}

form {
  display: flex;
  flex-direction: column; 
}

form div {
  margin-bottom: 1rem; 
}

label {
  display: block; 
  margin-bottom: 0.5rem; 
  font-weight: bold; 
}

input {
  width: 100%; 
  padding: 0.75rem; 
  border: 1px solid #ccc; 
  border-radius: 4px; 
  box-sizing: border-box; 
}

button {
  padding: 0.75rem 1.5rem; 
  border: none; 
  border-radius: 4px; 
  background-color: #007bff; 
  color: #fff; 
  font-size: 1rem; 
  cursor: pointer; 
  transition: background-color 0.3s ease; 
}

button:hover {
  background-color: #0056b3; 
}

p {
  text-align: center; 
  color: #666; 
}

.auth-register {
  margin-top: 1rem; 
}
</style>