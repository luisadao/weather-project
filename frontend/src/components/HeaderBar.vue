<script setup>
  import { useRouter } from 'vue-router'
  import axios from 'axios'
  import { useToast } from 'vue-toastification'
  const router = useRouter()
  
  const toast = useToast()

  const logout = async () => {
    try {
      await axios.post('/logout') 
      toast.success("Logged out successfully")
      router.push('/login') 
    } catch (error) {
      console.error('Logout failed:', error)
    }
  }
  </script>
  

<template>
    <div class="min-h-screen flex flex-col">
      <header class="header">
        <h1 class="title">Weather</h1>
        <button @click="logout" class="logout-button">
          Logout
        </button>
      </header>
    </div>
  </template>
  
  
  <style scoped>
  .header {
    width: 100%;
    background-color: #003366; 
    color: white;
    padding: 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    position: fixed;
    top: 0;
    left: 0;
    z-index: 10;
    
  }
  
  .title {
    font-size: 1.5rem;
    font-weight: bold;
  }
  
  .logout-button {
    background-color: #ff4d4d;
    color: white;
    border: none;
    padding: 0.5rem 1rem;
    border-radius: 0.25rem;
    cursor: pointer;
    font-weight: bold;
  }
  
  .logout-button:hover {
    background-color: #e60000; 
  }
  
  </style>
  