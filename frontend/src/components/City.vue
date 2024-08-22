<script setup>
import axios from 'axios'
import { ref, watch } from 'vue'
import { useToast } from 'vue-toastification'

const props = defineProps(['weather'])
const toast = useToast()

const formattedTime = (time) => {
    return new Date(time).toLocaleString();
}
const API_URL = 'http://localhost:8080/weather'
const iconUrl = ref('')

const updateIcon = (icon) => {
    iconUrl.value = `https://openweathermap.org/img/wn/${icon}@2x.png`;
}

watch(() => props.weather, (newWeather) => {
  if (newWeather && newWeather.icon) {
    updateIcon(newWeather.icon);
  }
}, { immediate: true });

const fetchWeather = async (weather) => {
    try {
        const response = await axios.get(`${API_URL}/${weather.city}`)
        weather.value = response.data
        updateIcon(response.data.icon);
        toast.success(`Weather in ${weather.city} updated`)

    } catch (error) {
        toast.error("There was an error updating the weather")
        console.log(error)
    }
}

</script>

<template>
    <div class="mt-3 d-flex">
        <div class="flex-shrink-1">
            <h5>{{ weather.city }}</h5>
            <img :src="iconUrl" class="img">
            <p>{{ weather.description }}</p>
            <p>{{ weather.temperature }}°</p>
            <p>{{ formattedTime(weather.timestamp) }}</p>
            <button @click="fetchWeather(weather)" type="button" class="btn btn-success">Update weather</button>
        </div>

    </div>
</template>

<style scoped>
.dashboard {
    display: flex;
    flex-wrap:wrap;
    gap: 20px;
    justify-content: center;
    padding: 20px;
    
}

.img {
    border: 1px solid grey;
    border-radius: 5px;
    background-color: lightgrey;
    margin-bottom: 10px; 
    width: 100px; 
    height: 100px; 
    object-fit: cover; 
}

.flex-shrink-1 {
    display: flex;
    flex-direction: column; 
    align-items: center; 
    text-align: center; 
    padding: 1rem; 
    border: 1px solid lightgrey; 
    border-radius: 8px; 
    background-color: #f9f9f9; 
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1); 
}

h5 {
    margin-bottom: 0.5rem; 
}

p {
    margin: 0.5rem 0; 
}

.btn {
    padding: 0.5rem 1rem; 
    border: none; 
    border-radius: 4px; 
    background-color: #28a745; 
    color: white; 
    cursor: pointer; 
    transition: background-color 0.3s ease; 
}

.btn:hover {
    background-color: #218838; 
}
</style>