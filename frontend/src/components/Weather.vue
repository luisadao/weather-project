<script setup>
import { ref, watch, onMounted } from 'vue'
import axios from 'axios'
import City from './City.vue';
import HeaderBar from './HeaderBar.vue';
import { useToast } from 'vue-toastification'
const weathers = ref([])
const toast = useToast()

const fetchWeather = async () => {
    try {
        const response = await axios.get("/weather")
        weathers.value = response.data
    } catch (error) {
        console.log(error)
    }
}


onMounted(async () => {
    await fetchWeather()
    toast.success("Weather Updated")
    console.log(weathers.value)
})


</script>

<template>
    <HeaderBar>
    </HeaderBar>
    <div>
            <div class="dashboard" >
                <City v-for="weather in weathers" :weather="weather" :key="weather.id" class="city-widget">
                </City>
            </div>
        
    </div>
</template>


<style scoped>
.city-widget {
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 16px;
  width: 200px;
  margin: 10px;
  text-align: center;
  box-shadow: 2px 2px 12px rgba(0, 0, 0, 0.1);
}

.dashboard {
    margin: 200px;
  display: flex;
  flex-direction: row; 
  justify-content: center;
  align-items: center; 
  text-align: center;
}

.header {
  width: 100%;
  background-color: #003366; /* Dark blue color */
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
  background-color: #ff4d4d; /* Red color for logout button */
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 0.25rem;
  cursor: pointer;
  font-weight: bold;
}

.logout-button:hover {
  background-color: #e60000; /* Darker red color on hover */
}

main {
  margin-top: 4rem; /* Add margin to offset the fixed header */
}

</style>