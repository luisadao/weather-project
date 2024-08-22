<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { useToast } from 'vue-toastification'

const name = ref('')
const email = ref('')
const password = ref('')
const error = ref('')
const toast = useToast()

const router = useRouter()

const register = async () => {
try {
    await axios.post('/register', {
    name: name.value,
    email: email.value,
    password: password.value
    })
    toast.success("Registered sucessfully")    
    router.push('/login')
} catch (err) {
    error.value = 'Registration failed. Please try again.'
    toast.error("There was an error during the registration")
}
}
</script>

<template>
    <div class="auth-container">
        <h1>Register</h1>
        <form @submit.prevent="register">
        <div>
            <label for="name">Name</label>
            <input v-model="name" type="text" id="name" required />
        </div>
        <div>
            <label for="email">Email</label>
            <input v-model="email" type="email" id="email" required />
        </div>
        <div>
            <label for="password">Password</label>
            <input v-model="password" type="password" id="password" required />
        </div>
        <button type="submit">Register</button>
        </form>
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
