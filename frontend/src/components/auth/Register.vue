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
        <h1 class="auth-header">Register</h1>
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
    <p class="auth-login">
      Already have an account? <router-link to="/login">Login here</router-link>
    </p>
    </template>

<style scoped>
.auth-container {
  max-width: 400px;
  margin: 2rem auto; /* Center the container and add vertical spacing */
  padding: 2rem;
  background-color: #f9f9f9; /* Light background color for the form container */
  border-radius: 8px; /* Rounded corners for the container */
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); /* Subtle shadow for depth */
}

.auth-header {
  text-align: center; /* Center align the header text */
  margin-bottom: 1.5rem; /* Add spacing below the header */
  color: #333; /* Darker text color for better readability */
}

form {
  display: flex;
  flex-direction: column; /* Arrange form elements in a column */
}

form div {
  margin-bottom: 1rem; /* Space between form fields */
}

label {
  display: block; /* Ensure labels take full width */
  margin-bottom: 0.5rem; /* Space between label and input */
  font-weight: bold; /* Bold labels for better emphasis */
}

input {
  width: 100%; /* Full width for input fields */
  padding: 0.75rem; /* Padding inside the input fields */
  border: 1px solid #ccc; /* Light border for input fields */
  border-radius: 4px; /* Rounded corners for input fields */
  box-sizing: border-box; /* Include padding and border in element's total width and height */
}

button {
  padding: 0.75rem 1.5rem; /* Padding inside the button */
  border: none; /* Remove default border */
  border-radius: 4px; /* Rounded corners for the button */
  background-color: #007bff; /* Primary button color */
  color: #fff; /* White text color */
  font-size: 1rem; /* Font size for button text */
  cursor: pointer; /* Pointer cursor on hover */
  transition: background-color 0.3s ease; /* Smooth transition for background color */
}

button:hover {
  background-color: #0056b3; /* Darker shade on hover */
}

p {
  text-align: center; /* Center align the text */
  color: #666; /* Lighter text color for error messages */
}

.auth-login {
  margin-top: 1rem; /* Space above the registration link */
}
</style>
