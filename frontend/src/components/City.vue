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
    flex-wrap: wrap;
    gap: 20px;
    justify-content: center;
    padding: 20px;
}

.img{
    border: 1px solid grey;
    border-radius: 5px;
    background-color: lightgrey;
}
</style>