<script setup>
import axios from 'axios'
    const props = defineProps(['weather'])

    const formattedTime = (time) => {
      return new Date(time).toLocaleString();
    }
    const API_URL = 'http://localhost:8080/weather'
    const ICON_URL = `https://openweathermap.org/img/wn/10d@2x.png`
    const fetchWeather = async (weather) => {
    try {
        const response = await axios.get(`${API_URL}/${weather.city}`)
        weather.value = response.data
    } catch (error) {
        console.log(error)
    }
}

</script>

<template>
    <div class="mt-3 d-flex">
        <div class="flex-shrink-1">
            <h5>{{ weather.city }}</h5>
            <!-- TODO: Fix url da imagem -->
            <img :src="ICON_URL">
            <p>{{ weather.description }}</p>
            <p>{{ weather.temperature }}°</p>
            <p>{{ formattedTime(weather.timestamp) }}</p>
            <button @click="fetchWeather(weather)" type="button" class="btn btn-success">Update time</button>
        </div>
        
    </div>
</template>

<style scoped>
.dashboard {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
  padding: 20px;
}
</style>